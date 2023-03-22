package main

import (
	"fmt"
)

// 函数作为参数

func sayhello(name string) {
	fmt.Printf("%v say hello!\n", name)
}
func qucan(name string) string {
	return name
}
func test(name string, f func(string)) {
	f(name)
}

// 函数作为返回值

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func cal(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	sayhello(qucan("Tom"))
	test("Alice", sayhello)

	f := cal("+")
	fmt.Printf("f(2, 3): %v\n", f(2, 3))
	f1 := cal("-")
	fmt.Printf("f1(2, 1): %v\n", f1(2, 1))

}
