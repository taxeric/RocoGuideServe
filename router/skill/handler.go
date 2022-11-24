package skill

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetSkillPageRequest struct {
	Page     *int   `json:"page" form:"page" binding:"required" `
	Amount   *int   `json:"amount" form:"amount" binding:"required" `
	Keywords string `json:"keywords" form:"keywords"`
}

func getSkillPages(c *gin.Context) {
	var request GetSkillPageRequest
	err := c.ShouldBind(&request)
	if err != nil {
		panic("failed")
	}
	list, total := model.GetSkillByName(*request.Page, *request.Amount, request.Keywords)
	c.JSON(
		http.StatusOK, base.ResponseListEntity{
			Code:  http.StatusOK,
			Data:  &list,
			Total: total,
		})

}
