package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var fileType = map[string]string{
	"jpg": "image/jpeg",
	"png": "image/png",
	"jpeg":"image/jpeg",
}

func md5File(content []byte) string {
	h := md5.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}
func writeToLocal(content []byte, key string) (int, error) {
	filename := util.UploadPath + key
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	return f.Write(content)
}
func ReadFromLocal(key string) ([]byte, error) {
	filename := util.UploadPath + key
	return ioutil.ReadFile(filename)
}

type UploadFileReq struct {
	UserId int `json:"user_id"`
}

func uploadFile(ctx *gin.Context) {
	req := UploadFileReq{
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		log.Errorf("parse file error: %s", err.Error())
		return
	}
	value,ok:=form.Value["userId"]
	if !ok{
		util.FailedResponse(ctx,ParaError,ParaErrorMsg)
		return
	}
	req.UserId,err=strconv.Atoi(value[0])
	if err!=nil{
		util.FailedResponse(ctx,ParaError,ParaErrorMsg)
		return
	}
	files := form.File["files"]
	result := make([]model.FileInfo, len(files))
	for i, f := range files {
		name := f.Filename
		info:=strings.Split(name,".")
		if len(info)<2{
			continue
		}
		if _,ok:=fileType[info[len(info)-1]];!ok{
			continue
		}
		t, err := f.Open()
		if err != nil {
			util.FailedResponse(ctx, OpenFileError, OpenFileErrorMsg)
			log.Errorf("open file error: %s", err.Error())
			return
		}
		content, err := ioutil.ReadAll(t)
		log.Infof("file size is %d, read size is %d", f.Size, len(content))
		if err != nil {
			util.FailedResponse(ctx, ReadFileError, ReadFileErrorMsg)
			log.Errorf("read file error: %s", err.Error())
			return
		}
		key := md5File(content)
		n, err := writeToLocal(content, key)
		if err != nil {
			util.FailedResponse(ctx, WriteFileError, WriteFileErrorMsg)
			log.Errorf("write into file error: %s", err.Error())
			return
		}
		if n != len(content) {
			util.FailedResponse(ctx, WriteFileError, WriteFileErrorMsg)
			log.Errorf("content size is: %d,actually write size is: %d", len(content), n)
			return
		}
		p, err := dao.InsertFile(name, key, req.UserId)
		if err != nil {
			util.FailedResponse(ctx, SQLError, SQLErrorMsg)
			return
		}
		result[i] = *p
	}
	util.OKResponse(ctx, result)
}


func getFileList(ctx *gin.Context)  {
	value:=ctx.Query("user_id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		util.FailedResponse(ctx,ParaError,ParaErrorMsg)
		return
	}
	file:=model.FileInfo{
		UserId:   id,
	}
	list,err:=dao.GetFileByCondition(&file)
	if err!=nil{
		util.FailedResponse(ctx,SQLError,SQLErrorMsg)
		return
	}
	util.OKResponse(ctx,list)
}
func getFile(ctx *gin.Context) {
	value := ctx.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		log.Errorf("parse string to int error: $s", err.Error())
		return
	}
	file := model.FileInfo{
		Id: id,
	}
	list, err := dao.GetFileByCondition(&file)
	if err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	if len(list) == 0 {
		util.FailedResponse(ctx, FileNotExist, FileNotExistMsg)
		return
	}
	file = list[0]
	content, err := ReadFromLocal(file.Key)
	if err != nil {
		util.FailedResponse(ctx, ReadFileError, ReadFileErrorMsg)
		log.Errorf("read file error: %s", err.Error())
		return
	}
	info := strings.Split(file.Filename, ".")
	if len(info) < 2 {
		util.FailedResponse(ctx, ReadFileError, ReadFileErrorMsg)
		log.Errorf("filename is: %s", file.Filename)
		return
	}
	ctx.Header("content-disposition", `attachment; filename=`+file.Filename)
	contentType, ok := fileType[info[len(info)-1]]
	if !ok {
		util.FailedResponse(ctx, ReadFileError, ReadFileErrorMsg)
		log.Errorf("file type is: %s", info[len(info)-1])
		return
	}
	ctx.Data(http.StatusOK, contentType, content)
}
func deleteFile(ctx *gin.Context)  {
	value := ctx.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		log.Errorf("parse string to int error: $s", err.Error())
		return
	}
	file:=model.FileInfo{
		Id:       id,
	}
	list,err:=dao.GetFileByCondition(&file)
	if err!=nil{
		util.FailedResponse(ctx,SQLError,SQLErrorMsg)
		return
	}
	if len(list)==0{
		util.FailedResponse(ctx,FileNotExist,FileNotExistMsg)
		return
	}
	file.Key=list[0].Key
	if err:=dao.DeleteFile(id);err!=nil{
		util.FailedResponse(ctx,SQLError,SQLErrorMsg)
		return
	}
	file.Id=0
	list,err=dao.GetFileByCondition(&file)
	if err!=nil{
		util.FailedResponse(ctx,SQLError,SQLErrorMsg)
		return
	}
	if len(list)==0{
		if err=os.Remove(util.UploadPath + file.Key);err!=nil{
			util.FailedResponse(ctx,DeleteFileError,DeleteFileErrorMsg)
			log.Errorf("remove local file error: %s",err.Error())
			return
		}
	}
	util.OKResponse(ctx,nil)
}