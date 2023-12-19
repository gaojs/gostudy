package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func loginPostHandler(c *gin.Context) {
	// usr := c.PostForm("username")
	// pwd := c.PostForm("password")
	// usr, _ := c.GetPostForm("username")
	// pwd, _ := c.GetPostForm("password")
	usr := c.DefaultPostForm("username", "gao")
	pwd := c.DefaultPostForm("password", "***")
	c.JSON(http.StatusOK, gin.H{"用户名": usr, "密码": pwd})
}

func postHtml() {
	// 默认路由*Engine
	r := gin.Default()
	// 添加模板文件的解析
	r.LoadHTMLFiles("template/login.html")
	// http://localhost/login
	r.GET("/login", loginGetHandler)
	r.POST("/login", loginPostHandler)
	r.Run("localhost:80") // 默认是8080
}
func PostDemo() {
	postHtml()
}
