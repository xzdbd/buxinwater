package models

type Userinfo struct {
	Id       int64  `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
	role     int64
}
