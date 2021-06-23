package handler

import (
	"context"
	"dora/config"
	"dora/modules/datasource"
	"dora/modules/model/dto"
	"dora/pkg/utils/ginutil"
	"dora/pkg/utils/logx"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"time"
)

type PublicResource struct {
}

func NewPublicResource() ginutil.Resource {
	return &PublicResource{
	}
}

func (pub *PublicResource) Register(router *gin.RouterGroup) {
	// 上报
	router.POST("/public/report", pub.TransToNsq)
	router.POST("/report", pub.TransToNsq)
	router.POST("/v2/report", pub.TransToNsq)

	// 测试用
	router.Any("/http/delay", pub.HTTPDelay)
	router.Any("/http/error", pub.HTTPError)

	router.GET("/test", pub.Test)
	router.GET("/mail", pub.SendMail)
	router.GET("/dingding", pub.DingDing)
}

func (pub *PublicResource) TransToNsq(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// 解析
	var eventData map[string]interface{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &eventData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// TODO 校验 appId 合法性
	// 校验
	if val, ok := eventData["_appId"]; !ok || val == "" {
		c.String(http.StatusBadRequest, "missing key \"_appId\"")
		return
	}
	if val, ok := eventData["category"]; !ok || val == "" {
		c.String(http.StatusBadRequest, "missing key \"category\"")
		return
	}

	// 添加ip
	var ip = c.ClientIP()
	eventData["ip"] = ip

	// 序列化
	marshal, err := json.Marshal(eventData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// 给mq
	nsgConfig := config.GetNsq()
	err = datasource.NsqProducerInstance().Publish(nsgConfig.Topic, marshal)
	if err != nil {
		logx.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

func (pub *PublicResource) HTTPDelay(c *gin.Context) {
	var p dto.DelayParam
	if err := c.ShouldBind(&p); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}

	if p.Delay != 0 {
		time.Sleep(time.Duration(p.Delay) * time.Second)
		c.String(200, fmt.Sprintf("response delay %v second", p.Delay))
		return
	}

	c.String(200, "ok")
}

func (pub *PublicResource) HTTPError(c *gin.Context) {
	var p dto.ErrorParam
	if err := c.ShouldBind(&p); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}
	if p.Status != 0 && p.Status >= 100 && p.Status <= 512 {
		c.String(p.Status, fmt.Sprintf("http error test: %v", http.StatusText(p.Status)))
		return
	}
	c.String(http.StatusInternalServerError, http.StatusText(500))
}

func (pub *PublicResource) Test(c *gin.Context) {
	ctx := context.Background()
	val, err := datasource.RedisInstance().Get(ctx, "ddd").Result()
	if err != nil {
		if err == redis.Nil {
			logx.Println("key does not exists")
			return
		}
		panic(err)
	}
	logx.Println("redis ping: ", val)
}

func (pub *PublicResource) SendMail(c *gin.Context) {
	mailCof := config.GetMail()
	m := datasource.BuilderEmail("msg@nancode.cn", fmt.Sprintf("Dora System Robot <%s>", mailCof.Username),
		"test", "hello world")
	err := datasource.GetMailPool().Send(m, 3*time.Second)

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	ginutil.JSONOk(c, nil)
}

func (pub *PublicResource) DingDing(c *gin.Context) {
	//msg := c.DefaultQuery("msg", "测试 [鼓掌]")
	//
	//secret := pub.Conf.DingDing[0].Secret
	//accessToken := pub.Conf.DingDing[0].AccessToken
	//
	//data := service.NewDingTalkMsg(msg)
	//err := service.SendDingDing(data, secret, accessToken)
	//if err != nil {
	//	ginutil.JSONError(c, http.StatusInternalServerError, err)
	//}
	//ginutil.JSONOk(c, nil)
}