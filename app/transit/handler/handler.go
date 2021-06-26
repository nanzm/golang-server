package handler

import (
	"dora/config"
	"dora/modules/datasource/nsq"
	"dora/pkg/utils"
	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var uptime = time.Now()

type PublicResource struct{}

func NewPublicResource() ginutil.Resource {
	return &PublicResource{}
}

func (pub *PublicResource) Register(router *gin.RouterGroup) {
	router.Any("/", pub.Info)
	router.HEAD("/ping", pub.Ping)
	router.GET("/ping", pub.Ping)

	// img 上报
	router.GET("/collect", pub.CollectQueryData)
	// post 请求
	router.POST("/collect", pub.CollectBodyData)
}

func (pub *PublicResource) Info(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name":    "dora",
		"build":   config.Build,
		"compile": config.Compile,
		"version": config.Version,
		"uptime":  utils.TimeFromNow(uptime),
		"now":     utils.CurrentTime(),
	})
}

func (pub *PublicResource) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (pub *PublicResource) CollectQueryData(c *gin.Context) {
	eventMap := c.Request.URL.Query()
	c.JSON(http.StatusOK, eventMap)
}

func (pub *PublicResource) CollectBodyData(c *gin.Context) {
	eventRaw, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if string(eventRaw) == "" {
		c.String(http.StatusOK, "empty")
		return
	}

	// 处理
	events, err := HandleEvent(c, eventRaw)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	// 给mq
	nsqConf := config.GetNsq()
	err = nsq.ProducerInstance().Publish(nsqConf.Topic, events)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}
