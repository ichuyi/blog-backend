package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"time"
)

type AddTodoReq struct {
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

func addTodo(ctx *gin.Context) {
	req := AddTodoReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	todo, err := dao.InsertTodo(req.Content, req.UserId)
	if err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, *todo)
}

type UpdateTodoReq struct {
	Id int `json:"id"`
}

func finishTodo(ctx *gin.Context) {
	req := UpdateTodoReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	todo := model.Todo{
		Id:         req.Id,
		Finish:     FINISH,
		FinishTime: model.JsonTime(time.Now()),
	}
	if err := dao.UpdateTodoById(&todo); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, nil)
}

func deleteTodo(ctx *gin.Context) {
	req := UpdateTodoReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	if err := dao.DeleteTodoById(req.Id); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, nil)
}

type TodoListReq struct {
	Id int `json:"id"`
}

func getTodoList(ctx *gin.Context) {
	req := TodoListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	todo := model.Todo{
		UserId: req.Id,
	}
	if list, err := dao.GetTodoList(&todo); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)

	} else {
		util.OKResponse(ctx, list)
	}
	return
}
