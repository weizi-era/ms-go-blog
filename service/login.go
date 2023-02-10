package service

import (
	"errors"
	"log"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginResponse, error) {

	passwd = utils.Md5Crypt(passwd, "mszlu")

	log.Println(passwd)

	user := dao.GetUserInfo(userName, passwd)

	if user == nil {
		return nil, errors.New("账号密码不正确")
	}

	// 生成token jwt技术进行生成令牌
	token, err := utils.Award(&user.Uid)
	if err != nil {
		log.Println(err)
	}

	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var lr = &models.LoginResponse{
		token,
		userInfo,
	}

	return lr, nil
}
