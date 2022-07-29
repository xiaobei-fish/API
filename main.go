package main

import (
	"NewTest3/common"
	"NewTest3/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/viper"
	"os"
)

func main(){
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r=route.CollectRoute(r)
	port := viper.GetString("server.port")
	if port!="" {
		r.Run(":" + port)
	}else {
		r.Run("3000") //默认端口号：8080
	}
}

func InitConfig() {
	Dir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(Dir + "/config")
	fmt.Println(Dir)
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
}