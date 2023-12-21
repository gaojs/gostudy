package orm

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string
	Price int
}

// 通过函数，指定表名
func (Product) TableName() string {
	return "product"
}

type User struct {
	Id   int    // gorm.Model
	Name string `gorm:"default:'def'"`
	Age  int    // `gorm:"default:20"`
}
