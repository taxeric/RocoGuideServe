package spirit

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getSpiritListRequest struct {
	Page     *int   `json:"page" form:"page" binding:"required"`
	Amount   *int   `json:"amount" form:"amount" binding:"required"`
	Keywords string `json:"keywords" form:"keywords"`
}

type spiritDetailsRequest struct {
	Id                    *int64   `json:"id" form:"id"`
	Avatar                *string  `json:"avatar" form:"avatar" binding:"required"`
	Number                *int     `json:"number" form:"number" binding:"required"`
	Name                  *string  `json:"name" form:"name" binding:"required"`
	Description           *string  `json:"description" form:"description" binding:"required"`
	PrimaryAttributesId   *int     `json:"primaryAttributes" form:"primaryAttributes"`
	SecondaryAttributesId *int     `json:"secondaryAttributes" form:"secondaryAttributes"`
	RacePower             *int     `json:"racePower" form:"racePower" binding:"required"`
	RaceAttack            *int     `json:"raceAttack" form:"raceAttack" binding:"required"`
	RaceDefense           *int     `json:"raceDefense" form:"raceDefense" binding:"required"`
	RaceMagicAttack       *int     `json:"raceMagicAttack" form:"raceMagicAttack" binding:"required"`
	RaceMagicDefense      *int     `json:"raceMagicDefense" form:"raceMagicDefense" binding:"required"`
	RaceSpeed             *int     `json:"raceSpeed" form:"raceSpeed" binding:"required"`
	GroupId               *int     `json:"group" form:"group" binding:"required"`
	Height                *float32 `json:"height" form:"height" binding:"required"`
	Weight                *float32 `json:"weight" form:"weight" binding:"required"`
	Hobby                 *string  `json:"hobby" form:"hobby" binding:"required"`
	Skill                 *[]int   `json:"skills" form:"skills"`
}

type spiritByIdRequest struct {
	Id *int `json:"id" form:"id" binding:"required"`
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
		Msg:   "success",
		Data:  &list,
		Total: total,
	})
}

func getSpiritDetailsById(c *gin.Context) {
	var request spiritByIdRequest
	err := c.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: 400,
			Msg:  "请求参数错误",
		})
		return
	}
	data, err := model.GetSpiritDetailsById(*request.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: 400,
			Msg:  "请求参数错误",
		})
		return
	}
	c.JSON(200, base.ResponseEntity{
		Code: 200,
		Msg:  "success",
		Data: &data,
	})
}

func insertSpirit(c *gin.Context) {
	var request spiritDetailsRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: 400,
			Msg:  "请求参数错误",
		})
		return
	}
	var spirit entity.Spirit
	spirit.Avatar = *request.Avatar
	spirit.Number = *request.Number
	spirit.Name = *request.Name
	spirit.Description = *request.Description
	spirit.PrimaryAttributes.Id = request.PrimaryAttributesId
	spirit.SecondaryAttributes.Id = request.SecondaryAttributesId
	spirit.RacePower = *request.RacePower
	spirit.RaceAttack = *request.RaceAttack
	spirit.RaceDefense = *request.RaceDefense
	spirit.RaceMagicAttack = *request.RaceMagicAttack
	spirit.RaceMagicDefense = *request.RaceMagicDefense
	spirit.RaceSpeed = *request.RaceSpeed
	spirit.Group.Id = *request.GroupId
	spirit.Height = *request.Height
	spirit.Weight = *request.Weight
	spirit.Hobby = *request.Hobby
	spirit.Skills = make([]entity.Skill, 0)
	if request.Skill != nil {
		for _, v := range *request.Skill {
			var s = new(entity.Skill)
			s.Id = int64(v)
			spirit.Skills = append(spirit.Skills, *s)
		}
	}
	id := model.InsertSpirit(&spirit)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}

func updateSpirit(c *gin.Context) {
	var request spiritDetailsRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.BadResponseEntity{
			Code: 400,
			Msg:  "请求参数错误",
		})
		return
	}
	var spirit entity.Spirit
	spirit.Avatar = *request.Avatar
	spirit.Number = *request.Number
	spirit.Name = *request.Name
	spirit.Description = *request.Description
	spirit.PrimaryAttributes.Id = request.PrimaryAttributesId
	spirit.SecondaryAttributes.Id = request.SecondaryAttributesId
	spirit.RacePower = *request.RacePower
	spirit.RaceAttack = *request.RaceAttack
	spirit.RaceDefense = *request.RaceDefense
	spirit.RaceMagicAttack = *request.RaceMagicAttack
	spirit.RaceMagicDefense = *request.RaceMagicDefense
	spirit.RaceSpeed = *request.RaceSpeed
	spirit.Group.Id = *request.GroupId
	spirit.Height = *request.Height
	spirit.Weight = *request.Weight
	spirit.Hobby = *request.Hobby
	id := model.UpdateSpirit(&spirit)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}
