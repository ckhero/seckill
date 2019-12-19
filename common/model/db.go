package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"seckill/common/constant"
	"seckill/config"
)
var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
func InitDB()  {
	db = connDB()

	if constant.GlobalProd == config.RunEnv {
		db.SetLogger(logrus.StandardLogger())
	} else {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(10)
}

func connDB() (db *gorm.DB) {
	dbConfig := config.AppConfig.DB
	connUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := gorm.Open(dbConfig.Driver, connUrl)
	if err != nil {
		panic("sql connect fail")
	}
	return
}
func CloseDB() {
	if db != nil {
		_ = db.Close()
	}
}
