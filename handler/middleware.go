package handler

import (
	"blog-backend/util"
	"github.com/gin-gonic/gin"
log "github.com/sirupsen/logrus")

func authorization(ctx *gin.Context)  {
	if id,err:=ctx.Cookie("current_user_id");err!=nil{
		log.Errorf("unauthorized, error: %s",err.Error())
		util.FailedResponse(ctx,Unauthorized,UnauthorizedMsg)
		ctx.Abort()
	}else{
		ctx.Set("user_id",id)
		ctx.Next()
	}
}
