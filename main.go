package main

import (
	"os"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//TODO:在router.go中设置路由规则
//TODO:补全代码
func main() {
	r := gin.Default()
	InitConfig()
	router.Init(r)
	common.InitDB()
	go oss.Init()
	port := viper.GetString("server.port")
	panic(r.Run(port))
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	// 这里不需要加文件名
	//viper.AddConfigPath(workDir + "/config.yml")
	viper.AddConfigPath(workDir+"/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
