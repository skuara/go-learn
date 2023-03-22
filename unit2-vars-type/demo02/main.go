package main

import "fmt"

// 全局变量,定义在函数外的变量
var (
	m1 = 2
	m2 = 1.2
	m3 = "honey"
)

// 匿名变量 声明
func getNameage() (string, int) {
	return "tom", 20
}

func main() {
	// 声明+赋值
	var num int = 12
	fmt.Println("num =", num)
	// 声明 不赋值，取默认值
	var num2 int
	fmt.Println("num2 = ", num2)
	// 不声明类型，通过值推断类型
	var num3 = 13.52
	fmt.Println("num3 =", num3)
	// 省略var  短变量声明
	sex := "男"
	fmt.Println("sex =", sex)

	// 一次性声明和赋值多个变量
	var n1, n2, n3 int
	n1 = 1
	n2 = 2
	n3 = 3
	fmt.Println(n1, n2, n3)

	var n4, n5, n6 = 1, "man", 13.2
	fmt.Println(n4, n5, n6)

	n9, n7, n8 := 1, 2, "女"
	fmt.Println(n7, n8, n9)

	fmt.Println(m1, m2, m3)

	// _ 匿名变量
	_, age := getNameage()
	fmt.Printf("age: %v\n", age)

}
