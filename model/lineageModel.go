package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func InsertLineage(lineage entity.Lineage) int64 {
	var sql = "insert into lineage(name,introduce,icon) values (?,?,?)"
	row, _ := utils.Database.Exec(sql, lineage.Name, lineage.Introduce, lineage.Icon)
	id, _ := row.LastInsertId()
	return id
}

func UpdateLineage(lineage entity.Lineage) int64 {
	var sql = "update lineage set name=?,introduce=?,icon=? where id=?"
	row, _ := utils.Database.Exec(sql, lineage.Name, lineage.Introduce, lineage.Icon, lineage.Id)
	id, _ := row.RowsAffected()
	return id
}

func GetAllLineage() []entity.Lineage {
	var sql = "select * from lineage"
	row, _ := utils.Database.Query(sql)
	var list = make([]entity.Lineage, 0)
	for row.Next() {
		var lineage entity.Lineage
		row.Scan(
			&lineage.Id,
			&lineage.Name,
			&lineage.Introduce,
			&lineage.Icon,
		)
		list = append(list, lineage)
	}
	return list
}
