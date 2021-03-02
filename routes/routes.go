package routes

import (
	"rocky/controllers"
	_ "rocky/docs" // 千万不要忘了导入把你上一步生成的docs

	"net/http"
	"rocky/logger"
	"rocky/middlewares"

	gs "github.com/swaggo/gin-swagger"

	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controllers.GetPostListHandler2)
		// 投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
