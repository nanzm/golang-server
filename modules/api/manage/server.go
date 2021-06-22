package manage

import (
	"dora/config"
	"dora/modules/api/manage/core"
	"dora/modules/api/manage/rest"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"
	"dora/pkg/utils/logx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Serve() *http.Server {
	core.Setup()
	defer core.TearDown()

	if err := ginutil.InitTrans("zh"); err != nil {
		logx.Fatal(err)
	}
	app := gin.New()
	gin.SetMode(gin.ReleaseMode)
	// session
	store := cookie.NewStore([]byte(config.GetApp().Secret))
	app.Use(sessions.Sessions("dora", store))

	app.Use(middleware.GinZap(), middleware.Recovery(false))
	rest.Register(app)

	svr := &http.Server{
		Addr:         ":8222",
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

	return svr
}
