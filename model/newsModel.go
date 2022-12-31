package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
	"fmt"
)

func GetNewsList(page int, amount int) ([]entity.News, int) {
	var sql = "select id,title,create_time,update_time,type,url from news order by id desc limit ? offset ?"
	var sqlAmount = "select COUNT(*) from news"
	row, _ := utils.Database.Query(sql, amount, (page-1)*amount)
	defer row.Close()
	var list = make([]entity.News, 0)
	for row.Next() {
		var news entity.News
		row.Scan(&news.Id, &news.Title, &news.CreateTime, &news.UpdateTime, &news.Type, &news.Url)
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

func InsertNews(url string, title string, updateTim string) int64 {
	return insertData(1, url, title, updateTim)
}

func insertData(newsType int, contentOrUrl string, title string, updateTime string) int64 {
	var sql string
	if newsType == 1 {
		sql = "insert into news(type,url,title, update_time) values (?,?,?,?)"
	} else {
		sql = "insert into news(type,content,title, update_time) values (?,?,?,?)"
	}
	result, _ := utils.Database.Exec(sql, newsType, contentOrUrl, title, updateTime)
	id, _ := result.LastInsertId()
	return id
}

func UpdateNews(news entity.News) int64 {
	var sql = "update news set url=?,title=? where id=?"
	row, err := utils.Database.Exec(sql, news.Url, news.Title, news.Id)
	if err != nil {
		fmt.Println(err)
	}
	newsId, _ := row.RowsAffected()
	return newsId
}
