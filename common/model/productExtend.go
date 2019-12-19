package model

import (
	"github.com/jinzhu/gorm"
)

type ProductExtend struct {
	Base
	ProductExtendId int32 `gorm:"product_extend_id;primary_key"`
	ProductExtendProductId int32 `gorm:"product_extend_product_id"`
	ProductExtendDesc string `gorm:"product_extend_desc"`
	Product Product `gorm:"ForeignKey:product_id;AssociationForeignKey:product_extend_product_id"`
}

/**
表名
*/
func (pe *ProductExtend) TableName() string {
	return "product_extend"
}

func (pe *ProductExtend)Create()  {
	GetDB().Save(pe)
}

func (pe *ProductExtend) FindOne(productExtendId int) (*ProductExtend, error) {
	if err := GetDB().First(pe, productExtendId).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, nil
		}
		return nil, err
	}
	GetDB().Model(pe).Related(new(Product), "Product").Find(pe.Product)
	return pe, nil
}