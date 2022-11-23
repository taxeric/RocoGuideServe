package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetAllAttrs() []entity.SpiritAttributes {
	var sql = "select id,name from genius_attributes"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list = make([]entity.SpiritAttributes, 0)
	for row.Next() {
		var group entity.SpiritAttributes
		row.Scan(&group.Id, &group.Name)
		list = append(list, group)
	}
	return list
}
