package model

type Product struct {
	Base
	Id int32 `gorm:"id;primary_key"`
	ProductName string `gorm:"product_name"`
	ProductQty int32 `gorm:"product_qty"`
}

func (p *Product)Create()  {
	p.GetDB().Save(p)
}