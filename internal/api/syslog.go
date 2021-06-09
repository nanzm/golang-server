package api

import (
	"dora/internal/datasource"
	"dora/internal/dto"
	"dora/internal/model"
	"dora/config"
	"dora/pkg/ginutil"
	"dora/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SyslogResource struct {
	Conf *config.Conf
}

func NewSyslogResource() Resource {
	return &SyslogResource{
		Conf: config.GetConf(),
	}
}

func (s *SyslogResource) Register(router *gin.RouterGroup) {

}

func (s *SyslogResource) GetParseErrorList(c *gin.Context) {
	var u dto.ParseErrorParam
	if err := c.ShouldBind(&u); err != nil {
		errorTrans(c, err)
		return
	}

	var list []*model.SysLog
	var total int64
	err := datasource.GormInstance().Scopes(utils.Paginate(u.Current, u.PageSize)).
		Find(&list).Count(&total).Order("id desc").Error

	if err != nil {
		ginutil.JSONServerError(c, err)
	}

	ginutil.JSONListPages(c, list, u.Current, u.PageSize, total)
}
