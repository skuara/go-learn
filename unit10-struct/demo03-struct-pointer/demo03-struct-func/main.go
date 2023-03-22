package main

import "fmt"

// 结构体作为函数参数
type Person struct {
	id   int
	name string
}

// 1. 值传递  函数内部不会改变外面结构体的内容
func showPerson(p1 Person) {
	p1.id = 1
	p1.name = "tom"
	fmt.Printf("p1: %v\n", p1)
}

// 2. 指针传递   函数内部能够改变外部结构体的内容
func showPersonPtr(p1 *Person) {
	p1.id = 1
	p1.name = "tom"
	fmt.Printf("p1: %v\n", *p1)
}

func main() {
	var p1 Person = Person{
		1,
		"kite",
	}
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("----------------")
	// 相当于复制了一个新的结构体，原来的结构体内容不变
	showPerson(p1)
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("==============")
	// 指针传递调用后 值已经被修改了
	showPersonPtr(&p1)
	fmt.Println("-----------")
	fmt.Printf("p1: %v\n", p1)

}
