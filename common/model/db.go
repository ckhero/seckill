package model

import (
	"github.com/jinzhu/gorm"
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
	dbConfig := config.AppConfig.DB
	connUrl := dbConfig.Username + ":" + dbConfig.Password + "@" + dbConfig.Host + "/" +dbConfig.Database + "charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dbConfig.Drive, connUrl)
	if err != nil {
		panic("sql connect fail")
	}
	db.DB().SetMaxIdleConns(10)
}