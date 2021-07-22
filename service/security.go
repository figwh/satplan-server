package service

import (
	"errors"
	"fmt"
	"satplan/common"
	"satplan/dao/db"
	"satplan/entity"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetCurrentUserEmail(c *gin.Context) string {
	//get email
	claims := jwt.ExtractClaims(c)
	return fmt.Sprintf("%s", claims["mail"])
}

func GetUserInfo(c *gin.Context) (entity.CurrentUserInfo, error) {
	//get email
	email := GetCurrentUserEmail(c)

	//find user info by email
	sysUser := db.FindSysUserByEmail(email)
	if sysUser.Id == 0 {
		return entity.CurrentUserInfo{}, errors.New("找不到用户：" + email)
	}

	cui := entity.CurrentUserInfo{
		Id:   sysUser.Id,
		Name: sysUser.UserName,
	}

	return cui, nil
}

func IsPlatformAdmin(userID int) bool {
	return userID == int(common.PLATFORM_ADMIN)
}
