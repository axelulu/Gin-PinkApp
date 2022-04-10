package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"pinkacg/logic"
	"pinkacg/models"
	"strconv"
)

// HomeHandle 首页文章列表
func HomeHandle(c *gin.Context) {
	// 获取参数
	p := new(models.Home)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Home with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	home, err := logic.HomeList(p)
	if err != nil {
		zap.L().Error("logic.HomeList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, home)
}

// RecommendHandle 推荐文章列表
func RecommendHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostRecommendList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Home with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	// 逻辑处理
	home, err := logic.RecommendList(p, uid)
	if err != nil {
		zap.L().Error("logic.RecommendList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, home)
}

// RecommendPostListHandle 推荐文章列表
func RecommendPostListHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostRecommendList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostCategoryList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	uid, err := getCurrentUserID(c)
	// 逻辑处理
	post, err := logic.RecommendPostList(p, uid)
	if err != nil {
		zap.L().Error("logic.PostCategoryList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// PostCategoryListHandle 分类文章列表
func PostCategoryListHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostCategoryList)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostCategoryList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	post, err := logic.PostCategoryList(p)
	if err != nil {
		zap.L().Error("logic.PostCategoryList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// PostListByIdsHandle 根据用户ids获取文章列表
func PostListByIdsHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostListByIds)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostListByIds with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	post, err := logic.PostListByIds(p)
	if err != nil {
		zap.L().Error("logic.PostListByIds failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// PostByIdHandle 根据用户ID获取文章
func PostByIdHandle(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	post, err := logic.PostById(pid, uid)
	if err != nil {
		zap.L().Error("logic.PostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// PostViewByIdHandle 增加文章观看数
func PostViewByIdHandle(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 逻辑处理
	err = logic.PostViewById(pid)
	if err != nil {
		zap.L().Error("logic.PostViewById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostCreateHandle 文章创建
func PostCreateHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostPublish)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostPublish with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	exec, err := logic.PostPublish(p, uid)
	if err != nil {
		zap.L().Error("logic.PostPublish failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, exec)
}

// RankingHandle 文章排行榜
func RankingHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostRankingList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostRankingList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	post, err := logic.PostRanking(p)
	if err != nil {
		zap.L().Error("logic.PostRanking failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// DynamicHandle 动态文章
func DynamicHandle(c *gin.Context) {
	// 获取参数
	p := new(models.PostDynamicList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.PostDynamicList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	post, err := logic.PostDynamic(p, uid)
	if err != nil {
		zap.L().Error("logic.PostDynamic failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

// UserPostHandle 用户文章列表
func UserPostHandle(c *gin.Context) {
	// 获取参数
	p := new(models.UserPost)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.UserPost with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.GetUserPost(p, uid)
	if err != nil {
		zap.L().Error("logic.GetUserPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
