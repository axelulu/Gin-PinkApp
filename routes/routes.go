package routes

import (
	"net/http"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
	"web_app/settings"

	"github.com/gin-contrib/pprof"

	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(time.Duration(settings.Conf.RateLimitTime)*time.Second, settings.Conf.RateLimitNum))

	// swagger文档地址
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	// 应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/categoryList", controller.CategoryListHandle)
		v1.GET("/postList", controller.PostCategoryListHandle)
		v1.GET("/post/:id", controller.PostByIdHandle)
		v1.GET("/home", controller.HomeHandle)
		v1.GET("/user/:id", controller.UserHandle)
		v1.GET("/profile", controller.ProfileHandle)
		v1.GET("/ranking", controller.RankingHandle)
		v1.GET("/search", controller.SearchHandle)
	}

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"msg": "OK",
		})
	})

	// pprof性能测试
	pprof.Register(r)

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"msg": "404",
		})
	})

	return r
}
