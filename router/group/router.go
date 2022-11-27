package group

import "github.com/gin-gonic/gin"

func LoadSpiritGroup(g *gin.RouterGroup) {
	group := g.Group("/group")
	loadSpiritGroupV1(group)
}

func loadSpiritGroupV1(g *gin.RouterGroup) {
	group := g.Group("/v1")
	group.GET("/groups", getAllGroup)
}
