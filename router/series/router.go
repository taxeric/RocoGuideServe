package series

import "github.com/gin-gonic/gin"

func LoadSeries(g *gin.RouterGroup) {
	var group = g.Group("/series")
	loadSeriesV1(group)
}

func loadSeriesV1(g *gin.RouterGroup) {
	var group = g.Group("/v1")
	group.POST("/series", insertSeries)
	group.GET("/series", getAllSeries)
	group.PUT("/series", updateSeries)
}
