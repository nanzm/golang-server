package handler

import (
	"dora/config"
	"dora/pkg/utils"
	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
	"net/http"
)


type PublicResource struct{}

func NewPublicResource() ginutil.Resource {
	return &PublicResource{}
}

func (pub *PublicResource) Register(router *gin.RouterGroup) {
	router.Any("/", pub.Info)
	router.HEAD("/ping", pub.Ping)
	router.GET("/ping", pub.Ping)
}

func (pub *PublicResource) Info(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name":    "dora-manage",
		"build":   config.Build,
		"compile": config.Compile,
		"version": config.Version,
		"uptime":  utils.TimeFromNow(config.Uptime),
		"now":     utils.CurrentTime(),
	})
}

func (pub *PublicResource) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
