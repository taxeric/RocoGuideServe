package series

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type series struct {
	Id   int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

func insertSeries(c *gin.Context) {
	var request series
	var err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var e entity.SpiritSeries
	e.Name = request.Name
	var id = model.InsertSeries(e)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}

func getAllSeries(c *gin.Context) {
	var list = model.GetAllSeries()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  list,
		Total: len(list),
	})
}

func updateSeries(c *gin.Context) {
	var request series
	var err = c.ShouldBind(&request)
	if err != nil || request.Id <= 0 {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var e entity.SpiritSeries
	e.Name = request.Name
	e.Id = request.Id
	var id = model.UpdateSeries(&e)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}
