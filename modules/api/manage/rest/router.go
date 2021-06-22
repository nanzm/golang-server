package rest

import (
	"dora/modules/api/manage/handler"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	// cors
	g.Use(middleware.CORSMiddleware(), middleware.RateLimitMiddleware())

	api := g.Group("/api")
	ginutil.SetupResource(api,
		handler.NewDashboardResource(),
		handler.NewIssueResource(),
		handler.NewProjectResource(),
		handler.NewSyslogResource(),
		handler.NewUserResource(),
	)
}
