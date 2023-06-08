package mall

import (
	"github.com/gin-gonic/gin"
	v1 "bookstore/api/v1"
	"bookstore/middleware"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	userRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate) //修改用户信息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)    //获取用户信息
		mallUserRouter.POST("/user/logout", mallUserApi.UserLogout)  //登出
	}
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
		userRouter.POST("/user/login", mallUserApi.UserLogin)       //登陆

	}

}
