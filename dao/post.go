package dao

import (
	"log"
	"ms-go-blog/models"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {

	var posts []models.Post

	page = (page - 1) * pageSize

	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetSearchPost(condition string) ([]models.Post, error) {
	var posts []models.Post

	rows, err := DB.Query("select * from blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// GetAllPost 查询所有文章
func GetAllPost() ([]models.Post, error) {

	var posts []models.Post

	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetAllPostCount() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}

func GetAllPostCountByCategoryId(categoryId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id=?", categoryId)
	_ = rows.Scan(&count)
	return
}

func GetPostPageByCategoryId(cid, page, pageSize int) ([]models.Post, error) {
	var posts []models.Post

	pages := (page - 1) * pageSize

	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,?", cid, pages, pageSize)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostDetailById(pid int) (models.Post, error) {

	var post models.Post

	row := DB.QueryRow("select * from blog_post where pid=?", pid)
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt)

	if err != nil {
		return post, err
	}

	return post, nil
}

func SavePost(post *models.Post) {

	result, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId,
		post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}

	pid, _ := result.LastInsertId()

	post.Pid = int(pid)

}

func UpdatePost(post *models.Post) {

	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title, post.Content, post.Markdown, post.CategoryId,
		post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
	}
}
