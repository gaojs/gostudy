package orm

import (
	"time"

	"gorm.io/gorm"
)

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
	gorm.Model        // ID int
	Name       string `gorm:"default:'def'"`
	Age        int    `gorm:"default:20"`
}

type Address struct {
	Country  string
	Province string
}

type UserInfo struct {
	ID         int64     `xorm:"pk autoincr"`    // 主键，自增(int)
	Name       string    `xorm:"varchar(24)"`    // 字符串(varchar)
	Age        int       `xorm:"int default 18"` // 默认值(int)
	Male       bool      `xorm:"index notnull"`  // 建立索引(tinyint)
	Email      string    `xorm:"-"`              // 不映射(varchar)
	Income     float64   `xorm:"->"`             // 只写入不读取(double)
	CreateTime time.Time `xorm:"created"`        // 自动设置时间(datetime)
	UpdateTime time.Time `xorm:"updated"`        // 自动设置时间(datetime)
	DeleteTime time.Time `xorm:"deleted"`        // 自动设置时间(datetime)
	Version    int       `xorm:"version"`        // 插入时为1，更新时自增
	Addr       Address   `xorm:"json"`           // 结构体存为JSON(Text)
	// Slice      []byte    // 切片(Blob)
	// Array      [3]byte   // 数组(Blob)
	// Score   map[string]any // 映射(Text)
	// Complex complex128     // varchar(64)
	// Float   float32        // float
}
