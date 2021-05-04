package ginutil

import (
	"dora/pkg/httputil"
	"dora/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONOk(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusOK, httputil.NewOkJSON(data, message...))
}

func JSONFail(c *gin.Context, bizCode int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, httputil.NewFailJSON(bizCode, msg))
}

func JSONValidatorFail(c *gin.Context, bizCode int, validate interface{}, message ...string) {
	c.AbortWithStatusJSON(http.StatusOK, httputil.NewValidatorFailJSON(bizCode, validate, message...))
}

func JSONError(c *gin.Context, httpStatus int, err error, msg ...string) {
	c.AbortWithStatusJSON(httpStatus, httputil.NewErrorJSON(-1, err.Error(), msg...))
	e := c.Error(err)
	if e != nil {
		logger.Printf("err: %v", e)
	}
}

// 不分页
func JSONList(c *gin.Context, list interface{}, total int) {
	c.JSON(http.StatusOK, httputil.NewOkJSON(gin.H{
		"list":  list,
		"total": total,
	}))
}

// 分页
func JSONListPages(c *gin.Context, list interface{}, current, pageSize int, total int64) {
	if current == 0 {
		current = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	c.JSON(http.StatusOK, httputil.NewOkJSON(gin.H{
		"list":     list,
		"current":  current,
		"pageSize": pageSize,
		"total":    total,
	}))
}

// BadRequest
func JSONBadRequest(c *gin.Context, err error) {
	JSONError(c, http.StatusBadRequest, err)
}

// ServerError
func JSONServerError(c *gin.Context, err error) {
	JSONError(c, http.StatusInternalServerError, err)
}

// Cookie
func Cookie(c *gin.Context, name, value string, maxAge int) {
	c.SetCookie(name, value, maxAge, "/", "", false, false)
}

// FoundRedirect redirect with the StatusFound
func FoundRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
	c.Abort()
}

// MovedRedirect redirect with the StatusMovedPermanently
func MovedRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusMovedPermanently, location)
	c.Abort()
}

// TemporaryRedirect redirect with the StatusTemporaryRedirect
func TemporaryRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusTemporaryRedirect, location)
	c.Abort()
}
