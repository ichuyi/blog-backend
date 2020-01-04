package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddPostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Labels  []int  `json:"labels"`
}

func addPost(ctx *gin.Context) {
	req := AddPostReq{}
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
	post, err := dao.InsertPost(req.Title, req.Content, req.Labels, id)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, *post)
}

type PostDetail struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	CreateTime model.JsonTime `json:"create_time"`
	UpdateTime model.JsonTime `json:"update_time"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Label      []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"label"`
}

func getPostList(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	id, err := strconv.Atoi(value.(string))
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	post := &model.Post{
		UserId: id,
	}
	postList, err := dao.GetPostByCondition(post)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	result := make([]PostDetail, len(postList))
	for i, p := range postList {
		result[i] = PostDetail{
			Id:         p.Id,
			UserId:     p.UserId,
			CreateTime: p.CreateTime,
			UpdateTime: p.UpdateTime,
			Title:      p.Title,
			Content:    p.Content,
		}
		result[i].Label = make([]struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}, len(p.Label))
		for j, l := range p.Label {
			la, err := dao.GetLabelById(l)
			if err != nil || la == nil {
				util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
				return
			}
			result[i].Label[j].Name = la.Name
			result[i].Label[j].Id = la.Id
		}
	}
	util.OKResponse(ctx, result)

}
func getPost(ctx *gin.Context) {
	v := ctx.Query("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		util.FailedResponse(ctx, util.ParaError, util.ParaErrorMsg)
		return
	}
	post := &model.Post{
		Id: id,
	}
	postList, err := dao.GetPostByCondition(post)
	if err != nil {
		util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
		return
	}
	if len(postList) == 0 {
		util.FailedResponse(ctx, util.PostNotExist, util.PostNotExistMsg)
		return
	}
	result := make([]PostDetail, len(postList))
	for i, p := range postList {
		result[i] = PostDetail{
			Id:         p.Id,
			UserId:     p.UserId,
			CreateTime: p.CreateTime,
			UpdateTime: p.UpdateTime,
			Title:      p.Title,
			Content:    p.Content,
		}
		result[i].Label = make([]struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}, len(p.Label))
		for j, l := range p.Label {
			la, err := dao.GetLabelById(l)
			if err != nil || la == nil {
				util.FailedResponse(ctx, util.SQLError, util.SQLErrorMsg)
				return
			}
			result[i].Label[j].Name = la.Name
			result[i].Label[j].Id = la.Id
		}
	}
	util.OKResponse(ctx, result[0])
}
