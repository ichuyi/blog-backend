package dao

import (
	"blog-backend/model"
	log "github.com/sirupsen/logrus"
)

func GetFileByCondition(file *model.FileInfo) ([]model.FileInfo, error) {
	list := make([]model.FileInfo, 0)
	err := blogEngine.Table("file_info").Desc("id").Find(&list, file)
	if err != nil {
		log.Errorf("get file info error: %s", err.Error())
		return nil, err
	} else {
		return list, nil
	}
}
func InsertFile(name string, key string, userId int, contentType string) (*model.FileInfo, error) {
	file := model.FileInfo{
		Filename:    name,
		Key:         key,
		UserId:      userId,
		ContentType: contentType,
	}
	_, err := blogEngine.Table("file_info").Insert(&file)
	if err != nil {
		log.Errorf("insert file info error: %s", err.Error())
		return nil, err
	}
	return &file, nil
}
func DeleteFile(id int) error {
	file := model.FileInfo{
		Id: id,
	}
	_, err := blogEngine.Table("file_info").Delete(&file)
	if err != nil {
		log.Errorf("delete file error: %s, id is : %d", err.Error(), id)
		return err
	}
	return nil
}
