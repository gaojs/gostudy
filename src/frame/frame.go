package frame

import (
	"fmt"
	"io/ioutil"
	"main/frame/gin"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>Hello HTTP!</h1>")
	b, err := ioutil.ReadFile("hello.html")
	if err != nil {
		fmt.Println("ReadFile failed, err=", err)
		return
	}
	n, err := fmt.Fprintln(w, string(b))
	if err != nil {
		fmt.Println("Fprintln failed, err=", err)
		return
	}
	fmt.Println("Fprintln len=", n)
}

func httpDemo() {
	// http://localhost/hello
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("http serve failed, err=", err)
		return
	}
}

func Demo() {
	println("frame()")
	gin.Demo()
	httpDemo()
}
