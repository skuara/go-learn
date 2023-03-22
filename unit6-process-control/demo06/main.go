package main

import "fmt"

// 1. 单独在select使用break和不使用break没有啥区别
// 2. 单独在switch语句，并且没有fallthrough，使用break和不使用break没有啥区别
// 3. 单独在表达式switch语句，并且有fallthrough，使用break能够终止fallthrough后的case语句的执行
// 4. 带标签的break可以跳出多层select和switch作用域范围，让break更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带break只能跳出当前语句块

func f1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

func f2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		// break
		fallthrough
	case 3:
		fmt.Println("3")
		break
	}
}

// break 跳出循环
func f3() {
MYLABLE:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break MYLABLE
		}
	}
	fmt.Println("end")
}

func main() {
	f1()
	f2()
	f3()
}
