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

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logx.Error(e)
			}
		} else {
			var pre string
			if c.Writer.Status() <= 400 {
				pre = color.HiGreenString("%v", c.Writer.Status())

			} else if c.Writer.Status() <= 500 {
				pre = color.HiYellowString("%v", c.Writer.Status())

			} else {
				pre = color.HiRedString("%v", c.Writer.Status())
			}

			logx.Zap.Info(pre,
				//zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				//zap.String("user-agent", c.Request.UserAgent()),
				//zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
		}
	}
}
