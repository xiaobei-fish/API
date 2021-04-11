package route

import (
	"NewTest3/controller"
	"NewTest3/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine)*gin.Engine{
	r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())			 //使用中间件
	r.POST("/api/register", controller.Register)                     //注册
	r.POST("/api/login", controller.Login)                           //登录
	r.GET("/api/info", middleware.AuthMiddleware(), controller.Info) //用户信息
	r.POST("/api/add", middleware.AuthMiddleware(), controller.Add)  //增加事件
	r.POST("/api/alt", middleware.AuthMiddleware(), controller.Alt)	 //修改事件
	r.POST("/api/que", middleware.AuthMiddleware(), controller.Que)	 //查询事件
	r.POST("/api/del", middleware.AuthMiddleware(), controller.Del)  //删除事件

	return r
}
