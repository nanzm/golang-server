package rest

import (
	"dora/internal/apps/manage/handler"
	"dora/internal/middleware"
	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	// cors
	g.Use(middleware.CORSMiddleware(), middleware.RateLimitMiddleware())

	public := g.Group("/")
	ginutil.SetupResource(public,
		handler.NewPublicResource(),
	)

	api := g.Group("/api")
	ginutil.SetupResource(api,
		handler.NewDashboardResource(),
		handler.NewIssueResource(),
		handler.NewProjectResource(),
		handler.NewSyslogResource(),
		handler.NewAlarmResource(),
		handler.NewUserResource(),
	)
}
