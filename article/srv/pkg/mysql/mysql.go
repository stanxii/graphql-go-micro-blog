package db

import (
	"srv/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
)

var DB *gorm.DB

func init() {
	mysql := config.Config.Mysql
	db, err := gorm.Open(
		mysql.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			mysql.User,
			mysql.Password,
			mysql.Host,
			mysql.Port,
			mysql.Table,
		),
	)

	if err != nil {
		log.Fatal("mysql error:", err)
		panic(err)
	}

	db.Debug()

	DB = db
}