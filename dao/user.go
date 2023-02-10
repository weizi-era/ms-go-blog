package dao

import (
	"log"
	"ms-go-blog/models"
)

func GetUserNameById(userId int) string {

	var userName string

	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}

	_ = row.Scan(&userName)

	return userName
}

func GetUserInfo(userName, passwd string) *models.User {

	var user = &models.User{}

	row := DB.QueryRow("select * from blog_user where user_name=? and password=? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}

	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return user
}
