package service

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func GetPostDetail(pid int) (*models.PostRes, error) {

	post, err := dao.GetPostDetailById(pid)
	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}

	return postRes, nil
}

func SearchPost(condition string) []models.SearchResp {
	var searchResps []models.SearchResp

	posts, err := dao.GetSearchPost(condition)
	if err != nil {
		log.Println(err)
	}

	for _, post := range posts {

		searchResps = append(searchResps, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}

	return searchResps

}
