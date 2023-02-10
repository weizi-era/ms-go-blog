package service

import (
	"html/template"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {

	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)

	var postMores []models.PostMore

	for _, post := range posts {

		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}

		postMores = append(postMores, postMore)

	}

	total := dao.GetAllPostCount()
	var pages []int
	pagesCount := (total-1)/10 + 1

	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}

	return hr, nil
}
