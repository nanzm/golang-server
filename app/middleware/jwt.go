package middleware

import (
	"dora/pkg/ginutil"
	"dora/pkg/jwtutil"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			ginutil.JSONFail(c, 401, "未登录")
			return
		}

		mc, err := jwtutil.ParseToken(token)
		if err != nil {
			ginutil.JSONFail(c, 401, "无效的 token")
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("uid", mc.Id)
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
