package routes

import (
	"GIN/gin_demo/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	//加载HTML文件
	r.LoadHTMLFiles("./templates/user/register.html", "./templates/user/login.html", "./templates/user/index.html",
		"./templates/spot/spotinfo.html", "./templates/spot/spotDisplay.html",
	)

	//景区Spot路由组
	spotGroup := r.Group("spot")
	{
		//增加景点spot
		spotGroup.POST("/spots", controller.CreateSpot)
		//删除某个spot
		spotGroup.DELETE("/spots/:id", controller.DeleteSpotById)
		//查看所有的spot
		spotGroup.GET("/spots", controller.GetSpotList)
		//修改某个spot
		spotGroup.PUT("/spots/:id", controller.UpdateSpot)

		//查看景区排序的所有spot
		spotGroup.GET("/spots/score", controller.GetSpotOrder)
		//查看景区地址排序的所有spot
		spotGroup.GET("/spots/score/1", controller.GetSpotOrder1)

		//景区信息录入
		spotGroup.GET("/info", controller.InfoInterface)
		spotGroup.POST("/info", controller.Info)
	}

	//用户User路由组
	userGroup := r.Group("user")
	{
		//增加用户User
		userGroup.POST("/users", controller.CreateUser)
		//删除某个User
		userGroup.DELETE("/users/:id", controller.DeleteUserById)
		//查看所有的User
		userGroup.GET("/users", controller.GetUserList)
		//修改某个User
		userGroup.PUT("/users/:id", controller.UpdateUser)

		//用户注册
		userGroup.GET("/register", controller.RegisterInterface)
		userGroup.POST("/register", controller.Register)

		//用户登录
		userGroup.GET("/login", controller.LoginInterface)
		userGroup.POST("/login", controller.Login)
	}
	return r
}
