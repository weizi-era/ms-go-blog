package api

import (
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)

func (*APIHandler) Login(w http.ResponseWriter, r *http.Request) {

	// 接收用户名和密码 返回对应的json数据
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)

	loginResponse, err := service.Login(username, passwd)
	if err != nil {
		common.Failure(w, err)
		return
	}
	common.Success(w, loginResponse)
}
