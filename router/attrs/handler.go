package attrs

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllAttrs(c *gin.Context) {
	var list = model.GetAllAttrs()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Data:  list,
		Total: len(list),
	})
}
