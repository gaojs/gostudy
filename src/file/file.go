package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func fileOpen() {
	var f *os.File // 文件句柄
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer f.Close() // 延迟关闭
	fmt.Printf("文件指针=%v, %T\n", f, f)
}

func fileReader() {
	var f *os.File // 文件句柄
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer f.Close() // 延迟关闭
	fmt.Println("文件内容如下：")
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // EndOfFile
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件内容读取结束！")
}

func fileReadAll() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println("打开读取错误：", err)
		return
	}
	fmt.Println("文件内容如下：")
	fmt.Println(string(content))
}

func fileWriter() {
	file, err := os.OpenFile("1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer file.Close()
	// 带缓存的Writer，需要刷新
	writer := bufio.NewWriter(file)
	writer.WriteString("hello")
	writer.Flush() // 必须刷新才存盘
	fmt.Println("文件内容写入结束！")
}

func fileExist() {
	info, err := os.Stat("main.go")
	if err != nil {
		fmt.Println("查看文件属性错误：", err)
		return
	}
	fmt.Println("文件名：", info.Name())
	fmt.Println("文件大小：", info.Size())
	fmt.Println("文件模式：", info.Mode())
	fmt.Println("修改时间：", info.ModTime())
	fmt.Println("文件信息：", info.Sys())
}

func Demo() {
	// fileOpen()
	// fileReader()
	// fileReadAll()
	// fileWriter()
	fileExist()
}
