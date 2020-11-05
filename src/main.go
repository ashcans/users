package main

import (
	"github.com/ashcans/repos/src/middleware"
	"github.com/ashcans/users/src/config"
	"github.com/ashcans/users/src/controller"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := config.LoadDefault()

	// 分配路由
	router := httprouter.New()
	//用户
	router.POST("/user/:id", middleware.JSONType(controller.CreateUser))
	router.GET("/user/:id", middleware.JSONType(controller.GetUser))
	router.PATCH("/user/:id", middleware.JSONType(controller.UpdateUser))
	router.DELETE("/user/:id", middleware.JSONType(controller.DeleteUser))
	//角色
	router.POST("/role/:id", middleware.JSONType(controller.CreateRole))
	router.GET("/role/:id", middleware.JSONType(controller.GetRole))
	router.PATCH("/role/:id", middleware.JSONType(controller.UpdateRole))
	router.DELETE("/role/:id", middleware.JSONType(controller.DeleteRole))

	// 连接数据库
	log.Fatal(http.ListenAndServe(config.ListenAddress, router))
}
