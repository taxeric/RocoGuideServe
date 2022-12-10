package router

import (
	"RocoGuide/middleware"
	"RocoGuide/router/abnormal"
	"RocoGuide/router/attrs"
	"RocoGuide/router/environment"
	"RocoGuide/router/group"
	"RocoGuide/router/news"
	"RocoGuide/router/series"
	"RocoGuide/router/skill"
	"RocoGuide/router/skilltype"
	"RocoGuide/router/spirit"
	"github.com/gin-gonic/gin"
	"net/http"
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
	registerAPI(skilltype.LoadSkillType)
	registerAPI(series.LoadSeries)
	registerAPI(environment.LoadSkillEnvironment)
	registerAPI(abnormal.LoadAbnormalState)
	engin := gin.Default()
	g := engin.Group("/api")
	g.Use(middleware.Cors())
	for _, api := range apis {
		api(g)
	}
	g.StaticFS("/res", http.Dir("./res"))
	engin.Run(":" + port)
}
