package model

type FileInfo struct {
	Id       int    `json:"id" xorm:"pk autoincr"`
	Filename string `json:"filename"`
	Key      string `json:"key" xorm:"notnull varchar(32)"`
	UserId   int    `json:"user_id" xorm:"notnull"`
}
