package main

import "fmt"

// 函数类型与函数变量

type fun func(int, int) int

func sum(a int, b int) int {
	return a + b
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	var f fun
	f = sum
	s := f(2, 3)
	fmt.Printf("s: %v\n", s)
	f = max
	m := f(2, 3)
	fmt.Printf("m: %v\n", m)

}
