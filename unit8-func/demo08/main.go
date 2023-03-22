package main

import "fmt"

// 递归   函数内部调用函数自身就是递归
// 1. 递归就是自己调用自己
// 2. 必须先定义函数的退出条件，没有退出条件，递归将成为死循环
// 3. go语言递归函数很有可能会产生一大堆goroutine,也很可能会出现栈空间内存溢出的问题

// for循环方式
func f1() {
	a := 1
	// a=5!=5*4*3*2*1
	for i := 1; i <= 5; i++ {
		a *= i
	}
	fmt.Println(a)
}

// 递归的方式
// 阶乘
func f2(a int) int {
	if a == 1 {
		return 1
	} else {
		a = a * f2(a-1)
		return a
	}
}

// 斐波那契数列
// f(n)=f(n-1)+f(n-2) 且f(2)=f(1)=1
func f3(n int) int {
	if n == 2 || n == 1 {
		return 1
	}
	return f3(n-1) + f3(n-2)
}

func main() {
	f1()
	fmt.Printf("f2(5): %v\n", f2(5))
	fmt.Printf("f3(5): %v\n", f3(5))
}
