package model

type User struct {
	Id       int    `json:"id" xorm:"pk autoincr int"`
	Username string `json:"username" xorm:"notnull unique"`
	Password string `json:"password" xorm:"notnull"`
}
