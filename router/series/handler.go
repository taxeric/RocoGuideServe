package series

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllSeries(c *gin.Context) {
	var list = model.GetAllSeries()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  list,
		Total: len(list),
	})
}
