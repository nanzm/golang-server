package cmd

import (
	"context"
	"dora/app/boots"
	"dora/app/mqConsumer"
	"dora/app/task"
	"dora/config"
	"dora/pkg"
	"dora/pkg/ginutil"
	"dora/pkg/logger"

	"dora/app/api"
	"dora/app/datasource"
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

var configFile string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start dora server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("server starting...")

		c := getConf()
		serverRun(c)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.yml)")
}

func getConf() *config.Conf {
	var path string
	if configFile != "" {
		path = configFile
		logger.Printf("config flags: %v", configFile)
	} else {
		path = "./config.yml"
		logger.Println("use default config")
	}
	return config.ParseConf(path)
}

func serverRun(conf *config.Conf) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// mail
	datasource.GetMailPool()

	// redis
	datasource.RedisInstance()

	// msql
	datasource.GormInstance()

	// 启动初始化
	boots.Run()

	// nsq 消费上报数据 -> 放入日志服务
	datasource.NsqConsumerRegister(conf.Nsq, mqConsumer.Consumer())

	// 启动定时任务
	// 1：告警监控
	// 2：创建issues
	task.StartCron()

	// 启动 http 服务
	startHttpServer(conf)
}

func startHttpServer(conf *config.Conf) {
	// 初始化翻译器
	if err := ginutil.InitTrans("zh"); err != nil {
		panic(err)
	}

	g := gin.New()
	//gin.SetMode("release")

	g.Use(pkg.GinZap(logger.L, false), pkg.RecoveryWithZap(logger.L, false))

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
		logger.Printf("Dora server listen at http://127.0.0.1%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	gracefulStop(srv)
}

// 优雅关闭
func gracefulStop(srv *http.Server) {
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
		logger.Warn("http server exit timeout of 5 seconds.")
	default:
	}

	logger.Warn("http server exited.")

	datasource.StopNsq()
	datasource.StopSlsLog()
	datasource.StopRedisClient()
	datasource.StopDataBase()
	datasource.StopMailPool()
}
