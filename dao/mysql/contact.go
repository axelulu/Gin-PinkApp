package mysql

import (
	"database/sql"
	"web_app/models"
)

func GetContactListByUserId(uid int64) (contact []*models.ContactDetail, err error) {
	sqlStr := `select user_id, send_id, update_time from contact where user_id=?`
	err = db.Select(&contact, sqlStr, uid)
	return
}

func GetContactListByUserIdSendId(uid int64, sid int64) (contact []*models.ContactDetail, err error) {
	sqlStr := `select user_id, send_id, update_time from contact where (user_id=? and send_id=?) or (send_id=? and user_id=?)`
	err = db.Select(&contact, sqlStr, uid, sid, sid, uid)
	return
}

func InsertContactItem(uid int64, sid int64) (res sql.Result, err error) {
	tx, err := db.Begin()
	sqlStr := `insert into contact (user_id, send_id) values(?,?)`
	res, err = tx.Exec(sqlStr, uid, sid)
	sqlStr2 := `insert into contact (user_id, send_id) values(?,?)`
	res, err = tx.Exec(sqlStr2, sid, uid)
	if err != nil {
		err = tx.Rollback()
	}
	err = tx.Commit()
	return
}

func GetChatListByUserId(uid int64, sid int64) (chat []*models.Message, err error) {
	sqlStr := `select user_id, send_id, cmd, media, content, update_time from chat where user_id=? or send_id=? order by update_time`
	err = db.Select(&chat, sqlStr, uid, sid)
	return
}

func InsertChatItem(msg models.Message) (res sql.Result, err error) {
	sqlStr2 := `insert into chat (user_id, send_id, cmd, content, media) values(?,?,?,?,?)`
	res, err = db.Exec(sqlStr2, msg.UserId, msg.SendId, msg.Cmd, msg.Content, msg.Media)
	return
}
