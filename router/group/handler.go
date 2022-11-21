package group

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllGroup(c *gin.Context) {
	list := model.GetGroupList()
	c.JSON(http.StatusOK, base.ResponseListEntity{Code: http.StatusOK, Total: len(list), Data: list})
}
