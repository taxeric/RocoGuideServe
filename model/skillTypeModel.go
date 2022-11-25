package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetAllSkillType() []entity.SkillType {
	var sql = "select id,name from skill_type"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list = make([]entity.SkillType, 0)
	for row.Next() {
		var g entity.SkillType
		row.Scan(&g.Id, &g.Name)
		list = append(list, g)
	}
	return list
}
