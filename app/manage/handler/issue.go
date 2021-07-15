package handler

import (
	"dora/app/manage/model/dto"
	"dora/modules/logstore"
	"dora/modules/middleware"
	"dora/pkg/utils/ginutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

type IssueResource struct {
}

func NewIssueResource() ginutil.Resource {
	return &IssueResource{
	}
}

func (issue *IssueResource) Register(router *gin.RouterGroup) {
	// issue
	router.GET("/issues", middleware.JWTAuthMiddleware(), issue.List)
	router.GET("/issue", middleware.JWTAuthMiddleware(), issue.DetailByMd5)
	router.POST("/issue/ignore", middleware.JWTAuthMiddleware(), issue.Ignore)
	router.POST("/issue/solve", middleware.JWTAuthMiddleware(), issue.Solve)
}

func (issue *IssueResource) List(c *gin.Context) {
	var u dto.IssueListParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	list, err := logstore.GetClient().QueryMethods().ErrorList(u.AppId, u.Start, u.End)
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	ginutil.JSONListPages(c, list.List, 1, 100, int64(list.Total))
}

func (issue *IssueResource) DetailByMd5(c *gin.Context) {
	var u dto.IssueDetailParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	result, err := logstore.GetClient().QueryMethods().GetLogByMd5(u.AppId, u.Start, u.End, u.Md5)
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, result)
}

func (issue *IssueResource) Ignore(c *gin.Context) {

}

func (issue *IssueResource) Solve(c *gin.Context) {

}
