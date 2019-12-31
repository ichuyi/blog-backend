package handler

import (
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func authorization(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		util.FailedResponse(ctx, util.Unauthorized, util.UnauthorizedMsg)
		ctx.Abort()
		return
	}
	if payload := util.ParseMyToken(token); payload == nil {
		util.FailedResponse(ctx, util.Unauthorized, util.UnauthorizedMsg)
		ctx.Abort()
	} else {
		ctx.Set("user_id", strconv.Itoa(payload.UserId))
		ctx.Next()
	}
}
