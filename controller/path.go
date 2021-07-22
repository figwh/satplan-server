package controller

import (
	"net/http"
	"satplan/common"
	entity "satplan/entity"
	"satplan/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetPathBySenId(c *gin.Context) {
	senId := c.Query("senid")
	start, err := strconv.ParseInt(c.Query("start"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetRespResult(int(common.FAILED),
			"start time 格式不对", nil, 0))
		return
	}
	stop, err := strconv.ParseInt(c.Query("stop"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetRespResult(int(common.FAILED),
			"stop time格式不对", nil, 0))
		return
	}
	path := service.GetSenPath(senId, start, stop)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"query success", path, len(*path)))
}

func GetPathPlan(c *gin.Context) {
	var planPara entity.PlanPara
	c.ShouldBindBodyWith(&planPara, binding.JSON)

	path := service.GetPathPlan(planPara)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"query success", path, len(*path)))
}
