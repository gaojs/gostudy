package gin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

func ginRest() {
	r := gin.Default() // 返回GIN引擎Engine
	// 增删改查
	r.POST("/hello", postHello)  // 增
	r.DELETE("/hello", delHello) // 删
	r.PUT("/hello", putHello)    // 改
	r.GET("/hello", getHello)    // 查
	// http://localhost/hello
	r.Run("localhost:80") // 默认是8080
}

type User struct {
	Name   string
	Gender string
	Age    int
}

func helloTmpl(w http.ResponseWriter, r *http.Request) {
	// 1、定义模板(hello.tmpl)
	kua := func(name string) (string, error) {
		return name + "真帅气", nil
	}
	// 2、解析模板
	t := template.New("hello.tmpl")
	// 添加自定义函数，模板中可使用
	t.Funcs(template.FuncMap{"kua": kua})
	t, err := t.ParseFiles("hello.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed, err=", err)
		return
	}
	// 3、渲染模板
	// err = t.Execute(w, "template")
	usr := User{"Hill", " 男 ", 18}
	hbs := []string{"足球", "篮球", "乒乓球"}
	mp := map[string]any{
		"user":  usr,
		"tall":  177,
		"hobby": hbs,
	}
	err = t.Execute(w, mp)
	if err != nil {
		fmt.Println("Execute failed, err=", err)
		return
	}
}

func indexTmpl(w http.ResponseWriter, r *http.Request) {
	// 1、定义模板(.tmpl)
	// 2、解析模板
	t, err := template.ParseFiles("template/base.tmpl", "template/index.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed, err=", err)
		return
	}
	// 3、渲染模板
	err = t.Execute(w, "index")
	if err != nil {
		fmt.Println("Execute failed, err=", err)
		return
	}
}

func homeTmpl(w http.ResponseWriter, r *http.Request) {
	// 1、定义模板(.tmpl)
	// 2、解析模板
	t, err := template.ParseFiles("template/base.tmpl", "template/home.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed, err=", err)
		return
	}
	// 3、渲染模板
	err = t.Execute(w, "home")
	if err != nil {
		fmt.Println("Execute failed, err=", err)
		return
	}
}

func customTmpl(w http.ResponseWriter, r *http.Request) {
	// 1、定义模板(.tmpl)
	// 2、解析模板
	t := template.New("custom.tmpl").Delims("{[", "]}")
	t, err := t.ParseFiles("template/custom.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed, err=", err)
		return
	}
	// 3、渲染模板
	xss := "<script>alert('xss');</script>"
	err = t.Execute(w, xss)
	if err != nil {
		fmt.Println("Execute failed, err=", err)
		return
	}
}

func httpDemo() {
	// http://localhost/hello
	http.HandleFunc("/hello", helloTmpl)
	http.HandleFunc("/index", indexTmpl)
	http.HandleFunc("/home", homeTmpl)
	http.HandleFunc("/custom", customTmpl)
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		fmt.Println("http server start failed, err=", err)
		return
	}
}

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.tmpl",
		gin.H{"title": "<a href='http://baidu.com'>Baidu</a>"})
}

func jsonHandler(c *gin.Context) {
	data := gin.H{"name": "Hill", "age": 18}
	c.JSON(http.StatusOK, data)
}

func structHandler(c *gin.Context) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	data := Person{Name: "Hill", Age: 18}
	c.JSON(http.StatusOK, data)
}

func queryHandler(c *gin.Context) {
	// name := c.Query("name")
	name := c.DefaultQuery("name", "default")
	age, ok := c.GetQuery("age")
	if !ok { // 取不到，则设置
		age = strconv.Itoa(20)
	}
	data := gin.H{"name": name, "age": age}
	c.JSON(http.StatusOK, data)
}

func pathHandler(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	data := gin.H{"name": name, "age": age}
	c.JSON(http.StatusOK, data)
}

func ginHtml() {
	// 默认路由*Engine
	r := gin.Default()
	// 添加静态文件的解析
	r.Static("/static", "template/static")
	// 添加自定义函数（safe）
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 添加模板文件的解析
	// r.LoadHTMLFiles("template/hello.tmpl")
	r.LoadHTMLGlob("template/hello.*")
	// http://localhost/hello
	// r.GET("/hello", sayHello)
	// r.GET("/json", jsonHandler)
	// r.GET("/struct", structHandler)
	// http://localhost/query?name=gao&age=18
	// r.GET("/query", queryHandler)
	// http://localhost/hill/18
	r.GET("/:name/:age", pathHandler)
	r.Run("localhost:80") // 默认是8080
}

func GetDemo() {
	println("gin()")
	// httpDemo()
	// ginRest()
	ginHtml()
}
