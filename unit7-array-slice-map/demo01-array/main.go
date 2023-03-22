package main

import "fmt"

func main() {
	// var a [3]int
	// var b [2]string
	// var c [2]bool
	// fmt.Printf("a: %T\n", a)
	// fmt.Printf("b: %T\n", b)
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("c: %v\n", c)

	// 初始化列表
	var a = [2]int{1, 2}
	fmt.Printf("a: %v\n", a)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)

	// 数组初始化 默认省略长度 自动推断长度
	var b = [...]int{1, 2, 3}
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %T\n", b)
	var b1 = [...]string{"hello", "world"}
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("len(b1): %v\n", len(b1))

	// 指定索引值的方式来初始化
	var c = [...]int{1: 2, 3: 2}
	fmt.Printf("c: %v\n", c)
	fmt.Printf("len(c): %v\n", len(c))
	var c1 = [...]string{1: "tom", 2: "jack"}
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("len(c1): %v\n", len(c1))
	var c2 = [...]bool{1: true, 2: false}
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("len(c2): %v\n", len(c2))

	//
}
