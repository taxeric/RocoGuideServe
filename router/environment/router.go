package environment

import "github.com/gin-gonic/gin"

func LoadSkillEnvironment(g *gin.RouterGroup) {
	var group = g.Group("/environment")
	loadSkillEnvironment(group)
}

func loadSkillEnvironment(g *gin.RouterGroup) {
	var group = g.Group("/v1")
	group.POST("/environment", insertEnvironment)
	group.GET("/environment", getAllEnvironment)
	group.PUT("/environment", updateEnvironment)
}
