package routes

import (
	"net/http"
	"pinkacg/controller"
	"pinkacg/logger"
	"pinkacg/middlewares"
	"pinkacg/settings"
	"time"

	"github.com/gin-contrib/pprof"

	_ "pinkacg/docs" // 千万不要忘了导入把你上一步生成的docs

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
	v1.POST("/forgetPwd", controller.ForgetPwdHandler)

	// 版本更新
	v1.GET("/update", controller.UpdateHandle)

	// 脚本
	//v1.GET("/shell", controller.ShellHandle)
	//v1.GET("/getDouYinUrl", controller.GetDouYinUrlHandle)
	//v1.GET("/getDouYinPostUrl", controller.GetDouYinPostUrlHandle)

	// 邮箱发送
	v1.GET("/sendRegEmail", controller.SendRegEmailHandle)
	v1.GET("/sendForgetPwdEmail", controller.SendForgetPwdEmailHandle)

	// 应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// 聊天
		v1.GET("/contactList", controller.ContactListHandle)
		v1.GET("/contact/:id", controller.ContactItemHandle)
		v1.POST("/contact", controller.ContactCreateHandle)
		v1.GET("/chatList", controller.ChatListHandle)
		v1.GET("/chat", controller.Chat)

		// 分类
		v1.GET("/categoryList", controller.CategoryListHandle)

		// 文章
		v1.GET("/postList", controller.PostCategoryListHandle)
		v1.GET("/postListByIds", controller.PostListByIdsHandle)
		v1.POST("/post", controller.PostCreateHandle)
		v1.GET("/post/:id", controller.PostByIdHandle)

		// 文章查看
		v1.GET("/postView/:id", controller.PostViewByIdHandle)

		// 评论
		v1.POST("/comment", controller.CommentCreateHandle)
		v1.GET("/commentList", controller.CommentListHandle)

		// 首页
		v1.GET("/home", controller.HomeHandle)

		// 用户
		v1.GET("/user/:id", controller.UserHandle)
		v1.GET("/profile", controller.ProfileHandle)
		v1.GET("/userCenter/:id", controller.UserCenterHandle)
		v1.GET("/userPost", controller.UserPostHandle)
		v1.POST("/userInfoUpdate", controller.UserInfoUpdateHandle)
		v1.POST("/userPasswordUpdate", controller.UserPasswordUpdateHandle)
		v1.POST("/userEmailUpdate", controller.UserEmailUpdateHandle)
		// 邮箱发送
		v1.GET("/sendChangePwdEmail", controller.SendChangePwdEmailHandle)
		v1.GET("/sendChangeEmail", controller.SendChangeEmailHandle)

		// 关注
		v1.POST("/follow", controller.FollowHandle)
		v1.GET("/followStatus/:id", controller.FollowStatusHandle)
		v1.POST("/unFollow", controller.UnFollowHandle)
		v1.GET("/followList", controller.FollowListHandle)
		v1.GET("/fansList", controller.FansListHandle)

		// 喜欢
		v1.POST("/like", controller.LikeHandle)
		v1.POST("/unLike", controller.UnLikeHandle)

		// 硬币
		v1.POST("/coin", controller.CoinHandle)

		// 收藏
		v1.POST("/star", controller.StarHandle)
		v1.POST("/unStar", controller.UnStarHandle)

		// 排行
		v1.GET("/ranking", controller.RankingHandle)

		// 动态
		v1.GET("/dynamic", controller.DynamicHandle)

		// 搜索
		v1.GET("/search", controller.SearchHandle)

		// 上传
		v1.POST("/upload", controller.UploadHandle)

		// 日志申报
		v1.POST("/log", controller.LogHandler)
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
