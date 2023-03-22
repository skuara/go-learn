package main

import "fmt"

// 函数的返回值

func no_return() {
	fmt.Println("just so so ")
}

func sum(a int, b int) int {
	return a + b
}

func mult_return() (name string, age int) {
	name = "tom"
	age = 20
	// return name, age
	return
}

func mult_return2() (name string, age int) {
	n := "tom"
	a := 20
	return n, a
}

// 丢掉返回值
func drop_return() (name string, age int) {
	name = "Alex"
	age = 12
	return
}

func main() {
	// no_return()
	// fmt.Printf("sum(2, 3): %v\n", sum(2, 3))
	fmt.Println(mult_return())
	fmt.Println(mult_return2())
	_, a := drop_return()
	fmt.Printf("a: %v\n", a)

}
