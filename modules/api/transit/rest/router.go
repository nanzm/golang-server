package rest

import (
	"dora/config"
	"dora/modules/api/transit/handler"
	"dora/modules/middleware"
	"dora/pkg/utils"
	"dora/pkg/utils/ginutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var uptime = time.Now()

func Register(g *gin.Engine) {
	// cors
	g.Use(middleware.CORSMiddleware())

	g.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name":    "dora",
			"build":   config.Build,
			"compile": config.Compile,
			"version": config.Version,
			"uptime":  utils.TimeFromNow(uptime),
			"now":     utils.CurrentTime(),
		})
	})

	// pong
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
	g.HEAD("/ping", pingHandler)
	g.GET("/ping", pingHandler)

	pubRouter := g.Group("/")
	ginutil.SetupResource(pubRouter, handler.NewPublicResource())
}
