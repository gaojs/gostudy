package gin

import (
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
func PostDemo() {
	postHtml()
}
