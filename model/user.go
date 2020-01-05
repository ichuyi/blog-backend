package model

type User struct {
	Id              int    `json:"id" xorm:"pk autoincr int"`
	Username        string `json:"username" xorm:"notnull unique"`
	Password        string `json:"password" xorm:"notnull"`
	Description     string `json:"description" xorm:"varchar(255)"`
	NeteasePhone    string `json:"netease_phone" xorm:"varchar(255)"`
	NeteasePassword string `json:"netease_password" xorm:"varchar(255)"`
}
