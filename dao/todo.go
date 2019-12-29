package dao

import (
	"blog-backend/model"
	log "github.com/sirupsen/logrus"
)

func InsertTodo(content string, userId int) (todo *model.Todo, err error) {
	todo = &model.Todo{
		UserId:  userId,
		Content: content,
	}
	_, err = blogEngine.Table("todo").Insert(todo)
	if err != nil {
		log.Errorf("insert todo error: %s", err.Error())
	}
	return
}
func DeleteTodoById(todoId int) error {
	_, err := blogEngine.Table("todo").Delete(&model.Todo{
		Id: todoId,
	})
	if err != nil {
		log.Errorf("delete todo error: %s", err.Error())
	}
	return err
}
func UpdateTodoById(todo *model.Todo) error {
	_, err := blogEngine.Table("todo").Where("id= ? ", todo.Id).Update(todo)
	if err != nil {
		log.Errorf("update todo error: %s", err.Error())
	}
	return err
}
func GetTodoList(todo *model.Todo) (list []model.Todo, err error) {
	list = make([]model.Todo, 0)
	err = blogEngine.Table("todo").Find(&list, todo)
	if err != nil {
		log.Errorf("get todo error: %s", err.Error())
	}
	return
}
