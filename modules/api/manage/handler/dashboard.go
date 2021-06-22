package handler

import (
	"context"

	"dora/modules/datasource"
	"dora/modules/logstore"
	"dora/modules/middleware"
	"dora/modules/model/dto"

	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

type DashboardResource struct {
}

func NewDashboardResource() ginutil.Resource {
	return &DashboardResource{}
}

func (issue *DashboardResource) Register(router *gin.RouterGroup) {
	router.GET("/dashboard/events", middleware.JWTAuthMiddleware(), issue.QueryEventsByMd5)
	router.GET("/dashboard/events/chart/:type", issue.QueryChartData)

	router.GET("/logstore/switch", issue.SwitchLogStore)
}

func (issue *DashboardResource) QueryEventsByMd5(c *gin.Context) {
	var u dto.QueryEventsByMd5Param
	if err := c.ShouldBind(&u); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}

	s := logstore.GetClient()
	md5Log, err := s.QueryMethods().GetLogByMd5(u.Start, u.End, u.Md5)
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}
	ginutil.JSONOk(c, md5Log)
}

// 查询图表 数据
func (issue *DashboardResource) QueryChartData(c *gin.Context) {
	var u dto.ChartData
	if err := c.ShouldBind(&u); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}
	dataType := c.Param("type")

	s := logstore.GetClient()
	data, err := s.DefaultQuery(u.AppId, u.Start, u.End, u.Interval, dataType)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, data)
}

func (issue *DashboardResource) SwitchLogStore(c *gin.Context) {
	const StoreSwitch = "logStoreSwitch"

	result, err := datasource.RedisInstance().Get(context.Background(), StoreSwitch).Result()
	if err != nil && err != redis.Nil {
		ginutil.JSONServerError(c, err)
		return
	}

	if result == "" {
		result, err := datasource.RedisInstance().Set(context.Background(), StoreSwitch, time.Now(), time.Hour*24).Result()
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
		ginutil.JSONOk(c, result, "已切换为 elastic")
		return

	} else {
		result, err := datasource.RedisInstance().Del(context.Background(), StoreSwitch).Result()
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
		ginutil.JSONOk(c, result, "已切换为 阿里云sls")
		return

	}
}
