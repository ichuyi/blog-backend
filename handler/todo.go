package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddTodoReq struct {
	Content string `json:"content"`
}

func addTodo(ctx *gin.Context) {
	req := AddTodoReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	value, _ := ctx.Get("user_id")
	id, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	todo, err := dao.InsertTodo(req.Content, id)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, *todo)
}

type UpdateTodoReq struct {
	Id     int `json:"id"`
	Finish int `json:"finish"`
}

func updateTodo(ctx *gin.Context) {
	req := UpdateTodoReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	todo := model.Todo{
		Id:     req.Id,
		Finish: req.Finish,
	}
	if err := dao.UpdateTodoById(&todo); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, nil)
}

type DeleteTodoReq struct {
	Id int `json:"id"`
}

func deleteTodo(ctx *gin.Context) {
	req := DeleteTodoReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	if err := dao.DeleteTodoById(req.Id); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, nil)
}

func getTodoList(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	id, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	todo := model.Todo{
		UserId: id,
	}
	if list, err := dao.GetTodoList(&todo); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)

	} else {
		util.OKResponse(ctx, list)
	}
	return
}
