package database

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

type DBConfig struct {
	DbName    string
	User      string
	Password  string
	Host      string
	Charset   string
	ParseTime bool
	Local     string
	Port      int
}

func init() {
	//数据库连接池
	var err error

	//本地
	var dbConfig = DBConfig{DbName: "stest", User: "root", Password: "123456", Host: "127.0.0.1", Charset: "utf8", ParseTime: true, Local: "Local"}

	SqlDB, err = manager.New(dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.Host).Set(
		manager.SetCharset(dbConfig.Charset),
		manager.SetParseTime(dbConfig.ParseTime),
		manager.SetLoc(dbConfig.Local)).Open(true)

	if err != nil {
		log.Fatalln(err.Error())
	}

	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	if err := SqlDB.Ping(); err != nil {
		log.Fatalln(err.Error())
	}
}
