package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetGroupList() []entity.SpiritGroup {
	var sql = "select id,name from group_table"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list []entity.SpiritGroup = make([]entity.SpiritGroup, 0)
	for row.Next() {
		var group entity.SpiritGroup
		row.Scan(&group.Id, &group.Name)
		list = append(list, group)
	}
	return list
}
