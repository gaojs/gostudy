package orm

import (
	"fmt"

	"xorm.io/core"
	"xorm.io/xorm"
)

func insertDemo(engine *xorm.Engine) {
	u := UserInfo{Name: "hill", Male: true, Addr: Address{"ZJ", "杭州"}}
	fmt.Println("insertDemo:", u)
	n, err := engine.Insert(&u)
	fmt.Println("Insert:", n, err, u)
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
		u := UserInfo{ID: 2}
		// has, err := engine.Get(&u)
		// fmt.Println("Get", has, err, u)
		// [SQL] SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [2]
		has, err := engine.Exist(&u)
		fmt.Println("Exist", has, err, u)
		// [SQL] SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [2]
	}
	if false { // 查询Query
		r, err := engine.Query("select * from user")
		fmt.Println("Query:", r, err) // 返回map[string][]byte
		rs, err := engine.QueryString("select * from user")
		fmt.Println("QueryString:", rs, err) // 返回map[string]string
		ri, err := engine.QueryInterface("select * from user")
		fmt.Println("QueryInterface:", ri, err) // 返回map[string]any
	}
	if false { // 更新
		u := UserInfo{Name: "new2"}
		n, err := engine.Update(&u)
		fmt.Println("Update:", n, err)
	}
	if false { // 删除
		u := UserInfo{Name: "hill"}
		n, err := engine.Delete(&u)
		fmt.Println("Delete:", n, err)
		// [SQL] DELETE FROM `user` WHERE `name`=? [hill]
	}
	if true { // 查询
		u := UserInfo{}
		n, err := engine.ID(1).Get(&u)
		fmt.Println("Get:", n, err, u)
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
	engine.ShowSQL(true) // 显示SQL
	engine.SetMapper(core.GonicMapper{})
	fmt.Println("engine=", engine.DriverName())
	err = engine.Sync2(new(UserInfo)) // 同步表结构
	fmt.Println("Sync, err=", err)
	insertDemo(engine)
	qureyDemo(engine)
}
