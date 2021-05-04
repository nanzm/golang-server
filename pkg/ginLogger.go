package pkg

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// GinZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func GinZap(logger *zap.Logger, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			var pre string
			if c.Writer.Status() < 400 {
				pre = color.HiGreenString("%v %v %v ", c.Writer.Status(), c.Request.Method, path)
			} else {
				pre = color.HiRedString("%v %v %v ", c.Writer.Status(), c.Request.Method, path)
			}

			logger.Info(pre,
				//zap.Int("status", c.Writer.Status()),
				//zap.String("method", c.Request.Method),
				//zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				//zap.String("user-agent", c.Request.UserAgent()),
				//zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
		}
	}
}
