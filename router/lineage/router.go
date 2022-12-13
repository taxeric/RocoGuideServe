package lineage

import "github.com/gin-gonic/gin"

func LoadLineage(g *gin.RouterGroup) {
	var group = g.Group("/lineage")
	loadLineageV1(group)
}

func loadLineageV1(g *gin.RouterGroup) {
	var group = g.Group("/v1")
	group.GET("/lineage", getAllLineage)
	group.PUT("/lineage", updateLineage)
	group.POST("/lineage", insertLineage)
}
