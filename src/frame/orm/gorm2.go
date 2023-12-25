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
			TablePrefix:   "qm_", // 表名前缀
			SingularTable: true,  // 表名使用单数
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

func multiDemo(db *gorm.DB) {

}

func GormDemo2() {
	println("gorm()")
	db := connDb()
	fmt.Println("db=", db)
	multiDemo(db)
}
