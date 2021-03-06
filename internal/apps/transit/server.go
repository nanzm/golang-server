package transit

import (
	"context"
	"dora/internal/apps/transit/boot"
	"dora/internal/apps/transit/rest"
	"dora/internal/config"
	"dora/internal/middleware"
	"dora/pkg/utils/ginutil"
	"dora/pkg/utils/logx"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// http://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=Dora%20transit%0A
const banner = `
██████╗  ██████╗ ██████╗  █████╗     ████████╗██████╗  █████╗ ███╗   ██╗███████╗██╗████████╗
██╔══██╗██╔═══██╗██╔══██╗██╔══██╗    ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██║╚══██╔══╝
██║  ██║██║   ██║██████╔╝███████║       ██║   ██████╔╝███████║██╔██╗ ██║███████╗██║   ██║   
██║  ██║██║   ██║██╔══██╗██╔══██║       ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██║   ██║   
██████╔╝╚██████╔╝██║  ██║██║  ██║       ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║   ██║   
╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝       ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝   ╚═╝   
                                                                                            
`

func Serve() {
	fmt.Print(banner)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	boot.Setup()
	defer boot.TearDown()

	if err := ginutil.InitTrans("zh"); err != nil {
		logx.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	// pprof
	middleware.DebugRegister(app)

	app.Use(middleware.GinZap(), middleware.Recovery(false))

	// session
	store := cookie.NewStore([]byte(config.GetTransitSecret().Secret))
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
		logx.Printf("transit server listen at http://127.0.0.1%s", svr.Addr)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Fatalf("Error ListenAndServe : %s", err)
			return
		}
	}()

	ginutil.GracefulShutdown(ctx, svr)
}
