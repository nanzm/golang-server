package cmd

import (
	"context"
	"dora/config"
	"dora/internal/boots"
	"dora/internal/middleware"
	"dora/internal/task"
	"dora/pkg/ginutil"
	"dora/pkg/logger"

	"dora/internal/api"
	"dora/internal/datasource"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Start dora manage",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("manage starting...")

		c := getConf()
		manageRun(c)
	},
}

func init() {
	rootCmd.AddCommand(manageCmd)
	manageCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.yml)")
}

func manageRun(conf *config.Conf) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// mail
	datasource.GetMailPool()

	// redis
	datasource.RedisInstance()

	// database
	datasource.GormInstance()

	// 启动初始化
	boots.Run()

	// 启动定时任务
	// 1：告警监控
	// 2：创建issues
	task.StartCron()

	// 启动 http 服务
	startManageServer(conf)
}

func startManageServer(conf *config.Conf) {
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
		logger.Printf("Dora manage listen at http://127.0.0.1%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	stopManage(srv)
}

// 优雅关闭
func stopManage(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("正在退出...")

	logger.Println("shutdown manage Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("http manage shutdown err: ", err)
	}

	select {
	case <-ctx.Done():
		logger.Info("http manage exit timeout of 5 seconds.")
	default:
	}

	logger.Info("http manage exited.")

	datasource.StopSlsLog()
	datasource.StopRedisClient()
	datasource.StopDataBase()
	datasource.StopMailPool()
}
