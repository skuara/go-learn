package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

// 值接收者实现方法
func (p Person) working() {
	fmt.Printf(" %v is working \n", p.name)
}

// 指针接收者实现方法
func (p *Person) change(mail string) {
	p.email = mail
}

// 值接收者使用值的副本来调用方法，指针接收者使用实际值来调用方法
// when  值接收者 or 指针接收者
// 类型的本质，  给当前类型增加或者删除一个值时，是打算创建一个新的值还是修改原来的值
// 前者通过值类型  后者通过指针类型
func main() {
	p1 := Person{
		"tom",
		20,
		"tom@email.com",
	}
	p1.working()
	p1.change("tom_33@email.com")
	fmt.Printf("p1: %v\n", p1)
}
