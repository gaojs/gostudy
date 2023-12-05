// init函数演示

package function

var globalVar = initGlobalVar()

func initGlobalVar() int {
	println("全局变量的初始化，先于init()函数执行")
	return 90
}

func init() {
	println("init()函数，先于main()函数执行")
	println("init1()")
}

func init() {
	println("init2()")
	println("同一个源文件中，可以定义多个init()函数")
}

func funcInit() {
	println("funcInit()")
}
