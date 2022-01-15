package mysql

import (
	"database/sql"
	"pinkacg/models"
)

func GetContactListByUserId(uid int64) (contact []*models.ContactDetail, err error) {
	sqlStr := `select user_id, send_id, update_time from contacts where user_id=?`
	err = db.Select(&contact, sqlStr, uid)
	return
}

func GetContactItemByUserId(uid int64, sid int64) (contact *models.ContactDetail, err error) {
	contact = new(models.ContactDetail)
	sqlStr := `select user_id, send_id, update_time from contacts where user_id=? and send_id=?`
	err = db.Get(contact, sqlStr, uid, sid)
	return
}

func GetContactListByUserIdSendId(uid int64, sid int64) (contact []*models.ContactDetail, err error) {
	sqlStr := `select user_id, send_id, update_time from contacts where (user_id=? and send_id=?) or (send_id=? and user_id=?)`
	err = db.Select(&contact, sqlStr, uid, sid, sid, uid)
	return
}

func InsertContactItem(uid int64, sid int64) (res sql.Result, err error) {
	tx, err := db.Begin()
	sqlStr := `insert into contacts (user_id, send_id,update_time,create_time) values(?,?,NOW(),NOW())`
	res, err = tx.Exec(sqlStr, uid, sid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	res, err = tx.Exec(sqlStr, sid, uid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	return
}

func GetChatListByUserId(uid int64, chatList *models.ChatList) (chat []*models.Message, err error) {
	sqlStr := `select user_id, send_id, cmd, media, pic, content, update_time from chats where (user_id=? and send_id=?) or (user_id=? and send_id=?)order by update_time desc limit ?,?`
	offset := (chatList.Page - 1) * chatList.Size
	err = db.Select(&chat, sqlStr, uid, chatList.Sid, chatList.Sid, uid, offset, chatList.Size)
	return
}

func GetChatByUserId(uid int64, sendId int64) (chat *models.Message, err error) {
	chat = new(models.Message)
	sqlStr := `select user_id, send_id, cmd, media, pic, content, update_time from chats where (user_id=? and send_id=?) or (user_id=? and send_id=?)order by update_time desc`
	err = db.Get(chat, sqlStr, uid, sendId, sendId, uid)
	return
}

func InsertChatItem(msg models.Message) (res sql.Result, err error) {
	sqlStr2 := `insert into chats (user_id, send_id, cmd, pic, content, media,update_time,create_time) values(?,?,?,?,?,?,NOW(),NOW())`
	res, err = db.Exec(sqlStr2, msg.UserId, msg.SendId, msg.Cmd, msg.Pic, msg.Content, msg.Media)
	return
}
