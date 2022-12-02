package main

import (
	"RocoGuide/router"
	"RocoGuide/utils"
)

func main() {
	var config = utils.MysqlConfig{
		Ip:     "127.0.0.1",
		Port:   3306,
		Unm:    "root",
		Pwd:    "1248",
		DbName: "raiders",
	}
	err := utils.InitDb(config)
	if err != nil {
		return
	}
	router.InitAPI("8888")
}
