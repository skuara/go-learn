package main

import (
	"fmt"
)

// 方法   有接收者的函数
// 接受者并非一定是struct类型，也可以是类型别名，slice，map，channel，func类型都可以
// struct结合他的类型就相当于面向对象中的类，区别是struct可以和它的方法分开，并非一定要属于同一个文件，但一定要属于同一个包
// 方法有两种接收类型(T type）和(T *type)  他们之间有区别
// 方法本身就是一个函数，go语言中没有方法重载，也就是说同一个类型中所有的方法名都必须是唯一的
// 如果receiver是一个指针类型，则会自动解除引用
// 方法和type是分开的，意味着实例的行为和数据存储是分开的，他们通过receiver建立起关联关系。

type person struct {
	name string
}

func (per person) eat() {
	fmt.Println(per.name + " eating ")
}

func (per person) sleep() {
	fmt.Println(per.name + " sleeping")
}

type custome struct {
	name string
}

func (cus custome) load(pwd string) bool {
	if cus.name == "tom" && pwd == "123456" {
		return true
	} else {
		return false
	}
}

func main() {
	var p person
	p.name = "tom"
	p.eat()
	p.sleep()

	var c custome = custome{
		name: "tom",
	}
	fmt.Printf("c.load(\"123456\"): %v\n", c.load("123456"))

}
