// 序列化(反序列化)
package serialization

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func argsDemo() {
	var usr, pwd string
	fmt.Println("命令行参数：", os.Args)
	flag.StringVar(&usr, "u", "user", "用户名，默认为user")
	flag.StringVar(&pwd, "pwd", "pwd", "密码，默认为pwd")
	flag.Parse() // 解析os.Args
	fmt.Println("用户名：", usr)
	fmt.Println("密码：", pwd)
}

func serialize(o any) {
	js, err := json.Marshal(o)
	if err != nil {
		fmt.Println("序列化错误：", err)
		return
	}
	fmt.Println("序列化结果：", string(js))
}

func serialDemo() {
	m := make(map[string]interface{})
	m["name"] = "高"
	m["age"] = 20
	m["addr"] = "杭州"
	serialize(m) // Map序列化
	var s []map[string]interface{}
	s = append(s, m)
	serialize(s) // 切片序列化
}
func unserialDemo() {
	m := make(map[string]interface{})
	m["name"] = "高"
	m["age"] = 20
	m["addr"] = "杭州"
	js, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化错误：", err)
		return
	}
	var n any
	err2 := json.Unmarshal(js, &n)
	if err2 != nil {
		fmt.Println("反序列化错误：", err2)
		return
	}
	fmt.Println("反序列化结果：", n)
}

func Demo() {
	argsDemo()
	serialDemo()
	unserialDemo()
}
