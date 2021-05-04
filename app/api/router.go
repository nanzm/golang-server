package api

import (
	"dora/app/middleware"
	"dora/config"
	"dora/pkg/ginutil"
	"dora/pkg/utils"

	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var uptime = time.Now()

func Register(g *gin.Engine, conf *config.Conf) {
	// cors
	g.Use(middleware.CORSMiddleware(), middleware.RateLimitMiddleware())

	g.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name":    "dora",
			"build":   config.Build,
			"compile": config.Compile,
			"version": config.Version,
			"uptime":  utils.TimeFromNow(uptime),
			"now":     utils.CurrentTime(),
		})
	})

	// pong
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
	g.HEAD("/ping", pingHandler)
	g.GET("/ping", pingHandler)

	pubRouter := g.Group("/")
	setupResource(pubRouter, NewPublicResource())

	apiRouter := g.Group("/api")
	setupResource(apiRouter,
		NewDashboardResource(),
		NewIssueResource(),
		NewOrganizationResource(),
		NewProjectResource(),
		NewSyslogResource(),
		NewUserResource(),
	)

}

type Resource interface {
	Register(router *gin.RouterGroup)
}

func setupResource(rg *gin.RouterGroup, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rg)
	}
}

func errorTrans(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ginutil.JSONFail(c, http.StatusBadRequest, err.Error())
	}
	ginutil.JSONValidatorFail(c, http.StatusBadRequest, removeTopStruct(errs.Translate(ginutil.Trans)))
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
