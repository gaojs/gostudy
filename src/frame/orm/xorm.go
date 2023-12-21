package orm

import (
	"fmt"

	"xorm.io/xorm"
)

func insertDemo(engine *xorm.Engine) {
	engine.ShowSQL(true)
	u := User{Name: "hill", Age: 20}
	n, err := engine.Insert(&u)
	fmt.Println("Insert:", n, err, u)
	// [SQL] INSERT INTO `user` (`id`,`name`,`age`) VALUES (?,?,?) [0 hill 20]
}

func qureyDemo(engine *xorm.Engine) {
	if false { // 查看表信息
		ts, _ := engine.DBMetas()
		for _, v := range ts {
			t, _ := engine.TableInfo(v)
			fmt.Println(t)
		}
	}
	if false { // 查询
		engine.ShowSQL(true)
		u := User{Id: 2}
		// has, err := engine.Get(&u)
		// fmt.Println("Get", has, err, u)
		// [SQL] SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [2]
		has, err := engine.Exist(&u)
		fmt.Println("Exist", has, err, u)
		// [SQL] SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [2]
	}
	if false { // 查询
		r, err := engine.Query("select * from user")
		fmt.Println("Query:", r, err)
		rs, err := engine.QueryString("select * from user")
		fmt.Println("QueryString:", rs, err)
		ri, err := engine.QueryInterface("select * from user")
		fmt.Println("QueryInterface:", ri, err)
	}
	if false { // 更新
		u := User{Name: "new2"}
		n, err := engine.Update(&u)
		fmt.Println("Update:", n, err)
	}
	if true { // 删除
		u := User{Name: "hill"}
		n, err := engine.Delete(&u)
		fmt.Println("Delete:", n, err)
		// [SQL] DELETE FROM `user` WHERE `name`=? [hill]
	}
}

func XormDemo() {
	println("xorm()")
	// 建立数据库链接
	s := "root:123@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", s)
	if err != nil {
		fmt.Println("err=", err)
		panic("连接数据库失败！")
	}
	defer engine.Close()
	fmt.Println("engine=", engine.DriverName())
	err = engine.Sync2(new(User)) // 同步表结构
	fmt.Println("Sync, err=", err)
	insertDemo(engine)
	qureyDemo(engine)
}
