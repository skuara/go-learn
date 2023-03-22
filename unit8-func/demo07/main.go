package main

import (
	"fmt"
	"strings"
)

// closure  闭包
// 定义在一个函数内部的函数，本质上闭包是将函数内部和函数外部连接起来的桥梁  闭包=函数+引用环境

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name = name + suffix
			return name
		} else {
			return name
		}
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// f := add()
	// // 变量f是一个函数并且引用了外部变量x，此时x就是一个闭包，在f的生命周期内，变量x也一直有效
	// f(10)
	// f(20)
	// fmt.Println("-----------")
	// f1 := add()
	// f1(10)
	// f1(20)
	jpgSuffix := makeSuffixFunc(".jpg")
	txtSuffix := makeSuffixFunc(".txt")
	fmt.Printf("jpgSuffix(\"test\"): %s\n", jpgSuffix("test"))
	fmt.Printf("txtSuffix(\"test\"): %s\n", txtSuffix("test"))

	add, sub := calc(10)
	fmt.Println(add(1), sub(2))

}
