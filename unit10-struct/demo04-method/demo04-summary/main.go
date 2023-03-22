package main

import "fmt"

type Persion struct {
	name string
}

func showPerson(p Persion) {
	fmt.Printf("p: %p\n", &p)
	p.name = "kite"
	fmt.Printf("p: %v\n", p)
}

func showPerson2(p *Persion) {
	fmt.Printf("p: %p\n", p)
	p.name = "kite"
	fmt.Printf("p: %v\n", *p)

}

// 方法的值类型和指针类型的接受者
func (per Persion) showPerson() {
	fmt.Printf("per: %p\n", &per)
	per.name = "kite"
	fmt.Printf("p: %v\n", per)
}
func (per *Persion) showPerson2() {
	fmt.Printf("per: %p\n", per)
	per.name = "kite"
	fmt.Printf("per: %v\n", *per)
}

func main() {
	p1 := Persion{
		"tom",
	}
	fmt.Printf("p1: %T\n", p1)
	p2 := &Persion{name: "tom"}
	fmt.Printf("p2: %T\n", p2)
	fmt.Println("-----------")
	// 值类型传递，相当于复制，原值不变
	fmt.Printf("p1: %v\n", &p1)
	showPerson(p1)
	fmt.Printf("p1: %p\n", &p1)
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("-----------")
	// 指针类型可以改变原值
	fmt.Printf("p2: %p\n", p2)
	showPerson2(p2)
	fmt.Printf("p2: %p\n", p2)
	fmt.Printf("p2: %v\n", *p2)
	fmt.Println("===========")
	//  值类型的接受者和指针类型的接受者
	fmt.Printf("p1: %p\n", &p1)
	p1.showPerson()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("-----------")
	fmt.Printf("p2: %p\n", p2)
	p2.showPerson2()
	fmt.Printf("p2: %v\n", *p2)

}
