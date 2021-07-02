package handler

import (
	"dora/app/manage/model/dao"
	"dora/app/manage/model/dto"
	"dora/app/manage/model/entity"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"

	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type AlarmResource struct {
}

func NewAlarmResource() ginutil.Resource {
	return &AlarmResource{
	}
}

func (pro *AlarmResource) Register(router *gin.RouterGroup) {
	router.GET("/alarms", middleware.JWTAuthMiddleware(), pro.Get)
	router.POST("/alarm", middleware.JWTAuthMiddleware(), pro.CreateOrUpdate)
}

func (pro *AlarmResource) Get(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *AlarmResource) CreateOrUpdate(c *gin.Context) {
	uid, _ := c.Get("uid")
	var body dto.CreateProject
	if err := c.ShouldBind(&body); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	project := entity.Project{
		AppId:             uuid.New().String(),
		Name:              body.Name,
		Alias:             body.Alias,
		Type:              body.Type,
		GitRepositoryUrl:  body.GitRepositoryUrl,
		GitRepositoryName: body.GitRepositoryName,
	}

	d := dao.NewProjectDao()

	// 校验项目名 名字
	p, err := d.GetByName(body.Name)
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	if p.ID != 0 {
		ginutil.JSONFail(c, -1, "该项目名已存在")
		return
	}

	// 创建
	result, err := d.Create(&project, uid.(uint))
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	// 切换到当前项目
	settingDao := dao.NewUserSettingDao()
	err = settingDao.UpdateOrCreate(uid.(uint), result.ID)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, result)
}

func (pro *AlarmResource) Status(c *gin.Context) {
	uid, _ := c.Get("uid")
	var body dto.CreateProject
	if err := c.ShouldBind(&body); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	project := entity.Project{
		AppId:             uuid.New().String(),
		Name:              body.Name,
		Alias:             body.Alias,
		Type:              body.Type,
		GitRepositoryUrl:  body.GitRepositoryUrl,
		GitRepositoryName: body.GitRepositoryName,
	}

	d := dao.NewProjectDao()

	// 校验项目名 名字
	p, err := d.GetByName(body.Name)
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	if p.ID != 0 {
		ginutil.JSONFail(c, -1, "该项目名已存在")
		return
	}

	// 创建
	result, err := d.Create(&project, uid.(uint))
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	// 切换到当前项目
	settingDao := dao.NewUserSettingDao()
	err = settingDao.UpdateOrCreate(uid.(uint), result.ID)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, result)
}

func (pro *AlarmResource) Delete(c *gin.Context) {
	uid, _ := c.Get("uid")
	var body dto.CreateProject
	if err := c.ShouldBind(&body); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	project := entity.Project{
		AppId:             uuid.New().String(),
		Name:              body.Name,
		Alias:             body.Alias,
		Type:              body.Type,
		GitRepositoryUrl:  body.GitRepositoryUrl,
		GitRepositoryName: body.GitRepositoryName,
	}

	d := dao.NewProjectDao()

	// 校验项目名 名字
	p, err := d.GetByName(body.Name)
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	if p.ID != 0 {
		ginutil.JSONFail(c, -1, "该项目名已存在")
		return
	}

	// 创建
	result, err := d.Create(&project, uid.(uint))
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	// 切换到当前项目
	settingDao := dao.NewUserSettingDao()
	err = settingDao.UpdateOrCreate(uid.(uint), result.ID)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, result)
}

// 告警记录
func (pro *AlarmResource) GetLogs(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *AlarmResource) RemoveLogs(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}


// 联系方式
func (pro *AlarmResource) GetContacts(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *AlarmResource) CreateContacts(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *AlarmResource) UpdateContacts(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *AlarmResource) RemoveContacts(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

