package model

type Todo struct {
	Id         int      `json:"id" xorm:"int pk  autoincr"`
	UserId     int      `json:"user_id" xorm:"bigint"`
	CreateTime JsonTime `json:"create_time" xorm:"created"`
	FinishTime JsonTime `json:"finish_time" xorm:"datetime"`
	Finish     int      `json:"finish" xorm:"tinyint notnull"`
	Content    string   `json:"content" xorm:"text"`
}
