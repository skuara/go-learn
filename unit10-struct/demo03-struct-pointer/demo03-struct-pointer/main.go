package main

import "fmt"

func test1() {
	var name string
	name = "tom"
	fmt.Printf("name: %v\n", name)
	var p_name *string
	p_name = &name
	fmt.Printf("p_name: %v\n", p_name)
	fmt.Printf("p_name: %v\n", *p_name)
}

func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}
	p1 := Person{
		1,
		"tom",
		20,
	}
	// 定义一个结构体指针
	var p_Persion *Person
	fmt.Printf("p1: %p\n", &p1)
	p_Persion = &p1
	fmt.Printf("p_Persion: %p\n", p_Persion)
	// fmt.Printf("p_Persion.id: %p\n", p_Persion.id)   ???   p_Persion.id  int  not *int
	fmt.Printf("p_Persion: %v\n", *p_Persion)

	// var p p_Persion = &p1
	// fmt.Printf("p: %v\n", p)
	// fmt.Printf("p.id: %v\n", &p.id)
	// fmt.Printf("p: %v\n", *p)
}

func main() {
	test1()
	test2()

}
