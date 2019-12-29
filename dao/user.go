package dao

import (
	"blog-backend/model"
	log "github.com/sirupsen/logrus"
)

func InsertUser(username string, password string) (user model.User, err error) {
	user = model.User{
		Username: username,
		Password: password,
	}
	_, err = blogEngine.Table("user").Insert(&user)
	if err != nil {
		log.Errorf("insert user error: %s", err.Error())
	}
	return
}
func GetUserByCondition(user model.User) (*model.User, error) {
	has, err := blogEngine.Table("user").Get(&user)
	if err != nil {
		log.Errorf("get user error: %s", err.Error())
	}
	if err != nil || !has {
		return nil, err
	}
	return &user, nil
}
