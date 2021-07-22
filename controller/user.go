package controller

import (
	"net/http"
	"satplan/common"
	"satplan/service"

	entity "satplan/entity"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userInfo, _ := service.GetUserInfo(c)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询用户成功", []entity.CurrentUserInfo{userInfo}, 1))
}

func GetAllUsers(c *gin.Context) {
	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), "权限不足", nil, 0))
		return
	}

	users := service.FindAllUsers()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询用户成功", *users, len(*users)))
}
