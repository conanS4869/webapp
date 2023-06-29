package controller

import (
	"html/template"
	"net/http"
	"webapp/dao"
	"webapp/model"
	"webapp/utils"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		//修改之后cookie发送浏览器
		http.SetCookie(w, cookie)
	}
	// 去首页
	GetPageBooksByPrice(w, r)
}
func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag {
		GetPageBooksByPrice(w, r)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		sess := &model.Session{
			SessionID: utils.CreateUUID(),
			UserName:  user.UserName,
			UserID:    user.ID,
		}
		dao.AddSession(sess)
		cookie := http.Cookie{
			Name:     "user",
			Value:    sess.SessionID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)

	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在!")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已存在!"))
	} else {
		w.Write([]byte("用户名可用!"))
	}
}
