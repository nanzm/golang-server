package rest

import (
	"dora/modules/api/transit/handler"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	// cors
	g.Use(middleware.CORSMiddleware())

	pubRouter := g.Group("/")
	ginutil.SetupResource(pubRouter, handler.NewPublicResource())
}
