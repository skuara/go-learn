package main

import "fmt"

func f2() {
	// 初始条件写在外边
	i := 1
	for ; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3() {
	// 初始条件和改变条件
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

// 永真循环
func f4() {
	for {
		fmt.Println("我在循环")
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	f2()
	f3()
	f4()
}
