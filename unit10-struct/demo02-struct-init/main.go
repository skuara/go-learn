package main

import "fmt"

// go 没有面向对象的概念，可以用结构体来实现

type person struct {
	id, age     int
	name, email string
}

func main() {
	var tom person
	tom.id = 101
	tom.age = 20
	tom.name = "tom"
	tom.email = "tom@gmail.com"
	fmt.Printf("tom: %v\n", tom)
	kite := person{}
	fmt.Printf("kite: %v\n", kite)

	// 匿名结构体
	var dog struct {
		age  int
		name string
	}
	dog.age = 1
	dog.name = "buding"
	fmt.Printf("dog: %v\n", dog)

	// 结构体键值初始化
	pite := person{
		id:    1,
		age:   20,
		name:  "pite",
		email: "pite@gmail.com",
	}
	fmt.Printf("pite: %v\n", pite)
	// 结构体顺序初始化
	jerry := person{
		2,
		33,
		"jerry",
		"jerry@gmail.com",
	}
	fmt.Printf("jerry: %v\n", jerry)

	// 部分初始化
	bob := person{
		name: "bob",
		age:  20,
	}
	fmt.Printf("bob: %v\n", bob)

}
