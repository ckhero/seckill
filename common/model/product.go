package model

import "github.com/jinzhu/gorm"

type Product struct {
	Base
	ProductId int32 `gorm:"product_id;primary_key"`
	ProductName string `gorm:"product_name"`
	ProductQty int32 `gorm:"product_qty"`
	ProductExtends []ProductExtend `gorm:"ForeignKey:product_extend_product_id;AssociationForeignKey:product_id"`
}

/**
表名
 */
func (p *Product) TableName() string {
	return "product"
}

func (p *Product)Create() (err error) {
	err = GetDB().Save(p).Error
	return
}

func (p *Product) FindOne(productId int) (*Product, error) {
	if err := GetDB().First(p, productId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	//GetDB().Model(p).Related(new(ProductExtend), "ProductExtends").Find(p.ProductExtends)
	return p, nil
}

func (p *Product) GetProductExtends() ([]ProductExtend, error) {
	if err := GetDB().Model(p).Related(new(ProductExtend), "ProductExtends").Find(&p.ProductExtends).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, nil
		}
		return nil, err
	}
	return p.ProductExtends, nil
}