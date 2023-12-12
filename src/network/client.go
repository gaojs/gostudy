package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func sendData(conn net.Conn, s string) {
	n, err := conn.Write([]byte(s))
	if err != nil {

		fmt.Println("发送(Write)失败!err=", err)
		return
	}
	fmt.Printf("已发送(len=%d):\n%s\n", n, s)
}

func recvData(conn net.Conn) {
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("读取(Read)失败!err=", err)
		return
	}
	fmt.Printf("收到数据(len=%d):\n%s\n", n, string(buf[:n]))
}

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
		sendData(conn, s)
		recvData(conn)
	}
}

func main() {
	clientDemo()
}
