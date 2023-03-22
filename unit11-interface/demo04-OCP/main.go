package main

import "fmt"

// OCP设计原则   对拓展开放，对修改关闭  go不是面向对象的语言

type Pet interface {
	eat()
	sleep()
}

type Cat struct {
	name string
	age  int
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eating...")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleeping...")
}

func (cat Cat) eat() {
	fmt.Println("cat eating....")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleeping...")
}

type Person struct {
}

func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {

	var cat Pet
	cat = Cat{"tom", 2}
	var person Person
	person.care(cat)
	var dog Pet
	dog = Dog{"jeery", 3}
	person.care(dog)

}
