package main

import "fmt"

func grade() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("其他")
	}
}

// 多条件匹配
func week() {
	week := 10
	switch week {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("其他")
	}
}

// case可以是条件表达式
func age() {
	age := 80
	switch {
	case age >= 80:
		fmt.Println("老年")
	case age < 80 && age >= 30:
		fmt.Println("中年")
	default:
		fmt.Println("青年")
	}
}

//fallthrough 可以执行满足条件的下一个case
func f1() {
	num := 100
	switch num {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}

func main() {
	grade()
	week()
	age()
	f1()
}
