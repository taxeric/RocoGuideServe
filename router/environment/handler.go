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
	Id        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name" binding:"required"`
	Introduce string `json:"introduce" form:"introduce" bind:"required"`
	Effects   string `json:"effects" form:"effects" binding:"required"`
	Type      int    `json:"type" form:"type" binding:"required"`
	Icon      string `json:"icon" form:"icon" binding:"required"`
}

func insertEnvironment(c *gin.Context) {
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
	e.Name = request.Name
	e.Introduce = request.Introduce
	e.Effects = request.Effects
	e.Type = request.Type
	e.Icon = request.Icon
	var id = model.InsertEnvironment(e)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: id,
	})
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
	if err != nil || request.Id <= 0 {
		c.JSON(http.StatusBadRequest, base.ResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var e entity.SkillEnvironment
	e.Id = request.Id
	e.Name = request.Name
	e.Introduce = request.Introduce
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
