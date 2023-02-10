package api

import (
	"errors"
	"ms-go-blog/common"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/service"
	"ms-go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*APIHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Failure(w, errors.New("不识别此请求路径"))
		return
	}

	post, err := dao.GetPostDetailById(pid)
	if err != nil {
		common.Failure(w, err)
		return
	}

	common.Success(w, post)
}

func (*APIHandler) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {

	// 获取用户ID 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		common.Failure(w, errors.New("登录已过期"))
		return
	}
	uid := claims.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		// save
		params := common.GetRequestJsonParam(r)
		cid := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cid)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)

		pType := int(postType)
		// userId

		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			10,
			pType,
			time.Now(),
			time.Now(),
		}

		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// update

		params := common.GetRequestJsonParam(r)
		cid := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cid)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		pType := int(postType)
		// userId

		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			10,
			pType,
			time.Now(),
			time.Now(),
		}

		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*APIHandler) SearchPost(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)

	common.Success(w, searchResp)
}
