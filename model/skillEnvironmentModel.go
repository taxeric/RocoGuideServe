package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func InsertEnvironment(environment entity.SkillEnvironment) int64 {
	var sql = "insert into environment(name,introduce,effects,type,icon) values (?,?,?,?,?)"
	result, _ := utils.Database.Exec(sql, environment.Name, environment.Introduce, environment.Effects, environment.Type, environment.Icon)
	id, _ := result.LastInsertId()
	return id
}

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
			&environment.Introduce,
			&environment.Effects,
			&environment.Type,
			&environment.Icon,
		)
		list = append(list, environment)
	}
	return list
}

func UpdateEnvironment(environment entity.SkillEnvironment) int64 {
	var sql = "update environment set name=? introduce=? effects=? type=? icon=? where id=?"
	row, _ := utils.Database.Exec(
		sql,
		environment.Name,
		environment.Introduce,
		environment.Effects,
		environment.Type,
		environment.Icon,
		environment.Id,
	)
	environmentId, _ := row.RowsAffected()
	return environmentId
}
