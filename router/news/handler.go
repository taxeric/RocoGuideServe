package news

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getNewsRequest struct {
	Page   *int `json:"page" form:"page" binding:"required" example:"1"`
	Amount *int `json:"amount" form:"amount" binding:"required" example:"20"`
}

type insertNewsRequest struct {
	Id    int64  `json:"id" form:"id"`
	Title string `json:"title" form:"title" binding:"required"`
	Url   string `json:"url" form:"url" binding:"required"`
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

func insertNews(c *gin.Context) {
	var request insertNewsRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var id = model.InsertNews(request.Url, request.Title)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}

func updateNews(c *gin.Context) {
	var request insertNewsRequest
	err := c.ShouldBind(&request)
	if err != nil || request.Id <= 0 {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "请求参数错误",
		})
		return
	}
	var news = entity.News{
		Id:    request.Id,
		Url:   request.Url,
		Title: request.Title,
	}
	var id = model.UpdateNews(news)
	if id == 1 {
		c.JSON(http.StatusOK, base.ResponseEntity{
			Code: http.StatusOK,
			Msg:  "success",
			Data: request.Id,
		})
	} else {
		c.JSON(http.StatusOK, base.ResponseEntity{
			Code: http.StatusBadRequest,
			Msg:  "未找到指定id的情报",
		})
	}
}
