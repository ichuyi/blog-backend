package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type MongoConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DataBase string `json:"data_base"`
}
type ServiceConfig struct {
	Port string `json:"port"`
}
type MatrixConfig struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}
type SocketConfig struct {
	Raddr      string `json:"raddr"`
	Delimiter  string `json:"delimiter"`
	MaxMessage int    `json:"max_message"`
}
type UploadConfig struct {
	UploadFolder string `json:"upload_folder"`
}
type Config struct {
	MySQL   MySQLConfig   `json:"mysql"`
	Service ServiceConfig `json:"service"`
	Matrix  MatrixConfig  `json:"matrix"`
	Socket  SocketConfig  `json:"socket"`
	Mongo   MongoConfig   `json:"mongo"`
	Upload  UploadConfig  `json:"upload"`
}

var (
	execBasePath = getCurrentDirectory()
	//execBasePath = "/Users/mhiwy/Documents/blog-backend"
	configPath = execBasePath + "/config.json"
	ConfigInfo Config
	UploadPath string
)

func init() {
	parse()
	UploadPath = fmt.Sprintf("%s/%s/", execBasePath, ConfigInfo.Upload.UploadFolder)
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
func parse() {
	conf, err := os.Open(configPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = json.NewDecoder(conf).Decode(&ConfigInfo)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
