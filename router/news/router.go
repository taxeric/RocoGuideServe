package news

import (
	"github.com/gin-gonic/gin"
)

func LoadNews(g *gin.RouterGroup) {
	group := g.Group("/news")
	loadNewsV1(group)
}

func loadNewsV1(g *gin.RouterGroup) {
	group := g.Group("/v1")
	group.GET("/page", getNews)
}
