package model

type Post struct {
	Id         int      `json:"id" xorm:"pk autoincr int" bson:"_id"`
	UserId     int      `json:"user_id" xorm:"notnull"`
	CreateTime JsonTime `json:"create_time" xorm:"created"`
	Title      string   `json:"title"`
	Content    string   `json:"content" xorm:"-" bson:"content"`
	Label      []int    `json:"label" xorm:"varchar(255)"`
}
