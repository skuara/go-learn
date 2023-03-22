package main

import (
	"fmt"
)

// 函数的参数
// 类型必须指定
func f1(a int, b int) int { // 形参
	return a + b
}

// 声明函数的参数列表叫形参， 调用时传递的参数叫实参
// go通过传值的方式传参，传递给函数的是拷贝后的副本，函数内部访问，修改的是副本
func f2(a int) {
	a = 200
	fmt.Printf("a: %v\n", a)
}

// 需要注意 数组，切片，指针本身就是指针类型，拷贝传值也是拷贝指针，拷贝后仍然是底层数据结构，修改他们可能会影响外部结构的值
func f3(a []int) {
	a[0] = 2
	a[1] = 3
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %x\n", &a[0])
}

// 支持变长参数
func f4(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// r := f1(2, 3) // 实参
	// fmt.Printf("r: %v\n", r)
	// a := 100
	// f2(a)
	// fmt.Printf("a: %v\n", a)
	// a := []int{3, 4}
	// fmt.Printf("a: %x\n", &a[0])
	// f3(a)
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("a: %x\n", &a[0])
	f4(1, 2, 345, 222)

}
