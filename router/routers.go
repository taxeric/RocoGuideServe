package router

import (
	"RocoGuide/router/news"
	news2 "RocoGuide/router/spirit"
	"github.com/gin-gonic/gin"
)

type API func(*gin.RouterGroup)

var apis []API

func registerAPI(api ...API) {
	apis = append(apis, api...)
}

func InitAPI(port string) {
	registerAPI(news.LoadNews)
	registerAPI(news2.LoadSpirit)
	engin := gin.Default()
	g := engin.Group("/api")
	for _, api := range apis {
		api(g)
	}
	engin.Run(":" + port)
}
