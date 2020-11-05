package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetRole 获取角色
func GetRole(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// CreateRole 创建角色
func CreateRole(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// UpdateRole 更新角色
func UpdateRole(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}

// DeleteRole 删除角色
func DeleteRole(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	println(id)
}
