package handler

import (
	"blog-backend/dao"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
)

type AddLabelReq struct {
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}

func addLabel(ctx *gin.Context) {
	req := AddLabelReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	if label, err := dao.InsertLabel(req.Name, req.UserId); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
	} else {
		util.OKResponse(ctx, label)
	}
}

type GetLabelListReq struct {
	UserId int `json:"user_id"`
}

func getLabelList(ctx *gin.Context) {
	req := GetLabelListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	if list, err := dao.GetAllLabel(req.UserId); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
	} else {
		util.OKResponse(ctx, list)
	}
}
