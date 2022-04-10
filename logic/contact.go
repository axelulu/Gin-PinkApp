package logic

import (
	"database/sql"
	"go.uber.org/zap"
	"pinkacg/dao/mysql"
	"pinkacg/models"
)

// GetContactList 获取联系人列表
func GetContactList(uid int64) (contacts map[string]interface{}, err error) {
	contacts = make(map[string]interface{}, 2)
	var contact []*models.ContactDetail
	contact, err = mysql.GetContactListByUserId(uid)
	if err != nil {
		return nil, err
	}
	for _, list := range contact {
		user, err := mysql.GetUserById(list.SendId)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed", zap.Int64("list.UserId", list.SendId), zap.Error(err))
			return nil, mysql.ErrorUserMeta
		}
		list.SendUserMeta = user
		chat, _ := mysql.GetChatByUserId(uid, list.SendId)
		list.Msg = chat
	}
	contacts["list"] = contact
	contacts["total"] = len(contact)
	return
}

// GetContactItem 获取单个联系人信息
func GetContactItem(uid int64, sid int64) (contact *models.ContactDetail, err error) {
	contact, err = mysql.GetContactItemByUserId(uid, sid)
	user, err := mysql.GetUserById(contact.SendId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("contact.UserId", contact.UserId), zap.Error(err))
		return nil, mysql.ErrorUserMeta
	}
	contact.SendUserMeta = user
	chat, _ := mysql.GetChatByUserId(uid, sid)
	contact.Msg = chat
	return
}

// GetChatList 获取聊天列表
func GetChatList(uid int64, chatList *models.ChatList) (chats map[string]interface{}, err error) {
	chats = make(map[string]interface{}, 2)
	var chat []*models.Message
	chat, err = mysql.GetChatListByUserId(uid, chatList)
	if err != nil {
		return nil, err
	}
	chats["list"] = chat
	chats["total"] = len(chat)
	return
}

// AddContactList 创建对话
func AddContactList(uid int64, sid int64) (res sql.Result, err error) {
	var contact []*models.ContactDetail
	contact, err = mysql.GetContactListByUserIdSendId(uid, sid)
	if err != nil {
		return nil, mysql.ErrorUserMeta
	}
	if len(contact) > 0 || uid == sid {
		return nil, mysql.ErrorContactExist
	} else {
		res, err = mysql.InsertContactItem(uid, sid)
	}
	return
}

func InsertChatItem(msg models.Message) (res sql.Result, err error) {
	res, err = mysql.InsertChatItem(msg)
	return
}
