package handler

import (
	"blog-backend/dao"
	"blog-backend/model"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
)

type AddPostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	Labels  []int  `json:"labels"`
}

func addPost(ctx *gin.Context) {
	req := AddPostReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	post, err := dao.InsertPost(req.Title, req.Content, req.Labels, req.UserId)
	if err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	util.OKResponse(ctx, *post)
}

type GetPostListReq struct {
	UserId int `json:"user_id"`
}

type PostDetail struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	CreateTime model.JsonTime `json:"create_time"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Label      []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"label"`
}

func getPostList(ctx *gin.Context) {
	req := GetPostListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.FailedResponse(ctx, ParaError, ParaErrorMsg)
		return
	}
	post := &model.Post{
		UserId: req.UserId,
	}
	postList, err := dao.GetPostByCondition(post)
	if err != nil {
		util.FailedResponse(ctx, SQLError, SQLErrorMsg)
		return
	}
	result := make([]PostDetail, len(postList))
	for i, p := range postList {
		result[i] = PostDetail{
			Id:         p.Id,
			UserId:     p.UserId,
			CreateTime: p.CreateTime,
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
				util.FailedResponse(ctx, SQLError, SQLErrorMsg)
				return
			}
			result[i].Label[j].Name = la.Name
			result[i].Label[j].Id = la.Id
		}
	}
	util.OKResponse(ctx, result)

}
