package frame

import (
	"fmt"
	"main/frame/gin"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello HTTP!")
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
