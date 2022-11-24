package router

import (
	"RocoGuide/router/attrs"
	"RocoGuide/router/group"
	"RocoGuide/router/news"
	"RocoGuide/router/skill"
	"RocoGuide/router/spirit"
	"github.com/gin-gonic/gin"
)

type API func(*gin.RouterGroup)

var apis []API

func registerAPI(api ...API) {
	apis = append(apis, api...)
}

func InitAPI(port string) {
	registerAPI(news.LoadNews)
	registerAPI(spirit.LoadSpirit)
	registerAPI(group.LoadSpiritGroup)
	registerAPI(attrs.LoadAttrs)
	registerAPI(skill.LoadSkill)
	engin := gin.Default()
	g := engin.Group("/api")
	for _, api := range apis {
		api(g)
	}
	engin.Run(":" + port)
}
