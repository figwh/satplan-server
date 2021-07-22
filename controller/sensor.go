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

func AddSensor(c *gin.Context) {
	var sensorInDTO entity.SensorInDTO
	c.ShouldBindBodyWith(&sensorInDTO, binding.JSON)
	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), "权限不足", nil, 0))
		return
	}

	sensorId, err := service.AddSensor(&sensorInDTO)
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "新建凭证成功", []int{sensorId}, 1))
	}
}

func GetAllSensors(c *gin.Context) {
	sensors := service.GetAllSensors()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证成功", sensors, len(*sensors)))
}

func GetSensorBySatId(c *gin.Context) {
	satId := c.Param("satid")
	sensors, err := service.GetSensorBySatId(satId)
	if err != nil {
		log.Debug("GetSensorGroups: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证分组成功", []entity.Sensor{*sensors}, 1))
	}
}

func GetSensorById(c *gin.Context) {
	sensorId := c.Param("id")
	sensor := service.GetSensorById(sensorId)
	if sensor.Id == 0 {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证成功", nil, 0))
	} else {
		sensors := []entity.Sensor{*sensor}
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "查询凭证成功", &sensors, 1))
	}
}

func DeleteSensor(c *gin.Context) {
	sensorId := c.Param("id")
	err := service.DeleteSensorById(sensorId)
	if err != nil {
		log.Debug("DeleteSensor: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "凭证删除成功", []bool{true}, 1))
	}
}

func UpdateSensor(c *gin.Context) {
	sensorId := c.Param("id")
	var sensorInDTO entity.SensorInDTO
	c.ShouldBindBodyWith(&sensorInDTO, binding.JSON)

	err := service.UpdateSensor(sensorId, &sensorInDTO)
	if err != nil {
		log.Debug("UpdateSensor: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "更新凭证成功", []bool{true}, 1))
	}
}
