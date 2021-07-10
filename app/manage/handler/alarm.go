package handler

import (
	"dora/app/manage/model/dao"
	"dora/app/manage/model/dto"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"

	"github.com/gin-gonic/gin"
	"net/http"
)

type AlarmResource struct {
}

func NewAlarmResource() ginutil.Resource {
	return &AlarmResource{
	}
}

func (pro *AlarmResource) Register(router *gin.RouterGroup) {
	router.GET("/alarm", middleware.JWTAuthMiddleware(), pro.Get)
	router.POST("/alarm", middleware.JWTAuthMiddleware(), pro.CreateOrUpdate)
	router.GET("/alarms", middleware.JWTAuthMiddleware(), pro.List)
}

func (pro *AlarmResource) Get(c *gin.Context) {

}

func (pro *AlarmResource) List(c *gin.Context) {
	var u dto.AlarmListParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewAlarmDao()
	list, count, e := d.ListWithQuery(u.Current, u.PageSize, u.AppId)
	if e != nil {
		ginutil.JSONServerError(c, e)
		return
	}

	ginutil.JSONListPages(c, list, u.Current, u.PageSize, count)
}

func (pro *AlarmResource) CreateOrUpdate(c *gin.Context) {

}

func (pro *AlarmResource) Status(c *gin.Context) {

}

func (pro *AlarmResource) Delete(c *gin.Context) {

}

// 告警记录
func (pro *AlarmResource) GetLogs(c *gin.Context) {

}

func (pro *AlarmResource) RemoveLogs(c *gin.Context) {

}

// 联系方式
func (pro *AlarmResource) GetContacts(c *gin.Context) {

}

func (pro *AlarmResource) CreateContacts(c *gin.Context) {

}

func (pro *AlarmResource) UpdateContacts(c *gin.Context) {

}

func (pro *AlarmResource) RemoveContacts(c *gin.Context) {

}
