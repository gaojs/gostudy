package orm

import (
	"database/sql"
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
	// ID   int    //gorm.Model
	UUID uint   `gorm:"primaryKey;comment:唯一标识"` // 设置为主键
	Name string `gorm:"default:'hi';comment:姓名"` // 设置默认值
	Age  int    `gorm:"default:20;comment:年龄"`   // 设置默认值
}

type UserDetail struct {
	// [error] invalid field found for struct main/frame/orm.TestUser's field User:
	// define a valid foreign key for relations or implement the Valuer/Scanner interface
	User   User         `gorm:"embedded"`          // 指明内嵌，以免报错；User的所有字段，都会内嵌进来
	Email  *string      `gorm:"not null"`          // 设置为非空
	Birth  *time.Time   `gorm:"column:birthday"`   // 指明列名
	Active sql.NullTime `gorm:"column:actived_at"` // 设置注释
}

// 完全无关的俩结构
type Dog0 struct { // 舔狗
	ID   int `gorm:"primaryKey"`
	Name string
}

type Girl0 struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
}
