package dao

import (
	"blog-backend/model"
	"blog-backend/util"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func InsertPost(title string, content string, labels []int, userId int) (*model.Post, error) {
	post := model.Post{
		Title:   title,
		Content: content,
		Label:   labels,
		UserId:  userId,
	}
	_, err := blogEngine.Table("post").Insert(&post)
	if err != nil {
		log.Errorf("insert post into mysql error: %s", err.Error())
		return nil, err
	}
	collection := BlogClient.Database(util.ConfigInfo.Mongo.DataBase).Collection("post")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = collection.InsertOne(ctx, bson.M{"_id": post.Id, "content": post.Content})
	if err != nil {
		log.Errorf("insert post into mongo error: %s", err.Error())
		return nil, err
	}
	return &post, nil
}
func GetPostByCondition(post *model.Post) ([]*model.Post, error) {
	list := make([]*model.Post, 0)
	err := blogEngine.Table("post").Desc("update_time").Find(&list, post)
	if err != nil {
		log.Errorf("get post from mysql error: %s", err.Error())
		return nil, err
	}
	for _, p := range list {
		res := model.Post{}
		collection := BlogClient.Database(util.ConfigInfo.Mongo.DataBase).Collection("post")
		err = collection.FindOne(context.TODO(), bson.M{"_id": p.Id}).Decode(&res)
		if err != nil {
			log.Errorf("get post from mongo error: %s", err.Error())
			return nil, err
		}
		p.Content = res.Content
	}
	return list, nil
}
