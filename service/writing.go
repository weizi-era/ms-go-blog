package service

import (
	"log"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func Writing() (wr models.WritingRes) {

	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}

	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	wr.Categorys = category

	return

}
