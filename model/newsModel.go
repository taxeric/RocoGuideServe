package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetNewsList(page int, amount int) ([]entity.News, int) {
	var sql = "select id,title,create_time,update_time,content,type,url from news order by id desc limit ? offset ?"
	var sqlAmount = "select COUNT(*) from news"
	row, _ := utils.Database.Query(sql, amount, (page-1)*amount)
	defer row.Close()
	var list = make([]entity.News, 0)
	for row.Next() {
		var news entity.News
		row.Scan(&news.Id, &news.Title, &news.Content, &news.Type, &news.CreateTime, &news.UpdateTime, news.Url)
		list = append(list, news)
	}
	row, _ = utils.Database.Query(sqlAmount)
	defer row.Close()
	var total int
	for row.Next() {
		row.Scan(&total)
	}
	return list, total
}
