package middleware

import (
	"dora/internal/datasource/redis"
	"dora/pkg/utils/ginutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"strconv"
	"time"
)

func RateLimitMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		rdb := redis.Instance()

		ip := c.ClientIP()

		limiter := redis_rate.NewLimiter(rdb)
		res, err := limiter.Allow(c.Request.Context(), ip, redis_rate.PerMinute(100))
		if err != nil {
			ginutil.JSONBadRequest(c, err)
			return
		}

		c.Header("RateLimit-Remaining", strconv.Itoa(res.Remaining))

		if res.Allowed == 0 {
			// rate limited.
			ms := int(res.RetryAfter / time.Millisecond)
			retry := strconv.Itoa(ms) + "ms"

			c.Header("RateLimit-RetryAfter", retry)
			ginutil.JSONFail(c, 400, fmt.Sprintf("ip: %v 访问太快了， %v后重试", ip, retry))
			return
		}

		c.Next()
	}
}
