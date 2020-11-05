package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DeleteUser 获取用户
func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// DeleteUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// DeleteUser 更新用户
func UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// DeleteUser 删除用户
func DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}
