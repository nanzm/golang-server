package handler

import (
	"dora/app/manage/model/dto"
	"dora/config"
	"dora/modules/logstore"
	"dora/modules/middleware"
	"dora/pkg/utils"
	"net/http"

	"dora/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
)

type DashboardResource struct {
}

func NewDashboardResource() ginutil.Resource {
	return &DashboardResource{}
}

func (da *DashboardResource) Register(router *gin.RouterGroup) {
	router.POST("/dashboard/dsl-search", da.Search)
	router.GET("/dashboard/logs", middleware.JWTAuthMiddleware(), da.QueryLogsByMd5)
	router.GET("/dashboard/chart/:type", da.QueryChartData)

	router.GET("/system", da.SystemInfo)
}

// todo 校验数据安全性 不该看的东西不能看
func (da *DashboardResource) Search(c *gin.Context) {
	var u dto.SearchParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}

	s := logstore.GetClient()
	searchResult, err := s.DSLSearch(u.Body)
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSONOk(c, searchResult)
}

func (da *DashboardResource) QueryLogsByMd5(c *gin.Context) {
	var u dto.QueryEventsByMd5Param
	if err := c.ShouldBind(&u); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}

	s := logstore.GetClient()
	md5Log, err := s.QueryMethods().GetErrorLogsByMd5(u.AppId, u.Start, u.End, u.Md5)
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}
	ginutil.JSONOk(c, md5Log)
}

// 查询图表 数据
func (da *DashboardResource) QueryChartData(c *gin.Context) {
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

func (da *DashboardResource) SystemInfo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name":    "dora-manage",
		"build":   config.Build,
		"compile": config.Compile,
		"version": config.Version,
		"uptime":  utils.TimeFromNow(config.Uptime),
		"now":     utils.CurrentTime(),
	})
}
