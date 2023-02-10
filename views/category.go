package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {

	categoryTemplate := common.Template.Category

	path := r.URL.Path
	cidStr := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		log.Println("")
		categoryTemplate.WriteError(w, errors.New("路径不匹配"))
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员！！"))
		return
	}

	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)

	// 每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostByCategoryId(cid, page, pageSize)
	categoryTemplate.WriteData(w, categoryResponse)

}
