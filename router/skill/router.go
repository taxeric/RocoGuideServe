package skill

import "github.com/gin-gonic/gin"

func LoadSkill(r *gin.RouterGroup) {
	g := r.Group("/skill")
	loadSkillV1(g)
}

func loadSkillV1(r *gin.RouterGroup) {
	group := r.Group("/v1")
	group.GET("/skills", getSkillPages)
}
