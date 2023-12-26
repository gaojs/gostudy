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
		// DisableForeignKeyConstraintWhenMigrating: false, // 外键关系
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

func one2one(db *gorm.DB) {
	// err := db.AutoMigrate(&Dog0{}, &Girl0{})
	// fmt.Println("AutoMigrate, err=", err)
	if false { // BelongsTo属于1对1关系的演示
		db.Migrator().DropTable(&Dog0{}, &Girl0{})
		// 创建舔狗表的时候，会自动创建女神表
		err := db.AutoMigrate(&Dog0{})
		fmt.Println("AutoMigrate, err=", err)

		// 舔狗属于女神，它有女神ID和女神结构体
		d := Dog0{Name: "舔狗", Girl: Girl0{Name: "女神"}}
		db.Debug().Create(&d)
		// INSERT INTO `girl0` (`name`) VALUES ('女神') ON DUPLICATE KEY UPDATE `id`=`id` RETURNING `id`
		// INSERT INTO `dog0` (`name`,`girl0_id`) VALUES ('舔狗',2) RETURNING `id`
		fmt.Println("BelongsTo, dog=", d)
		d0 := Dog0{}
		db.Debug().Preload("Girl").First(&d0)
		// SELECT * FROM `girl0` WHERE `girl0`.`id` = 1
		// SELECT * FROM `dog0` ORDER BY `dog0`.`id` LIMIT 1
		fmt.Println("Preload, dog=", d0)
		// err = db.Debug().Model(&d0).Association("Girl").Delete(&d0.Girl)
		// // UPDATE `dog0` SET `girl`=NULL WHERE `dog0`.`id` = 1 AND `dog0`.`girl` = 1
		// fmt.Println("Association, err=", err, d0)
		g := Girl0{ID: d.Girl.ID}
		db.Debug().Delete(&g)
		fmt.Println("Delete, g=", g)
	}
	if true { // HasOne拥有1对1关系的演示
		// 各自创建舔狗表和女神表
		db.Migrator().DropTable(&Girl1{}, &Dog1{})
		err := db.AutoMigrate(&Girl1{}, &Dog1{})
		fmt.Println("AutoMigrate, err=", err)

		// 女神拥有舔狗，舔狗也有女神ID
		// g := Girl1{Name: "女神"} // 狗的字段全空
		g := Girl1{Name: "女神", Dog: Dog1{Name: "舔狗"}}
		db.Debug().Create(&g)
		// INSERT INTO `dog1` (`name`,`girl1_id`) VALUES ('舔狗',1) ON DUPLICATE KEY UPDATE `girl1_id`=VALUES(`girl1_id`) RETURNING `id`
		// INSERT INTO `girl1` (`name`) VALUES ('女神') RETURNING `id`
		fmt.Println("HasOne, girl=", g)
		g1 := Girl1{}
		db.Debug().Preload("Dog").First(&g1)
		// SELECT * FROM `dog1` WHERE `dog1`.`girl1_id` = 1
		// SELECT * FROM `girl1` ORDER BY `girl1`.`id` LIMIT 1
		fmt.Println("Preload, girl=", g1)
		// // err = db.Debug().Model(&g1).Association("Dog1").Delete(&g1.Dog1)
		// // UPDATE `dog0` SET `girl`=NULL WHERE `dog0`.`id` = 1 AND `dog0`.`girl` = 1
		// err = db.Debug().Model(&g1).Association("Dog").Clear()
		// // UPDATE `dog1` SET `girl`=NULL WHERE `dog1`.`girl` = 1
		// fmt.Println("Association, err=", err, g1)
		db.Debug().Delete(&g1)
		fmt.Println("Delete, g=", g1)
	}
}

