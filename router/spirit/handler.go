package spirit

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getSpiritListRequest struct {
	Page     *int   `json:"page" form:"page" binding:"required"`
	Amount   *int   `json:"amount" form:"amount" binding:"required"`
	Keywords string `json:"keywords" form:"keywords"`
}

func getSpiritList(c *gin.Context) {
	var request getSpiritListRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: 400,
			Msg:  "请求参数错误",
		})
		return
	}
	list, total := model.GetSpiritList(*request.Page, *request.Amount, request.Keywords)
	c.JSON(200, base.ResponseListEntity{
		Code:  200,
		Msg:   "成功",
		Data:  &list,
		Total: total,
	})
}
