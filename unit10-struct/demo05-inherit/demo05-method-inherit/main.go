package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println("eating....")
}
func (animal Animal) sleep() {
	fmt.Println("sleeping...")
}

type Cat struct {
	Animal
	kind string
}

func main() {
	cat := Cat{
		Animal{
			name: "buji",
			age:  20,
		},
		"buou",
	}
	cat.eat()
	cat.sleep()
	fmt.Printf("cat.kind: %v\n", cat.kind)
	fmt.Printf("cat.age: %v\n", cat.age)

}
