package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func OKResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, CommonResponse{
		Code:    0,
		Message: "ok",
		Result:  data,
	})
}

func FailedResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, CommonResponse{
		Code:    code,
		Message: msg,
		Result:  nil,
	})
}
