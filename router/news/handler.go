package news

import (
	"RocoGuide/base"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getNewsRequest struct {
	Page   *int `json:"page" form:"page" binding:"required" example:"1"`
	Amount *int `json:"amount" form:"amount" binding:"required" example:"20"`
}

func getNews(c *gin.Context) {
	var request getNewsRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}

	list, t := model.GetNewsList(*request.Page, *request.Amount)
	c.JSON(http.StatusOK, base.ResponseListEntity{
		Code:  http.StatusOK,
		Msg:   "success",
		Data:  &list,
		Total: t,
	})
}
