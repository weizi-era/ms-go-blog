package views

import (
	"errors"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}

	postDetailRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("Detail查询出错"))
		return
	}

	detail.WriteData(w, postDetailRes)
}
