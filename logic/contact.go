package logic

import (
	"database/sql"
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetContactList(uid int64) (contacts map[string]interface{}, err error) {
	contacts = make(map[string]interface{}, 2)
	var contact []*models.ContactDetail
	contact, err = mysql.GetContactListByUserId(uid)
	for _, list := range contact {
		user, _ := mysql.GetUserById(list.SendId)
		list.SendUserMeta = user
	}
	contacts["list"] = contact
	contacts["total"] = len(contact)
	return
}

func AddContactList(uid int64, sid int64) (res sql.Result, err error) {
	var contact []*models.ContactDetail
	contact, err = mysql.GetContactListByUserIdSendId(uid, sid)
	if len(contact) > 0 || uid == sid {
		err = errors.New("该对话已存在")
	} else {
		res, err = mysql.InsertContactItem(uid, sid)
	}
	return
}

func GetChatList(uid int64) (chats map[string]interface{}, err error) {
	chats = make(map[string]interface{}, 2)
	var chat []*models.Message
	chat, err = mysql.GetChatListByUserId(uid)
	chats["list"] = chat
	chats["total"] = len(chat)
	return
}

func InsertChatItem(msg models.Message) (res sql.Result, err error) {
	res, err = mysql.InsertChatItem(msg)
	return
}
