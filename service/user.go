package service

import (
	"errors"
	"satplan/common"
	"satplan/dao/db"
	"satplan/entity"
	"strings"
)

func GetUser(email string, password string) (*entity.User, error) {
	sysUser := db.FindSysUserByEmail(email)
	if sysUser.Id == 0 {
		return nil, errors.New("找不到用户：" + email)
	} else if !strings.EqualFold(common.DecryptString(sysUser.Password), password) {
		return nil, errors.New("用户名或密码错误")
	} else {
		return &entity.User{
			UserName: sysUser.Name,
			Email:    email,
		}, nil
	}
}

func FindAllUsers() *[]entity.SysUser {

	return nil
}
