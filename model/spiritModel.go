package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetSpiritList(page int, amount int, keywords string) ([]entity.SpiritListItem, int) {
	var sql = "select g.id,g.number,g.avatar,g.name,att.id,att.name ,att2.id,att2.name from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ?  order by g.id desc limit ? offset ?"
	var sqlAmount = "select COUNT(*) from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ?"
	row, err := utils.Database.Query(sql, "%"+keywords+"%", amount, (page-1)*amount)
	if err != nil {
		return *new([]entity.SpiritListItem), 0
	}
	defer row.Close()
	var list = make([]entity.SpiritListItem, 0)
	for row.Next() {
		var spirit entity.SpiritListItem
		row.Scan(&spirit.ID,
			&spirit.Number,
			&spirit.Avatar,
			&spirit.Name,
			&spirit.PrimaryAttributes.Id,
			&spirit.PrimaryAttributes.Name,
			&spirit.SecondaryAttributes.Id,
			&spirit.SecondaryAttributes.Name)
		list = append(list, spirit)
	}
	row, err = utils.Database.Query(sqlAmount, "%"+keywords+"%")
	if err != nil {
		return *new([]entity.SpiritListItem), 0
	}
	defer row.Close()
	var total int
	for row.Next() {
		row.Scan(&total)
	}
	return list, total
}
