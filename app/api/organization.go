package api

import (
	"dora/app/dao"
	"dora/app/dto"
	"dora/app/middleware"
	"dora/app/model"
	"dora/config"
	"dora/pkg/ginutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationResource struct {
	Conf *config.Conf
}

func NewOrganizationResource() Resource {
	return &OrganizationResource{
		Conf: config.GetConf(),
	}
}

func (org *OrganizationResource) Register(router *gin.RouterGroup) {
	router.GET("/user/organizations", middleware.JWTAuthMiddleware(), org.UserOrganizations)
	router.GET("/organization", middleware.JWTAuthMiddleware(), org.Get)
	router.POST("/organization", middleware.JWTAuthMiddleware(), org.Create)

	// 组织成员管理
	router.GET("/organization/member", middleware.JWTAuthMiddleware(), org.GetMembers)
	router.POST("/organization/member/add", middleware.JWTAuthMiddleware(), org.AddMember)
	router.POST("/organization/member/remove", middleware.JWTAuthMiddleware(), org.RemoveMember)
}

func (org *OrganizationResource) Get(c *gin.Context) {
	var u dto.QueryOrganizationDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewOrganizationDao()
	get, err := d.Get(u.OrganizationId)

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (org *OrganizationResource) Create(c *gin.Context) {
	var q dto.CreateOrganization
	if err := c.ShouldBind(&q); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	uid, _ := c.Get("uid")

	organization := model.Organization{
		Name:         q.Name,
		Introduction: q.Introduction,
		Type:         q.Type,
		CreateUid:    uid.(uint),
	}
	d := dao.NewOrganizationDao()
	result, err := d.Create(uid.(uint), &organization)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 切换当前组织
	settingDao := dao.NewUserSettingDao()
	err = settingDao.UpdateOrCreate(uid.(uint), 0, organization.ID)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, result)
}

func (org *OrganizationResource) UserOrganizations(c *gin.Context) {
	get, _ := c.Get("uid")
	organizationDao := dao.NewOrganizationDao()
	userOrgList, err := organizationDao.UserOrganization(get.(uint))
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONList(c, userOrgList, len(userOrgList))
}

func (org *OrganizationResource) GetMembers(c *gin.Context) {
	var q dto.GetOrganizationMembers
	if err := c.ShouldBind(&q); err != nil {
		errorTrans(c, err)
		return
	}
	organizationDao := dao.NewOrganizationDao()
	members, err := organizationDao.GetMembers(q.OrganizationId)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONList(c, members, len(members))
}

func (org *OrganizationResource) AddMember(c *gin.Context) {
	var q dto.AddOrganizationMembers
	if err := c.ShouldBind(&q); err != nil {
		errorTrans(c, err)
		return
	}
	organizationDao := dao.NewOrganizationDao()
	err := organizationDao.AddUser(q.OrganizationId, q.UserIds)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONOk(c, nil)
}

func (org *OrganizationResource) RemoveMember(c *gin.Context) {
	var q dto.RemoveOrganizationMembers
	if err := c.ShouldBind(&q); err != nil {
		errorTrans(c, err)
		return
	}
	organizationDao := dao.NewOrganizationDao()
	err := organizationDao.RemoveUser(q.OrganizationId, q.UserId)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONOk(c, nil)
}
