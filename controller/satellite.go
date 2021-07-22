package controller

import (
	"net/http"
	"satplan/common"
	"satplan/service"

	log "github.com/sirupsen/logrus"

	entity "satplan/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddSatellite(c *gin.Context) {
	var tle entity.Tle
	c.ShouldBindBodyWith(&tle, binding.JSON)

	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), "权限不足", nil, 0))
		return
	}
	satelliteId, err := service.AddSatellite(&tle)
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "新建凭证成功", []int{satelliteId}, 1))
	}
}

func GetAllSatellites(c *gin.Context) {
	satellites := service.GetAllSatellites()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证成功", satellites, len(*satellites)))
}

func GetSatelliteById(c *gin.Context) {
	satId := c.Param("id")

	satellite, err := service.GetSatelliteById(satId)
	if err != nil {
		log.Debug("GetSatelliteById: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"查询凭证分组成功", []entity.Satellite{*satellite}, 1))
	}
}

func UpdateSatellite(c *gin.Context) {
	satId := c.Param("id")
	var satDTO entity.SatDTO
	c.ShouldBindBodyWith(&satDTO, binding.JSON)
	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), "权限不足", nil, 0))
		return
	}
	err := service.UpdateSatellite(satId, &satDTO)
	if err != nil {
		log.Debug("UpdateSatellite: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证分组成功", nil, 1))
	}
}

func DeleteSatellite(c *gin.Context) {
	satId := c.Param("id")
	err := service.DeleteSatelliteById(satId)
	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), "权限不足", nil, 0))
		return
	}
	if err != nil {
		log.Debug("DeleteSatellite: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "凭证删除成功", []bool{true}, 1))
	}
}
