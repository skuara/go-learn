package main

import (
	"fmt"
)

// func parity_juge() {
// 	var num int
// 	fmt.Println("请输入一个数字")
// 	fmt.Scan(&num)
// 	if num%2 == 0 {
// 		fmt.Printf("num: %v 是偶数\n", num)
// 	} else {
// 		fmt.Printf("num: %v 是奇数\n", num)
// 	}
// }

// if else if
func grade() {
	grade := 80
	if grade >= 60 && grade <= 70 {
		fmt.Println("B")
	} else if grade > 70 {
		fmt.Println("A")
	} else {
		fmt.Println("C")
	}
}

func week() {
	// Monday  Tuesday Wednesday Thursday Friday Saturday Sunday
	var week string
	fmt.Println("请输入第一个字符")
	fmt.Scan(&week)
	if week == "S" {
		fmt.Println("请输入第二个字符")
		fmt.Scan(&week)
		if week == "a" {
			fmt.Println("Saturday")
		} else {
			fmt.Println("Sunday")
		}
	} else if week == "T" {
		fmt.Println("请输入第二个字符")
		fmt.Scan(&week)
		if week == "u" {
			fmt.Println("Tuesday")
		} else {
			fmt.Println("Thursday")
		}
	} else if week == "M" {
		fmt.Println("Monday")
	} else if week == "W" {
		fmt.Println("Wednesday")
	} else {
		fmt.Println("Friday")
	}
}

// if嵌套语句
func gender() {
	gender := "男"
	age := 20
	if gender == "男" {
		if age >= 18 {
			fmt.Println("成年男性")
		} else {
			fmt.Println("未成年男性")
		}
	} else {
		if age >= 18 {
			fmt.Println("成年女性")
		} else {
			fmt.Println("未成年女性")
		}
	}
}

func main() {
	a := 1
	b := 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	// 初始变量放在变量表达式中，age作用域只能在if的条件中的
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	flag1 := true
	if flag1 {
		fmt.Printf("flag1: %v\n", flag1)
	}
	// 注意不能使用0或者非0表示真假
	// 大括号必须存在即是只有一行
	// 左括号必须在if和else的同一行
	// parity_juge()
	grade()
	gender()
	// week()

}
