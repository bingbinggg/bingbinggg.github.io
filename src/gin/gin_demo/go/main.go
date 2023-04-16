package main

import (
	"GIN/gin_demo/dao"
	"GIN/gin_demo/entity"
	"GIN/gin_demo/routes"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8080的项目
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
