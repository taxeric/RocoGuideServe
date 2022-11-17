package main

import "RocoGuide/utils"

func main() {
	var config = utils.MysqlConfig{
		Ip:     "127.0.0.1",
		Port:   3306,
		Unm:    "unm",
		Pwd:    "pwd",
		DbName: "dbname",
	}
	err := utils.InitDb(config)
	if err != nil {
		return
	}
	println("初始化成功")
}
