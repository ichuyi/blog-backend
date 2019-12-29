package model

import (
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04") + `"`), nil
}

type Label struct {
	Id          int      `json:"id" xorm:"pk autoincr int"`
	UserId      int      `json:"user_id" xorm:"notnull"`
	Name        string   `json:"name" xorm:"notnull varchar(10)"`
	CreatedTime JsonTime `json:"created_time" xorm:"created"`
}
