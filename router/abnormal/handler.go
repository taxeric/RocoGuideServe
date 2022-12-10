package abnormal

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type abnormal struct {
	Id        int64   `json:"id"  form:"id"`
	Name      string  `json:"name" form:"name" binding:"required"`
	Introduce string  `json:"introduce" form:"introduce" binding:"required"`
	Icon      *string `json:"icon" form:"icon"`
}

func insertAbnormal(c *gin.Context) {
	var request abnormal
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var state = entity.AbnormalState{
		Name:      request.Name,
		Introduce: request.Introduce,
	}
	if request.Icon == nil {
		state.Icon = sql.NullString{Valid: false}
	} else {
		state.Icon = sql.NullString{Valid: true, String: *request.Icon}
	}
	var id = model.InsertAbnormalState(state)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: id,
	})
}

func getAllAbnormal(c *gin.Context) {
	var list = model.GetAllAbnormalState()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  list,
		Total: len(list),
	})
}

func updateAbnormal(c *gin.Context) {
	var request abnormal
	err := c.ShouldBind(&request)
	if err != nil || request.Id <= 0 {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var state = entity.AbnormalStateResponse{
		Id:        request.Id,
		Name:      request.Name,
		Introduce: request.Introduce,
	}
	if request.Icon == nil {
		state.Icon = ""
	} else {
		state.Icon = *request.Icon
	}
	var id = model.UpdateAbnormalState(state)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: id,
	})
}
