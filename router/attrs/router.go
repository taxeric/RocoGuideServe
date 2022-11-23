package attrs

import "github.com/gin-gonic/gin"

func LoadAttrs(g *gin.RouterGroup) {
	group := g.Group("/attrs")
	loadAttrsV1(group)
}

func loadAttrsV1(g *gin.RouterGroup) {
	group := g.Group("/v1")
	group.GET("/page", getAllAttrs)
}
