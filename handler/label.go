package handler

import (
	"blog-backend/dao"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddLabelReq struct {
	Name string `json:"name"`
}

func addLabel(ctx *gin.Context) {
	req := AddLabelReq{}
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
	if label, err := dao.InsertLabel(req.Name, id); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
	} else {
		util.OKResponse(ctx, label)
	}
}

func getLabelList(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	id, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	if list, err := dao.GetAllLabel(id); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
	} else {
		util.OKResponse(ctx, list)
	}
}
