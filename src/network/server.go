package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 延时关闭
	for {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取(Read)失败!err=", err)
			break
		}
		fmt.Printf("收到数据(len=%d):\n%s\n", n, string(buf[:n]))
		sent, err := conn.Write([]byte("received!"))
		fmt.Printf("发送数据(sent=%d)\n", sent)
	}
}

func serverDemo() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("监听(Listen)失败!err=", err)
		return
	}
	fmt.Println("监听(Listen)成功!", listen)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受(Accept)失败!err=", err)
			continue
		}
		fmt.Println("接受(Accept)成功!", conn)
		go process(conn)
	}
}

func main() {
	serverDemo()
}
