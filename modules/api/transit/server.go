package transit

import (
	"context"
	"dora/config"
	"dora/modules/api/transit/core"
	"dora/modules/api/transit/rest"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"
	"dora/pkg/utils/logx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Serve() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	core.Setup()
	defer core.TearDown()

	if err := ginutil.InitTrans("zh"); err != nil {
		logx.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(middleware.GinZap(), middleware.Recovery(false))

	// session
	store := cookie.NewStore([]byte(config.GetApp().Secret))
	app.Use(sessions.Sessions("dora", store))

	rest.Register(app)

	svr := &http.Server{
		Addr:         ":8221",
		Handler:      app,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 5 * 60,
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Fatalf("Error ListenAndServe : %s", err)
			return
		}
	}()

	ginutil.GracefulShutdown(ctx, svr)
}