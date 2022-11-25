package skilltype

import "github.com/gin-gonic/gin"

func LoadSkillType(g *gin.RouterGroup) {
	group := g.Group("/skilltype")
	loadSkillTypeV1(group)
}

func loadSkillTypeV1(g *gin.RouterGroup) {
	group := g.Group("/v1")
	group.GET("/page", getAllSkillType)
}
