package main

import "fmt"

// 匿名函数，go语言函数不能嵌套，可用匿名函数简单实现  可以既没有参数，也没有返回值

func main() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	i := max(1, 3)
	fmt.Printf("i: %v\n", i)
	// 也可以自己调用自己
	r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Printf("r: %v\n", r)

}
