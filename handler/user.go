package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SignReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(ctx *gin.Context) {
	req := SignReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	user, err := dao.InsertUser(req.Username, req.Password)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, user)
}
func signIn(ctx *gin.Context) {
	req := SignReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	if user, err := dao.GetUserByCondition(model.User{
		Username: req.Username,
	}); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
	} else if user == nil {
		util.FailedResponse(ctx, util.UserNotExist, util.UserNotExistMsg)
	} else if user.Password != req.Password {
		util.FailedResponse(ctx, util.PasswordError, util.PasswordErrorMsg)
	} else {
		token, err := util.GetToken(user)
		if err != nil {
			log.Errorf("get token error: %s", err.Error())
			util.FailedResponse(ctx, util.OtherError, util.OtherErrorMsg)
			return
		}
		ctx.SetCookie("token", token, util.MaxAge, "/", util.Domain, false, true)
		util.OKResponse(ctx, *user)
	}
	return
}

func signOut(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", util.Domain, false, true)
	util.OKResponse(ctx, nil)
}
