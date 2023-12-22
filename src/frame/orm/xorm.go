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
		// 添加了pk才能调用ID方法按主键查询
		n, err := engine.ID(1).Get(&u)
		fmt.Println("Get:", n, err, u)
	}
}

func relationDemo(engine *xorm.Engine) {
	// 同步表结构
	err := engine.Sync2(new(Teacher), new(Detail), new(Student), new(Course), new(Performance))
	fmt.Println("Sync, err=", err)
	if false { // 插入基础数据
		detail1 := &Detail{Id: 1, Tel: "139", Addr: "杭州"}
		detail2 := &Detail{Id: 2, Tel: "133", Addr: "南京"}
		stu1 := &Student{Id: 1, Name: "小明"}
		stu2 := &Student{Id: 2, Name: "小王"}
		stu3 := &Student{Id: 3, Name: "小张"}
		stu4 := &Student{Id: 4, Name: "小田"}
		tcher1 := &Teacher{Id: 1, Name: "李白", DetailId: 1}
		tcher2 := &Teacher{Id: 2, Name: "杜甫", DetailId: 2}
		course1 := &Course{Name: "剑术", TeacherId: 1}
		course2 := &Course{Name: "体操", TeacherId: 2}
		course3 := &Course{Name: "舞蹈", TeacherId: 2}
		perf1 := &Performance{CourseId: 1, StudentId: 1, Score: 100}
		perf2 := &Performance{CourseId: 1, StudentId: 2, Score: 60}
		perf3 := &Performance{CourseId: 1, StudentId: 3, Score: 80}
		n, err := engine.Insert(detail1, detail2, stu1, stu2, stu3, stu4)
		fmt.Println("Insert", n, err)
		n, err = engine.Insert(tcher1, tcher2, course1, course2, course3, perf1, perf2, perf3)
		fmt.Println("Insert", n, err)
	}
	if false { // 单表
		// 单表，单条查询
		s := &Student{Id: 1}
		s2 := &Student{}
		b, err := engine.Get(s)
		fmt.Println("Get", b, err, s)
		b, err = engine.ID(1).Get(s2)
		fmt.Println("ID.Get", b, err, s2)
		b, err = engine.Where("id=?", 1).Get(s2)
		fmt.Println("Where.Get", b, err, s2)
		// 软删除
		s3 := Student{Name: "小田"}
		n, err := engine.Delete(&s3)
		fmt.Println("Delete", n, err, s3)
		// 单表，多条查询
		sa := make([]Student, 0) // make slice
		err = engine.Find(&sa)
		fmt.Println("Find", err, sa)
	}
	if false { // 多表
		// 1对1查询（1教师：1教师详情）
		tcherDetails := make([]TeacherDetail, 0)
		err = engine.Table("teacher").Join("LEFT", "detail", "teacher.detail_id=detail.id").Find(&tcherDetails)
		fmt.Println("Find", err, tcherDetails)
		// 1对多查询（1个老师：N个课程；1个课程：1个老师）
		courseTchers := make([]CourseTeacher, 0)
		err = engine.Table("course").Join("LEFT", "teacher", "course.teacher_id=teacher.id").Find(&courseTchers)
		fmt.Println("Find", err, courseTchers)
		// 多对多查询（N课程：N学生）
		performs := make([]Performance, 0)
		err = engine.Table("performance").Join("LEFT", "course", "performance.course_id=course.id").
			Join("LEFT", "student", "performance.student_id=student.id").Find(&performs)
		fmt.Println("Find", err, performs)
	}
	if false { // Rows-Scan
		// 不定义自己的集合，使用rows的封装类型返回
		newStu := new(Student)
		rows, err := engine.Rows(newStu)
		if err != nil {
			fmt.Println("Rows", err, newStu)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(newStu)
			fmt.Println("Scan", err, newStu)
		}
		// 执行SQL
		sql1 := "select * from Detail where id = ?"
		queryRet, err := engine.Query(sql1, 1)
		fmt.Printf("Query, err=%v, type=%T\n", err, queryRet)
		for i, v := range queryRet { // []map[string][]uint8
			fmt.Println(i, v)
			for j, e := range v {
				fmt.Println(j, string(e))
			}
		}
		// 执行SQL
		sql2 := "insert into student (name) values (?)"
		ret, _ := engine.Exec(sql2, "小白")
		lastInsertId, _ := ret.LastInsertId()
		effectedRows, _ := ret.RowsAffected()
		fmt.Println("LastInsertId=", lastInsertId, "RowsAffected=", effectedRows)
	}
	if false { // 事件
		before := func(bean any) {
			fmt.Println("before", bean) // before &{0 }
		}
		after := func(bean any) {
			fmt.Println("after", bean) // after &{1 小明}
		}
		s := Student{}
		b, err := engine.Before(before).After(after).Get(&s)
		fmt.Println("bool=", b, "err=", err, "s=", s)
	}
	// 事务
	txExample(engine)
}

func txExample(engine *xorm.Engine) {
	fmt.Println("txExample: in")
	session := engine.NewSession()
	defer session.Close()

	//插入 教师详情
	detail := &Detail{
		Addr: "长安",
		Tel:  "999",
	}
	// insert后会更新 detail变量中的相关字段id
	_, err := session.Insert(detail)
	if err != nil { // 插入失败回滚
		fmt.Println("Insert, err=", err)
		session.Rollback()
		return
	}
	// 插入 教师
	tcher := &Teacher{
		Name:     "白居易",
		DetailId: detail.Id,
	}
	_, err = session.Insert(tcher)
	if err != nil { // 插入失败回滚
		fmt.Println("Insert, err=", err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		fmt.Println("Commit, err=", err)
		return
	}
	fmt.Println("txExample: out")
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
	// err = engine.Sync2(new(UserInfo)) // 同步表结构
	// fmt.Println("Sync, err=", err)
	// insertDemo(engine)
	// qureyDemo(engine)
	relationDemo(engine)
}
