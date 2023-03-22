package main

import "fmt"

type dog struct {
	name string
	age  int
}
type person struct {
	name string
	age  int
	dog  dog
}

func main() {
	var p person
	p.age = 20
	p.name = "tom"
	p.dog.age = 1
	p.dog.name = "lucky"
	fmt.Printf("p: %v\n", p)

}
