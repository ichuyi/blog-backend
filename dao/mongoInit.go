package dao

import (
	"blog-backend/util"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var BlogClient *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	BlogClient, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", util.ConfigInfo.Mongo.Host, util.ConfigInfo.Mongo.Port)))
	if err != nil {
		log.Fatalf("create mongo client error: %s", err.Error())
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = BlogClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("connect to mongo error: %s", err.Error())
	}
}
