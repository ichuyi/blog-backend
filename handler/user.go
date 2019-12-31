package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SignReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(ctx *gin.Context) {
	req := SignReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	user, err := dao.InsertUser(req.Username, req.Password)
	if err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, user)
}
func signIn(ctx *gin.Context) {
	req := SignReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	if user, err := dao.GetUserByCondition(model.User{
		Username: req.Username,
	}); err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
	} else if user == nil {
		util.FailedResponse(ctx, UserNotExist, UserNotExistMsg)
	} else if user.Password != req.Password {
		util.FailedResponse(ctx, PasswordError, PasswordErrorMsg)
	} else {
		ctx.SetCookie("current_user_id",strconv.Itoa(user.Id),MaxAge,"/",Domain,false,true)
		util.OKResponse(ctx, *user)
	}
	return
}

func signOut(ctx *gin.Context)  {
	ctx.SetCookie("current_user_id","",-1,"/",Domain,false,true)
	util.OKResponse(ctx,nil)
}
