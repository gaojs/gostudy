package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func clientDemo() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("拨号(Dial)失败!err=", err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入发送的内容:")
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToLower(s) == "q" {
			break
		}
		n, err := conn.Write([]byte(s))
		if err != nil {

			fmt.Println("发送(Write)失败!err=", err)
			return
		}
		fmt.Printf("已发送(len=%d):\n%s\n", n, s)
	}
}

func main() {
	clientDemo()
}
