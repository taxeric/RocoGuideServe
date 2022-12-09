package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func InsertSeries(series entity.SpiritSeries) int64 {
	var sql = "insert into series(name) values (?)"
	result, _ := utils.Database.Exec(sql, series.Name)
	id, _ := result.LastInsertId()
	return id
}

func GetAllSeries() []entity.SpiritSeries {
	var sql = "select * from series"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list = make([]entity.SpiritSeries, 0)
	for row.Next() {
		var singleSeries entity.SpiritSeries
		row.Scan(&singleSeries.Id, &singleSeries.Name)
		list = append(list, singleSeries)
	}
	return list
}

func UpdateSeries(series *entity.SpiritSeries) int64 {
	var sql = "update series set name=? where id=?"
	row, _ := utils.Database.Exec(sql, series.Name, series.Id)
	seriesId, _ := row.RowsAffected()
	return seriesId
}
