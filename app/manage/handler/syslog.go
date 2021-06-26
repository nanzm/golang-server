package handler

import (
	"dora/app/manage/model/dto"
	"dora/app/manage/model/entity"
	"dora/modules/datasource/gorm"
	"dora/pkg/utils"
	"dora/pkg/utils/ginutil"

	"github.com/gin-gonic/gin"
)

type SyslogResource struct {
}

func NewSyslogResource() ginutil.Resource {
	return &SyslogResource{
	}
}

func (s *SyslogResource) Register(router *gin.RouterGroup) {

}

func (s *SyslogResource) GetParseErrorList(c *gin.Context) {
	var u dto.ParseErrorParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.ErrorTrans(c, err)
		return
	}

	var list []*entity.SysLog
	var total int64
	err := gorm.Instance().Scopes(utils.Paginate(u.Current, u.PageSize)).
		Find(&list).Count(&total).Order("id desc").Error

	if err != nil {
		ginutil.JSONServerError(c, err)
	}

	ginutil.JSONListPages(c, list, u.Current, u.PageSize, total)
}
