package gin

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello GIN!"})
}

func Demo() {
	println("gin()")
	// http://localhost/hello
	r := gin.Default() // 返回GIN引擎
	r.GET("/hello", sayHello)
	r.Run(":80") // 默认8080
}
