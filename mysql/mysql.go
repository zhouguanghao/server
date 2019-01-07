package mysql

import (
	"database/sql"
	"github.com/davyxu/golog"
	_ "github.com/go-sql-driver/mysql"
)

var (
	log = golog.New("mysql")
	DB *sql.DB
)

func init() {
	log.Debugln("Mysql init.")
	var err error
	DB,err = sql.Open("mysql", "user:zgh1625347@tcp(127.0.0.1:3306)/user?charset=utf8")
	if err != nil {
		log.Errorf("Open mysql database error: %s\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Errorln(err)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(50)

	GetAllPlayerInfo()
}
