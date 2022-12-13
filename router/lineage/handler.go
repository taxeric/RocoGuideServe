package lineage

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type lineageRequest struct {
	Id        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name" binding:"required"`
	Introduce string `json:"introduce" form:"introduce" binding:"required"`
	Icon      string `json:"icon" form:"icon" binding:"required"`
}

func getAllLineage(c *gin.Context) {
	var list = model.GetAllLineage()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  list,
		Total: len(list),
	})
}

func insertLineage(c *gin.Context) {
	var request lineageRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var lineage entity.Lineage
	lineage.Name = request.Name
	lineage.Introduce = request.Introduce
	lineage.Icon = request.Icon
	var id = model.InsertLineage(lineage)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: id,
	})
}

func updateLineage(c *gin.Context) {
	var request lineageRequest
	err := c.ShouldBind(&request)
	if err != nil || request.Id <= 0 {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var lineage = entity.Lineage{
		Id:        request.Id,
		Name:      request.Name,
		Introduce: request.Introduce,
		Icon:      request.Icon,
	}
	var id = model.UpdateLineage(lineage)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: id,
	})
}
