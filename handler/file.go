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
)

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

func uploadFile(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	userId, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		log.Errorf("get userid error: %s", err.Error())
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		log.Errorf("parse file error: %s", err.Error())
		return
	}
	files := form.File["files"]
	result := make([]model.FileInfo, len(files))
	for i, f := range files {
		name := f.Filename
		contentType := f.Header.Get("Content-Type")
		t, err := f.Open()
		if err != nil {
			util.FailedResponse(ctx, util.OpenFileError, util.OpenFileErrorMsg)
			log.Errorf("open file error: %s", err.Error())
			return
		}
		content, err := ioutil.ReadAll(t)
		log.Infof("file size is %d, read size is %d", f.Size, len(content))
		if err != nil {
			util.FailedResponse(ctx, util.ReadFileError, util.ReadFileErrorMsg)
			log.Errorf("read file error: %s", err.Error())
			return
		}
		key := md5File(content)
		n, err := writeToLocal(content, key)
		if err != nil {
			util.FailedResponse(ctx, util.WriteFileError, util.WriteFileErrorMsg)
			log.Errorf("write into file error: %s", err.Error())
			return
		}
		if n != len(content) {
			util.FailedResponse(ctx, util.WriteFileError, util.WriteFileErrorMsg)
			log.Errorf("content size is: %d,actually write size is: %d", len(content), n)
			return
		}
		p, err := dao.InsertFile(name, key, userId, contentType)
		if err != nil {
			util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
			return
		}
		result[i] = *p
	}
	util.OKResponse(ctx, result)
}

func getFileList(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	id, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	file := model.FileInfo{
		UserId: id,
	}
	list, err := dao.GetFileByCondition(&file)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, list)
}
func getFile(ctx *gin.Context) {
	value := ctx.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		log.Errorf("parse string to int error: %s", err.Error())
		return
	}
	file := model.FileInfo{
		Id: id,
	}
	list, err := dao.GetFileByCondition(&file)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	if len(list) == 0 {
		util.FailedResponse(ctx, util.FileNotExist, util.FileNotExistMsg)
		return
	}
	file = list[0]
	content, err := ReadFromLocal(file.Key)
	if err != nil {
		util.FailedResponse(ctx, util.ReadFileError, util.ReadFileErrorMsg)
		log.Errorf("read file error: %s", err.Error())
		return
	}
	ctx.Header("content-disposition", `attachment; filename=`+file.Filename)
	ctx.Data(http.StatusOK, file.ContentType, content)
}
func deleteFile(ctx *gin.Context) {
	value := ctx.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		log.Errorf("parse string to int error: %s", err.Error())
		return
	}
	file := model.FileInfo{
		Id: id,
	}
	list, err := dao.GetFileByCondition(&file)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	if len(list) == 0 {
		util.FailedResponse(ctx, util.FileNotExist, util.FileNotExistMsg)
		return
	}
	file.Key = list[0].Key
	if err := dao.DeleteFile(id); err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	file.Id = 0
	list, err = dao.GetFileByCondition(&file)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	if len(list) == 0 {
		if err = os.Remove(util.UploadPath + file.Key); err != nil {
			util.FailedResponse(ctx, util.DeleteFileError, util.DeleteFileErrorMsg)
			log.Errorf("remove local file error: %s", err.Error())
			return
		}
	}
	util.OKResponse(ctx, nil)
}
