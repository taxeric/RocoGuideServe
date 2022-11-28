package skill

import (
	"RocoGuide/base"
	"RocoGuide/entity"
	"RocoGuide/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetSkillPageRequest struct {
	Page     *int   `json:"page" form:"page" binding:"required" `
	Amount   *int   `json:"amount" form:"amount" binding:"required" `
	Keywords string `json:"keywords" form:"keywords"`
}

type skillDetailRequest struct {
	Id                *int    `json:"id" form:"id"`
	Name              *string `json:"name" form:"name" binding:"required"`
	Description       *string `json:"description" form:"description" binding:"required"`
	Value             *int    `json:"value" form:"value" binding:"required"`
	Amount            *int    `json:"amount" form:"amount" binding:"required"`
	Speed             *int    `json:"speed" form:"speed" binding:"required"`
	IsGenetic         *bool   `json:"isGenetic" form:"isGenetic" binding:"required"`
	AdditionalEffects *string `json:"additionalEffects" form:"additionalEffects" binding:"required"`
	IsBe              *bool   `json:"isBe" form:"isBe" binding:"required"`
	SkillTypeID       *int    `json:"skillTypeId" form:"skillTypeId" binding:"required"`
	AttributesID      *int    `json:"attributesId" form:"attributesId" binding:"required"`
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

func insertSkill(c *gin.Context) {
	var request skillDetailRequest
	err := c.ShouldBind(&request)
	if err != nil {
		panic("failed")
	}
	var skill entity.Skill
	skill.Name = *request.Name
	skill.Description = *request.Description
	skill.AdditionalEffects = *request.AdditionalEffects
	skill.Value = *request.Value
	skill.Amount = *request.Amount
	skill.Speed = *request.Speed
	skill.Attributes.Id = request.AttributesID
	skill.SkillType.Id = request.SkillTypeID
	skill.IsGenetic = *request.IsGenetic
	skill.IsBe = *request.IsBe
	id := model.InsertSkill(&skill)
	c.JSON(http.StatusOK, base.ResponseEntity{
		Code: http.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(id, 10),
	})
}
