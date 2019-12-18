package model

import "github.com/jinzhu/gorm"

type Base struct {
	db *gorm.DB
}

func (base *Base) GetDB() *gorm.DB {
	if base.db == nil {
		base.SetDB(GetDB())
	}
	return base.db
}

func (base *Base) SetDB(db *gorm.DB) {
	base.db = db
}
