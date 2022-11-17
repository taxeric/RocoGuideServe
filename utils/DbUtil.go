package utils

import (
	"RocoGuide/constants"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var Database *sql.DB

type MysqlConfig struct {
	Ip     string
	Port   int
	Unm    string
	Pwd    string
	DbName string
}

func InitDb(config MysqlConfig) error {
	var dataSourceName = config.Unm + ":" + config.Pwd + "@tcp(" + config.Ip + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName
	fmt.Println(dataSourceName)
	Database, err := sql.Open(constants.DriverName, dataSourceName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = Database.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	Database.SetConnMaxLifetime(time.Minute * 3)
	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(10)
	return nil
}
