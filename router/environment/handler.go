package environment

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type skillEnvironment struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Effects string `json:"effects"`
	Type    int    `json:"type"`
	Icon    string `json:"icon"`
}

func getAllEnvironment(c *gin.Context) {
	var list = model.GetAllEnvironment()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  list,
		Total: len(list),
	})
}

func updateEnvironment(c *gin.Context) {
	var request skillEnvironment
	var err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var e entity.SkillEnvironment
	e.Id = request.Id
	e.Name = request.Name
	e.Effects = request.Effects
	e.Type = request.Type
	e.Icon = request.Icon
	var id = model.UpdateEnvironment(e)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}
