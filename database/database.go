package database

import (
	"fmt"
	"jtn/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_NAME))

	if err != nil {
		panic("database error")
	}

	err = db.DB().Ping()

	if err != nil {
		panic("DSN Error")
	}
}

func DBManager() *gorm.DB {
	return db
}
