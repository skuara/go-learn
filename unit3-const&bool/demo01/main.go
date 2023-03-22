package main

import "fmt"

func main() {
	// const constantName[type]= value
	const PI float64 = 36
	const PI2 = 3.14
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2)

	const (
		a = 100
		b = 200
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	const w, h = 100, 300
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)

	// iota  默认值0， 每声明一个常量增加1，遇到const关键字 重置为0
	const (
		a1 = iota
		_
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}
