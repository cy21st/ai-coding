package routes

import (
	"ai-go/handlers"
	"ai-go/middleware"
	"ai-go/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set trusted proxies
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 使用CORS中间件
	r.Use(middleware.Cors())

	// 处理404路由
	r.NoRoute(func(c *gin.Context) {
		utils.NotFound(c, "Route not found", nil)
	})

	// Public routes (不需要认证的路由)
	// 用户登录接口
	r.POST("/api/login", handlers.Login)
	// 用户登出接口
	r.POST("/api/logout", handlers.Logout)
	// 测试接口
	r.GET("/api/test", handlers.Test)

	// Protected routes (需要JWT认证的路由)
	authorized := r.Group("/api")
	authorized.Use(middleware.JWTAuth())
	{
		// Dashboard statistics (仪表盘统计)
		// 获取系统统计数据（用户、事件、属性、关联数量）
		authorized.GET("/statistics", handlers.GetStatistics)

		// User management routes (用户管理)
		// 创建新用户
		authorized.POST("/users", handlers.CreateUser)
		// 获取用户列表
		authorized.GET("/users", handlers.GetUserList)
		// 获取当前用户信息
		authorized.GET("/users/info", handlers.GetUserInfo)
		// 更新用户信息
		authorized.POST("/users/:id", handlers.UpdateUser)
		// 删除用户（软删除）
		authorized.POST("/users/:id/delete", handlers.DeleteUser)

		// Event management routes (事件管理)
		// 创建新事件
		authorized.POST("/events", handlers.CreateEvent)
		// 更新事件信息
		authorized.POST("/events/:id", handlers.UpdateEvent)
		// 删除事件（软删除）
		authorized.POST("/events/:id/delete", handlers.DeleteEvent)
		// 获取事件列表（支持分页和搜索）
		authorized.GET("/events", handlers.GetEventList)
		// 获取所有事件及其属性（支持分页和搜索）
		authorized.GET("/events/all", handlers.GetAllEventsWithAttributes)
		// 获取单个事件详细信息
		authorized.GET("/events/:id", handlers.GetEventInfo)
		// 获取事件的属性列表
		authorized.GET("/events/:id/attributes", handlers.GetEventAttributes)

		// Attribute management routes (属性管理)
		// 创建新属性
		authorized.POST("/attributes", handlers.CreateAttribute)
		// 更新属性信息
		authorized.POST("/attributes/:id", handlers.UpdateAttribute)
		// 删除属性（软删除）
		authorized.POST("/attributes/:id/delete", handlers.DeleteAttribute)
		// 获取属性列表（支持分页和搜索）
		authorized.GET("/attributes", handlers.GetAttributeList)
		// 获取单个属性详细信息
		authorized.GET("/attributes/:id", handlers.GetAttributeInfo)
		// 获取属性数据类型列表
		authorized.GET("/attribute-types", handlers.GetAttributeTypes)

		// Relation management routes (关联关系管理)
		// 创建新的事件-属性关联关系
		authorized.POST("/relations", handlers.CreateRelation)
		// 获取关联关系列表（支持分页）
		authorized.GET("/relations", handlers.GetRelationList)
		// 删除关联关系（软删除）
		authorized.POST("/relations/:id/delete", handlers.DeleteRelation)
	}

	return r
}
