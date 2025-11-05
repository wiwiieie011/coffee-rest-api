package models

type Drink struct{
	ID uint  `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Price int `json:"price"`
	Instock bool `json:"is_stock" gorm:"column:is_stock"`
	ContainsCoffeine bool `json:"containsCaffeine"`
	Volume int `json:"volume"`
	Description string `json:"description"`
}