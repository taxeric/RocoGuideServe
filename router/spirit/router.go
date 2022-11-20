package news

import (
	"github.com/gin-gonic/gin"
)

func LoadSpirit(g *gin.RouterGroup) {
	group := g.Group("/spirit")
	loadSpiritV1(group)
}

func loadSpiritV1(g *gin.RouterGroup) {
	group := g.Group("/v1")
	group.GET("/page", getSpiritList)
}
