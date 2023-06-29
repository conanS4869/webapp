package dao

import (
	"net/http"
	"webapp/model"
	"webapp/utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessID string) *model.Session {
	sqlStr := "select session_id,username,user_id  from sessions where session_id = ?"
	row := utils.Db.QueryRow(sqlStr, sessID)
	utils.Db.QueryRow(sqlStr, sessID)
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess
}
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
