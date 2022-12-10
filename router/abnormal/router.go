package abnormal

import "github.com/gin-gonic/gin"

func LoadAbnormalState(g *gin.RouterGroup) {
	var group = g.Group("/abnormal")
	loadAbnormalStateV1(group)
}

func loadAbnormalStateV1(g *gin.RouterGroup) {
	var group = g.Group("/v1")
	group.POST("/abnormal", insertAbnormal)
	group.GET("/abnormal", getAllAbnormal)
	group.PUT("/abnormal", updateAbnormal)
}
