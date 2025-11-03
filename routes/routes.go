package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/controllers"
	"github.com/lccxxo/blog/middlewares"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS - 允许所有源
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false, // 注意：AllowAllOrigins为true时，AllowCredentials必须为false
		MaxAge:           12 * time.Hour,
	}))

	// API路由组
	api := r.Group("/api")
	{
		// 公开的认证路由（不需要认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// 公开的文章接口
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticle)

		// 公开的分类和标签接口
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategory)
		api.GET("/tags", controllers.GetTags)
		api.GET("/tags/:id", controllers.GetTag)

		// 公开的评论接口
		api.GET("/articles/:id/comments", controllers.GetArticleComments)
		api.GET("/comments/:id", controllers.GetComment)
	}

	// 需要认证的路由（使用不同的组避免冲突）
	protected := api.Group("")
	protected.Use(middlewares.AuthMiddleware())
	{
		// 用户相关
		protected.GET("/profile", controllers.GetProfile)

		// 文章管理
		protected.POST("/articles", controllers.CreateArticle)
		protected.PUT("/articles/:id", controllers.UpdateArticle)
		protected.DELETE("/articles/:id", controllers.DeleteArticle)
		protected.GET("/my-articles", controllers.GetMyArticles)

		// 分类管理
		protected.POST("/categories", controllers.CreateCategory)
		protected.PUT("/categories/:id", controllers.UpdateCategory)
		protected.DELETE("/categories/:id", controllers.DeleteCategory)

		// 标签管理
		protected.POST("/tags", controllers.CreateTag)
		protected.PUT("/tags/:id", controllers.UpdateTag)
		protected.DELETE("/tags/:id", controllers.DeleteTag)

		// 评论管理
		protected.POST("/comments", controllers.CreateComment)
		protected.PUT("/comments/:id", controllers.UpdateComment)
		protected.DELETE("/comments/:id", controllers.DeleteComment)
		protected.GET("/my-comments", controllers.GetMyComments)

		// 上传管理
		protected.POST("/upload/image", controllers.UploadImage)
		protected.POST("/upload/video", controllers.UploadVideo)
		protected.DELETE("/upload/image", controllers.DeleteImage)
	}

	// 静态文件服务（提供上传的图片访问）
	r.Static("/uploads", "./uploads")

	return r
}
