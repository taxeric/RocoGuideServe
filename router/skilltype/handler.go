package skilltype

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllSkillType(c *gin.Context) {
	var list = model.GetAllSkillType()
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Data:  list,
		Total: len(list),
	})
}
