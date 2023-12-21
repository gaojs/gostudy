package gorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
}

func Demo() {
	println("gorm()")
	dbstr := "root:123@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	dl := mysql.Open(dbstr)
	cfg := gorm.Config{}
	db, err := gorm.Open(dl, &cfg)
	if err != nil {
		panic("连接数据库失败！")
	}
	fmt.Println("db=", db, "err=", err)
	db.AutoMigrate(&Product{})
	p1 := Product{Name: "苹果", Price: 1}
	fmt.Println("p1=", p1)
	db.Create(&p1) // 创建
	var p2 Product
	db.First(&p2) // 查询
	fmt.Println("p2=", p2)
	db.Model(&p2).Update("Price", "100")
	fmt.Println("p3=", p2)
	db.Delete(&p2) // 删除
	fmt.Println("p4=", p2)
}
