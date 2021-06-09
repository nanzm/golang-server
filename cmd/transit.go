package cmd

import (
	"context"
	"dora/config"
	"dora/internal/middleware"
	"dora/internal/mqConsumer"
	"dora/pkg/ginutil"
	"dora/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"dora/internal/api"
	"dora/internal/datasource"
	"net/http"
	"time"

	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var transitCmd = &cobra.Command{
	Use:   "transit",
	Short: "Start dora transit",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("transit starting...")

		c := getConf()
		transitRun(c)
	},
}

func init() {
	rootCmd.AddCommand(transitCmd)
	transitCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.yml)")
}

func transitRun(conf *config.Conf) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// redis
	datasource.RedisInstance()

	// nsq 消费上报数据 -> 放入日志服务
	datasource.NsqConsumerRegister(conf.Nsq, mqConsumer.Consumer())

	// 启动 http 服务
	startTransit(conf)
}

func startTransit(conf *config.Conf) {
	// 初始化翻译器
	if err := ginutil.InitTrans("zh"); err != nil {
		panic(err)
	}

	g := gin.New()
	//gin.SetMode("release")

	g.Use(middleware.GinZap(logger.L, false), middleware.RecoveryWithZap(logger.L, false))

	// session
	store := cookie.NewStore([]byte(conf.Secret))
	g.Use(sessions.Sessions("dora", store))

	api.Register(g, conf)

	srv := &http.Server{
		Addr:         ":8222",
		Handler:      g,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logger.Printf("Dora transit listen at http://127.0.0.1%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	stopTransit(srv)
}

// 优雅关闭
func stopTransit(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("正在退出...")

	logger.Println("shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("http server shutdown err: ", err)
	}

	select {
	case <-ctx.Done():
		logger.Info("http server exit timeout of 5 seconds.")
	default:
	}

	logger.Info("http server exited.")

	datasource.StopNsq()
	datasource.StopSlsLog()
}
