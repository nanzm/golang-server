package ginutil

import (
	"context"
	"dora/pkg/utils/logx"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type Resource interface {
	Register(router *gin.RouterGroup)
}

func SetupResource(rg *gin.RouterGroup, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rg)
	}
}


func GetOrigin(c *gin.Context) string {
	scheme := "http"
	host := c.Request.Host
	forwardedHost := c.GetHeader("X-Forwarded-Host")
	if forwardedHost != "" {
		host = forwardedHost
	}
	forwardedProto := c.GetHeader("X-Forwarded-Proto")
	if forwardedProto == "https" {
		scheme = forwardedProto
	}

	return fmt.Sprintf("%s://%s", scheme, host)
}

// 优雅关闭
func GracefulShutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()
	now := time.Now()

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeout); err != nil {
		logx.Fatal("Server Shutdown:", err)
	}
	logx.Infof("Server exiting，关闭耗时：", time.Since(now))
}
