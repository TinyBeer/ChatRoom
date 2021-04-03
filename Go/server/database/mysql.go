package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	sqldb *sql.DB
)

func Query(queryStr string, arg ...interface{}) (*sql.Rows, error) {
	return sqldb.Query(queryStr, arg...)
}

func Exec(queryStr string, arg ...interface{}) error {
	_, err := sqldb.Exec(queryStr, arg...)
	return err
}

func Close() {
	sqldb.Close()
}

func Init() (err error) {
	// connect to mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.database"))
	// dataSourceName
	sqldb, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to open database, err:" + err.Error())
	}

	err = sqldb.Ping()
	if err != nil {
		panic("failed to connect to mysql database, err:" + err.Error())
	}
	return nil
}
