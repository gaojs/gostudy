package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}
func delHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}
func putHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}
func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func Demo() {
	println("gin()")
	r := gin.Default() // 返回GIN引擎Engine
	// 增删改查
	r.POST("/hello", postHello)  // 增
	r.DELETE("/hello", delHello) // 删
	r.PUT("/hello", putHello)    // 改
	r.GET("/hello", getHello)    // 查
	// http://localhost/hello
	r.Run("localhost:80") // 默认是8080
}
