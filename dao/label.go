package dao

import (
	"blog-backend/model"
	log "github.com/sirupsen/logrus"
)

func InsertLabel(tag string, userId int) (label *model.Label, err error) {
	label = &model.Label{
		Name:   tag,
		UserId: userId,
	}
	_, err = blogEngine.Table("label").Insert(label)
	if err != nil {
		log.Errorf("insert label error: %s", err.Error())
	}
	return
}
func GetAllLabel(userId int) (labelList []model.Label, err error) {
	labelList = make([]model.Label, 0)
	if err = blogEngine.Table("label").Desc("id").Where("user_id = ?", userId).Find(&labelList); err != nil {
		log.Errorf("get label list error: %s", err.Error())
	}
	return
}
func GetLabelById(id int) (*model.Label, error) {
	label := &model.Label{
		Id: id,
	}
	has, err := blogEngine.Table("label").Get(label)
	if err != nil {
		log.Errorf("get label by id error: %s", err.Error())
		return nil, err
	} else if !has {
		return nil, nil
	} else {
		return label, nil
	}
}
