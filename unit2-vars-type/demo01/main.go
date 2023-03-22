package main

import "fmt"

func main() {
	//1. 声明
	var age int
	//2. 赋值
	age = 18
	//3. 输出
	fmt.Println("age =", age)
	// 声明赋值可以合成
	var age2 int = 19
	fmt.Println("age2 =", age2)
	// 变量不可以重复赋值， 赋值的时候不可以赋予不匹配的类型
}
