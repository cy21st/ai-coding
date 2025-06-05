package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求的中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许的来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// 允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		// 允许携带认证信息（cookies, authorization headers）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 预检请求的缓存时间
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
