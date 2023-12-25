package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func connDb() (db *gorm.DB) {
	// 建立数据库链接
	mysqlConfig := mysql.Config{ // 可以配置更多项，不只是DSN(DataSourceName)
		DSN:                       "root:123@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string字符串类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用时间精度，MYSQL5.6之前不支持
		DontSupportRenameIndex:    true,  // 删除再重建索引，MariaDB不支持直接重命名
		DontSupportRenameColumn:   true,  // 用Change重命名，MariaDB不支持直接重命名
		SkipInitializeWithVersion: false, //根据MYSQL版本自动配置
	}
	gormConfig := gorm.Config{ // 可以配置更多项
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "qm_", // 表名前缀
			SingularTable: true, // 表名使用单数
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键（代码中指定外键关系）
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		fmt.Println("err=", err)
		panic("连接数据库失败！")
	}
	return db
}

func tableDemo(db *gorm.DB) {
	// 自动迁移创建表
	// err := db.AutoMigrate(&User{})
	// fmt.Println("err=", err)
	// 使用Migrator建表
	dbm := db.Migrator()
	if !dbm.HasTable(&User{}) {
		err := dbm.CreateTable(&User{})
		fmt.Println("err=", err)
	} else {
		err := dbm.DropTable(&User{})
		fmt.Println("err=", err)
	}
}

func recordDemo(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &UserDetail{})
	fmt.Println("AutoMigrate, err=", err)
	res := db.Create(&User{Name: "gao", Age: 12})
	fmt.Println("Create, res.RowsAffected=", res.RowsAffected)
	u := make([]User, 0)
	db.Where("name like ?", "%a%").Find(&u)
	fmt.Println("Find", u)
	ret := db.Unscoped().Delete(u)
	fmt.Println("Delete, ret.RowsAffected=", ret.RowsAffected)
}

func belongDemo(db *gorm.DB) {
	err := db.AutoMigrate(&Dog0{}, &Girl0{})
	fmt.Println("AutoMigrate, err=", err)
}

func GormDemo2() {
	println("gorm()")
	db := connDb()
	fmt.Printf("db=%T, %v\n", db, db)
	sqlDb, _ := db.DB()
	fmt.Printf("sqlDb=%T, %v\n", sqlDb, sqlDb)
	// tableDemo(db)
	// recordDemo(db)
	belongDemo(db)
}
