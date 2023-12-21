package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func productDemo(db *gorm.DB) {
	db.AutoMigrate(&Product{}) // 建立OR关系
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

func createDemo(db *gorm.DB) {
	u1 := User{Name: "hill", Age: 18}
	// fmt.Println(db.NewRecord(&u1))
	db.Create(&u1) // 创建记录
	u2 := User{Name: "jimmy"}
	db.Debug().Create(&u2) // 创建记录
	u3 := User{}
	db.Debug().Create(&u3) // 创建记录
}

func readDemo(db *gorm.DB) {
	// 等于
	var u User
	db.Where("name = ?", "def").Debug().First(&u)
	// SELECT * FROM `users` WHERE name = 'def' ORDER BY `users`.`id` LIMIT 1
	fmt.Println("user=", u)
	// 不等于
	var us []User = make([]User, 1)
	db.Where("name != ?", "def").Debug().Find(&us)
	// SELECT * FROM `users` WHERE name != 'def'
	fmt.Println("users=", us)
	// 集合查询，等于
	db.Where("name IN ?", []string{"hill", "def"}).Debug().Find(&us)
	// SELECT * FROM `users` WHERE name IN ('hill','def')
	fmt.Println("users=", us)
	// 模糊查询，包含i
	db.Where("name LIKE ?", "%i%").Debug().Find(&us)
	// SELECT * FROM `users` WHERE name LIKE '%i%
	fmt.Println("users=", us)
	// 多条件组合(逻辑与)
	db.Where("name = ? AND age >= ?", "hill", 10).Debug().Find(&us)
	// SELECT * FROM `users` WHERE name = 'hill' AND age >= 10
	fmt.Println("users=", us)
	// 结构体查询
	db.Where(&User{Name: "hill", Age: 18}).Debug().Find(&us)
	// SELECT * FROM `users` WHERE `users`.`name` = 'hill' AND `users`.`age` = 18
	fmt.Println("users=", us)
	// MAP查询
	db.Where(map[string]any{"name": "hill", "age": 18}).Debug().Find(&us)
	// SELECT * FROM `users` WHERE `age` = 18 AND `name` = 'hill'
	fmt.Println("users=", us)
	db.Where([]int{5, 8, 10}).Debug().Find(&us)
	// SELECT * FROM `users` WHERE `users`.`id` IN (5,8,10)
	fmt.Println("users=", us)
	// NOT查询，逻辑非
	db.Not("name IN ?", []string{"hill", "def"}).Debug().Find(&us)
	// SELECT * FROM `users` WHERE NOT name IN ('hill','def')
	fmt.Println("users=", us)
	// OR查询，逻辑或
	db.Where("name = ?", "hill").Or("name = ?", "def").Debug().Find(&us)
	// SELECT * FROM `users` WHERE name = 'hill' OR name = 'def'
	fmt.Println("users=", us)
	// 多条件组合(逻辑或)
	db.Where("name = ? OR age >= ?", "hill", 10).Debug().Find(&us)
	// SELECT * FROM `users` WHERE name = 'hill' OR age >= 10
	fmt.Println("users=", us)
	// 多条件组合(逻辑或)，内联写法(省去Where)
	db.Debug().Find(&us, "name = ? OR age >= ?", "hill", 10)
	// SELECT * FROM `users` WHERE name = 'hill' OR age >= 10
	fmt.Println("users=", us)
	// 查找不到则初始化一个
	var u2 User
	// db.Debug().FirstOrInit(&u2, User{Name: "non"})
	// SELECT * FROM `users` WHERE `users`.`name` = 'non' ORDER BY `users`.`id` LIMIT 1
	db.Debug().FirstOrCreate(&u2, User{Name: "non"})
	// INSERT INTO `users` (`name`,`age`) VALUES ('non',0) RETURNING `id`
	fmt.Println("user=", u2)
	db.Debug().Table("users").Select("name, age").Where("name=?", "hill").Scan(&u2)
	// SELECT name, age FROM `users` WHERE name='hill'
	fmt.Println("user=", u2)
	db.Debug().Raw("SELECT name, age FROM users WHERE name=?", "hill").Scan(&u2)
	// SELECT name, age FROM users WHERE name='hill'
	fmt.Println("user=", u2)
}

func updateDemo(db *gorm.DB) {
	var u User
	db.Debug().First(&u)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	u.Name = "new3"
	db.Debug().Save(&u) // 修改所有字段
	// UPDATE `users` SET `name`='new',`age`=18 WHERE `id` = 1
	db.Debug().Model(&u).Update("name", "new2")
	// UPDATE `users` SET `name`='new2' WHERE `id` = 1
	m1 := map[string]any{"name": "hi", "age": 38}
	db.Debug().Model(&u).Updates(m1)
	// UPDATE `users` SET `age`=38,`name`='hi' WHERE `id` = 1
	// 只更新age字段
	db.Debug().Model(&u).Select("age").Updates(m1)
	// UPDATE `users` SET `age`=38 WHERE `id` = 1
	// 不更新age字段
	db.Debug().Model(&u).Omit("age").Updates(m1)
	// UPDATE `users` SET `name`='hi' WHERE `id` = 1
	// 在原有字段上更新
	db.Debug().Model(&u).Update("age", gorm.Expr("age+?", 2))
	//  UPDATE `users` SET `age`=age+2 WHERE `id` = 1
}

func deleteDemo(db *gorm.DB) {
	var u = User{Id: 1, Name: "hill", Age: 80}
	db.Debug().Delete(&u) // 删除时，主键不能空
	// DELETE FROM `users` WHERE `users`.`id` = 1
}

func userDemo(db *gorm.DB) {
	// 把模型和数据表对应起来
	db.AutoMigrate(&User{}) // 建立OR关系
	// createDemo(db) // 创建记录
	// readDemo(db) // 查询记录
	// updateDemo(db) // 修改记录
	deleteDemo(db) // 删除记录
}

func GormDemo() {
	println("gorm()")
	// 建立数据库链接
	s := "root:123@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		fmt.Println("err=", err)
		panic("连接数据库失败！")
	}
	fmt.Println("db=", db)
	// productDemo(db)
	userDemo(db)
}
