package main

import "fmt"

// init 函数    先于main函数执行，实现包级别的一些初始化操作

// init 函数不能被其他函数调用 既没有输入参数也没有返回值，每个包可以有多个init函数，包的每个源文件也可以有多个init函数
// 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个顺序执行
// 不同包的init函数按照包的导入的依赖关系决定执行顺序
// 初始化顺讯   变量的初始化--》 init初始化 --》 main函数

var i int = initvar()

func init() {
	fmt.Println("init...")
}

func initvar() int {
	fmt.Println("initvar...")
	return 100
}

func main() {
	fmt.Println("main...")

}
