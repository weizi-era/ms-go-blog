package service

import (
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	// 查询所有的分类
	categorys, _ := dao.GetAllCategory()

	// 查询所有的文章 进行月份整理
	posts, _ := dao.GetAllPost()

	pigeonholeMap := make(map[string][]models.Post)

	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}
