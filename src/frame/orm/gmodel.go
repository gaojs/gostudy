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
type Dog struct { // 舔狗
	ID   int `gorm:"primaryKey"`
	Name string
}

type Girl struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
}

// 一对一的关系
// BelongsTo属于
type Dog0 struct { // 舔狗
	ID   int `gorm:"primaryKey"`
	Name string
	// 靠该字段，关联女神
	GirlID int `gorm:"column:girl_id"` // 所属女神的ID
	// [error] define a valid foreign key for relations
	// 有GirlID后，Girl才不会报上述的错误
	Girl Girl0 `gorm:"ForeignKey:GirlID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 所属女神
}

type Girl0 struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
}

// 一对一的关系
type Dog1 struct { // 舔狗
	ID   int `gorm:"primaryKey"`
	Name string
	// 靠该字段，关联女神
	GirlID int `gorm:"column:girl_id"` // 所属女神的ID
}

// HasOne拥有
type Girl1 struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
	// invalid recursive type Dog1
	Dog Dog1 `gorm:"ForeignKey:GirlID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 拥有的舔狗
}

// 一对多的关系
type DogInfo struct { // 舔狗的信息
	ID     int `gorm:"primaryKey"`
	Money  int
	Dog2ID int `gorm:"column:dog_id"` // 所属舔狗的ID
}

type Dog2 struct { // 舔狗
	ID   int `gorm:"primaryKey"`
	Name string
	Info DogInfo
	// 靠该字段，关联女神
	Girl2ID int `gorm:"column:girl_id"` // 所属女神的ID
}

// HasMany拥有
type Girl2 struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
	Dogs []Dog2 // 拥有的舔狗
}

// 多对多的关系
type Dog3 struct { // 舔狗
	ID    int `gorm:"primaryKey"`
	Name  string
	Girls []Girl3 `gorm:"many2many:dog_girl"` // 拥有的女神
}

// HasMany拥有
type Girl3 struct { // 女神
	ID   int `gorm:"primaryKey"`
	Name string
	Dogs []Dog3 `gorm:"many2many:dog_girl"` // 拥有的舔狗
}

// 多态
type Tire struct { // 轮胎
	ID   int `gorm:"primaryKey"`
	Name string
	// 下述字段，进行关联多态
	OwnerID   int
	OwnerType string
}

type Bike struct { // 单车
	ID    int `gorm:"primaryKey"`
	Name  string
	Tires []Tire `gorm:"polymorphic:Owner;polymorphicValue:bike"` // 多态
}

type Car struct { // 汽车
	ID    int `gorm:"primaryKey"`
	Name  string
	Tires []Tire `gorm:"polymorphic:Owner;polymorphicValue:car"` // 多态
}

// 外键标签
type Poet struct { // 诗人
	ID   int `gorm:"primaryKey"`
	Name string
	// Peoms []Peom // 靠该字段，关联
	// 诗歌的外键，是OwnerName，关联诗人的Name
	Peoms []Peom `gorm:"foreignKey:Owner;references:Name"`
}

type Peom struct { // 诗歌
	ID   int `gorm:"primaryKey"`
	Name string
	// PoetID int // 靠该字段，关联
	Owner string
}
