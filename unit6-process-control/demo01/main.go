package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 20
	if a > 20 {
		fmt.Printf("a: %v\n", a)
	} else {
		fmt.Printf("b: %v\n", b)
	}
	switch a {
	case 10:
		fmt.Printf("a: %v\n", a)
	default:
		fmt.Println("default")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	x := [...]int{1, 2, 3}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
}
