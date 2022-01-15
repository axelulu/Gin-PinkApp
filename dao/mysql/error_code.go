package mysql

import "errors"

var (
	ErrorValidateCode       = errors.New("验证码错误")
	ErrorUserNotLogin       = errors.New("用户未登录")
	ErrorUserExist          = errors.New("用户已存在")
	ErrorEmailExist         = errors.New("邮箱已存在")
	ErrorUserNotExist       = errors.New("用户不存在")
	ErrorInvalidPassword    = errors.New("用户名或密码错误")
	ErrorHashBcryptPassword = errors.New("用户密码加密错误")
	ErrorCatEmpty           = errors.New("分类为空")
	ErrorUserMeta           = errors.New("用户信息错误")
	ErrorPostMeta           = errors.New("文章信息错误")
	ErrorUserChat           = errors.New("用户聊天消息错误")
	ErrorContactExist       = errors.New("该对话已存在")
	ErrorUserFollowed       = errors.New("已经关注此用户")
	ErrorUserUnFollowed     = errors.New("已经取消关注此用户")
	ErrorPostLiked          = errors.New("已经喜欢了该文章")
	ErrorPostUnLiked        = errors.New("已经取消喜欢了该文章")
	ErrorPostStared         = errors.New("已经收藏了该文章")
	ErrorPostUnStared       = errors.New("取消收藏了该文章")
	ErrorPostCoined         = errors.New("已经投币了该文章")
)
