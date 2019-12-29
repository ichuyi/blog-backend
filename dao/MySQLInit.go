package dao

import (
	"blog-backend/model"
	"blog-backend/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"xorm.io/core"
)

var blogEngine *xorm.Engine

func init() {
	connect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", util.ConfigInfo.MySQL.User, util.ConfigInfo.MySQL.Password, util.ConfigInfo.MySQL.Host, util.ConfigInfo.MySQL.Port, util.ConfigInfo.MySQL.Database)
	var err error
	blogEngine, err = xorm.NewEngine("mysql", connect)
	if err != nil {
		log.Fatalf(err.Error())
	}
	blogEngine.ShowSQL(true)
	blogEngine.Logger().SetLevel(core.LOG_DEBUG)
	err = blogEngine.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Info("success to connect to MySQL,connect info :", connect)
	err = blogEngine.Sync2(new(model.Label), new(model.Todo), new(model.User), new(model.Post), new(model.FileInfo))
	if err != nil {
		log.Errorf(err.Error())
	}
}
