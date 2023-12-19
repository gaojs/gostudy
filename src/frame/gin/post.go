package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"usr"`
	Password string `form:"password" json:"pwd"`
}

func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginPost(c *gin.Context) {
	// usr := c.PostForm("username")
	// pwd := c.PostForm("password")
	// usr, _ := c.GetPostForm("username")
	// pwd, _ := c.GetPostForm("password")
	// usr := c.DefaultPostForm("username", "gao")
	// pwd := c.DefaultPostForm("password", "***")
	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"用户名": u.Username, "密码": u.Password})
	}
}

func uploadPost(c *gin.Context) {
	// 多文件上传，可以用for循环
	// form, _ := c.MultipartForm()
	// files := form.File["file"]
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.SaveUploadedFile(f, f.Filename)
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func postHtml() {
	// 默认路由*Engine
	r := gin.Default()
	// 添加模板文件的解析
	r.LoadHTMLFiles("template/login.html")
	// http://localhost/login
	r.GET("/login", loginGet)
	r.POST("/login", loginPost)
	r.POST("/upload", uploadPost)
	r.Run("localhost:80") // 默认是8080
}

func redirectBaidu(c *gin.Context) {
	// 直接重定向，URL地址栏都会跟着改变
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func redirectHello(c *gin.Context) {
	// c.Request.URL.Path = "/baidu"
	// gin.Default().HandleContext(c)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

func m1(c *gin.Context) {
	fmt.Println("m1 in")
	c.Set("usr", "hill")
	c.Next() // 调用后续的处理函数
	// c.Abort() // 不执行后续处理函数
	fmt.Println("m1 out")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in")
	u, err := c.Get("usr")
	fmt.Println(u, err)
	// c.Next() // 调用后续的处理函数
	c.Abort() // 不执行后续处理函数
	fmt.Println("m2 out")
}

func redirectDemo() {
	// r.Use(Logger(), Recovery())
	r := gin.Default() // gin.New()
	// http://localhost/hello
	// r.GET("/baidu", redirectBaidu)
	// r.GET("/hello", redirectHello)
	r.Use(m1, m2)
	r.GET("/both", hello)
	r.Run("localhost:80")
}

func PostDemo() {
	// postHtml()
	redirectDemo()
}
