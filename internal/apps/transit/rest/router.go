package rest

import (
	"dora/internal/apps/transit/handler"
	"dora/internal/middleware"
	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	// cors
	g.Use(middleware.CORSMiddleware())

	pubRouter := g.Group("/")
	ginutil.SetupResource(pubRouter, handler.NewPublicResource())
}
