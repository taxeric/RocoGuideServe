package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetAllEnvironment() []entity.SkillEnvironment {
	var sql = "select * from environment"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list = make([]entity.SkillEnvironment, 0)
	for row.Next() {
		var environment entity.SkillEnvironment
		row.Scan(
			&environment.Id,
			&environment.Name,
			&environment.Effects,
			&environment.Icon,
			&environment.Type,
		)
		list = append(list, environment)
	}
	return list
}

func UpdateEnvironment(environment entity.SkillEnvironment) int64 {
	var sql = "update environment set name=? effects=? type=? icon=? where id=?"
	row, _ := utils.Database.Exec(
		sql,
		environment.Name,
		environment.Effects,
		environment.Type,
		environment.Icon,
		environment.Id,
	)
	environmentId, _ := row.RowsAffected()
	return environmentId
}