func one2many(db *gorm.DB) {
	// 各自创建舔狗表和女神表
	db.Migrator().DropTable(&Girl2{}, &Dog2{}, &DogInfo{})
	err := db.AutoMigrate(&Girl2{}, &Dog2{}, &DogInfo{})
	fmt.Println("AutoMigrate, err=", err)

	g := Girl2{Name: "女神", Dogs: []Dog2{
		{Name: "舔狗1", Info: DogInfo{Money: 10000}},
		{Name: "舔狗2", Info: DogInfo{Money: 100}}}}
	db.Debug().Create(&g)
	fmt.Println("HasMany, girl=", g)
	g2 := make([]Girl2, 0)
	db.Debug().Preload("Dogs.Info", "money < 1000").Preload("Dogs", "name=?", "舔狗1").Find(&g2)
	// SELECT * FROM `dog_info` WHERE `dog_info`.`dog_id` = 1 AND money < 1000
	// SELECT * FROM `dog2` WHERE `dog2`.`girl_id` = 1 AND name='舔狗1'
	// SELECT * FROM `girl2`
	fmt.Println("Preload, girl=", g2)
	db.Debug().Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money < 1000")
	}).Find(&g2)
	fmt.Println("Joins, girl=", g2)
	// err = db.Debug().Model(&g2).Association("Dogs").Clear()
	// fmt.Println("Association, err=", err, g2)
}

func many2many(db *gorm.DB) {
	// 删除舔狗表和女神表时，不会删除关联关系表
	db.Migrator().DropTable(&Girl3{}, &Dog3{}, "dog_girl")
	// 自动创建Dog和关联关系表
	err := db.AutoMigrate(&Girl3{})
	fmt.Println("AutoMigrate, err=", err)

	d1 := Dog3{Name: "舔狗1", ID: 1}
	d2 := Dog3{Name: "舔狗2", ID: 2}
	fmt.Println("dogs=", d1, d2)
	g1 := Girl3{Name: "女神1", Dogs: []Dog3{d1, d2}}
	g2 := Girl3{Name: "女神2", Dogs: []Dog3{d1, d2}}
	fmt.Println("girls=", g1, g2)
	db.Debug().Create(&g1)
	db.Debug().Create(&g2)
	// INSERT INTO `girl3` (`name`) VALUES ('女神1') RETURNING `id`
	// INSERT INTO `dog3` (`name`,`id`) VALUES ('舔狗1',1),('舔狗2',2) ON DUPLICATE KEY UPDATE `id`=`id` RETURNING `id`
	// INSERT INTO `dog_girl` (`girl3_id`,`dog3_id`) VALUES (2,1),(2,2) ON DUPLICATE KEY UPDATE `girl3_id`=`girl3_id`
	db.Debug().Preload("Girls").Find(&d1)
	db.Debug().Preload("Girls").Find(&d2)
	fmt.Println("dogs=", d1, d2)
	girls := []Girl3{}
	db.Debug().Model(&d1).Preload("Dogs").Association("Girls").Find(&girls)
	fmt.Println("girls=", girls)
	db.Debug().Model(&d2).Association("Girls").Clear()
	// DELETE FROM `dog_girl` WHERE `dog_girl`.`dog3_id` = 2
	db.Debug().Model(&d2).Association("Girls").Replace(&g1, &g2)
	// DELETE FROM `dog_girl` WHERE `dog_girl`.`dog3_id` = 2 AND `dog_girl`.`girl3_id` NOT IN (1,2)
}

func polyDemo(db *gorm.DB) {
	db.Migrator().DropTable(&Tire{}, &Bike{}, &Car{})
	// 自动创建Dog和关联关系表
	err := db.AutoMigrate(&Tire{}, &Bike{}, &Car{})
	fmt.Println("AutoMigrate, err=", err)

	b := Bike{Name: "单车1", Tires: []Tire{{Name: "细轮胎"}}}
	c := Car{Name: "汽车1", Tires: []Tire{{Name: "粗轮胎"}, {Name: "越野胎"}}}
	db.Create(&b)
	db.Create(&c)
}

func tagDemo(db *gorm.DB) {
	// 删除舔狗表和女神表时，不会删除关联关系表
	db.Migrator().DropTable(&Poet{}, &Peom{})
	// 自动创建Dog和关联关系表
	err := db.AutoMigrate(&Poet{}, &Peom{})
	fmt.Println("AutoMigrate, err=", err)
	p := Poet{Name: "李白", Peoms: []Peom{{Name: "静夜思"}, {Name: "蜀道难"}}}
	db.Create(&p)
	fmt.Println("Create, p=", p)
}

func GormDemo2() {
	println("gorm()")
	db := connDb()
	fmt.Printf("db=%T, %v\n", db, db)
	sqlDb, _ := db.DB()
	fmt.Printf("sqlDb=%T, %v\n", sqlDb, sqlDb)
	// tableDemo(db)
	// recordDemo(db)
	// one2one(db)
	// one2many(db)
	// many2many(db)
	polyDemo(db)
	// tagDemo(db)
}
