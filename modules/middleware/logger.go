package middleware

import (
	"dora/pkg/utils/logx"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinZap() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		var pre string
		if c.Writer.Status() <= 400 {
			pre = color.HiGreenString("%v", c.Writer.Status())

		} else if c.Writer.Status() <= 500 {
			pre = color.HiYellowString("%v", c.Writer.Status())

		} else {
			pre = color.HiRedString("%v", c.Writer.Status())
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logx.Zap.Error(pre,
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("error", e),
					zap.Duration("latency", latency),
				)
			}
		} else {
			logx.Zap.Info(pre,
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.Duration("latency", latency),
			)
		}
	}
}
