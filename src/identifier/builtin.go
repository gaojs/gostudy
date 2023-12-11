// 内置函数
package identifier

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func stringDemo() {
	// 10进制转换成其他进制（2,8,16）
	fmt.Println("10进制转换成其他进制（2,8,16）")
	fmt.Println(strconv.FormatInt(123, 2))
	fmt.Println(strconv.FormatInt(123, 8))
	fmt.Println(strconv.FormatInt(123, 16))
	// 字符串替换，字符串分割
	fmt.Println("字符串替换，字符串分割")
	s := "我爱北京，我爱杭州，我爱地球！"
	fmt.Println("原字符串：", s)
	s2 := strings.Replace(s, "我", "你", 2)
	fmt.Println("替换后的：", s2)
	ss := strings.Split(s, "，")
	fmt.Println("分割后的：", len(ss), ss)
}

func timeDemo() {
	var now time.Time = time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Weekday(), now.YearDay())
	fmt.Println(now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Clock())
	fmt.Printf("当前日期和时间：%d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	// 使用Time.Format必须使用如下这个固定的日期和时间
	fmt.Println(now.Format("时间：2006-01-02 15:04:05"))
	for i := 0; i < 10; i++ { // 休眠和时间戳
		time.Sleep(time.Second + time.Millisecond*2)
		fmt.Println(time.Now().Unix(), time.Now().UnixMicro())
	}
}

func recoverDemo() {
	// 捕捉Panic错误
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func divideZero() {
	defer recoverDemo()
	var a int = 1
	var b = 0
	res := a / b
	println("res=", res, "未打印出来")
}

func readConf(conf string) error {
	if conf == "init.conf" {
		return nil
	}
	return errors.New("配置文件错误！")
}

func panicDemo() {
	err := readConf("conf.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("未产生Panic！")
}

func builtinDemo() {
	// stringDemo()
	// timeDemo()
	divideZero()
	println("builtin: len, new, make...")
	defer recoverDemo()
	panicDemo()
	fmt.Println("Panic之后！未打印出来")
}
