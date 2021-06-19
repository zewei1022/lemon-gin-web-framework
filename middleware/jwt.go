package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/model/response"
	"github.com/zewei1022/lemon-gin-web-framework/utils"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(response.TokenEmpty, "缺少Token令牌", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(response.TokenInvalid, "Token令牌错误", c)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			response.FailWithDetailed(response.TokenExpired, "Token令牌已过期", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
