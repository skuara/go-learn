package main

import "fmt"

func main() {
	// 算术运算符
	a := 100
	b := 10
	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))

	a++
	fmt.Printf("a: %v\n", a)
	b--
	fmt.Printf("b: %v\n", b)

	// 关系运算符
	fmt.Printf("(a > b): %v\n", (a > b))
	fmt.Printf("(a < b): %v\n", (a < b))
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a != b): %v\n", (a != b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b))

	// 逻辑运算符  && ||  !
	c := true
	d := false
	fmt.Printf("(c && d): %v\n", (c && d))
	fmt.Printf("(c || d): %v\n", (c || d))
	fmt.Printf("(!c): %v\n", (!c))
	fmt.Printf("(!d): %v\n", (!d))

	// 位运算符
	e := 4
	f := 8
	fmt.Printf("e: %b\n", e)
	fmt.Printf("f: %b\n", f)

	fmt.Printf("(e & f): %b\n", (e & f))   // 按位相与
	fmt.Printf("(e | f): %b\n", (e | f))   // 按位相或
	fmt.Printf("(e ^ f): %b\n", (e ^ f))   // 按位相异或，二进制位数相异时，结果为1
	fmt.Printf("(e << 2): %b\n", (e << 2)) // 二进制左移动
	fmt.Printf("(f >> 2): %b\n", (f >> 2)) // 二进制右移动

	// 赋值运算符
	g := 100
	g = 10
	fmt.Printf("g: %v\n", g)
	a += 1
	fmt.Printf("a: %v\n", a)
	a &= 1
	fmt.Printf("a: %v\n", a)
	a = 2
	a ^= 1
	fmt.Printf("a: %b\n", a)

}
