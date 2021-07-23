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
		return nil, errors.New("cannot find user: " + email)
	} else if !strings.EqualFold(common.DecryptString(sysUser.Password), password) {
		return nil, errors.New("wrong username or password")
	} else {
		return &entity.User{
			UserName: sysUser.UserName,
			Email:    email,
		}, nil
	}
}

func GetAllUsers() *[]entity.SysUser {
	return db.FindAllSysUsers()
}
