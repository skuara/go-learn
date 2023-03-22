package main

import "fmt"

// 类型断言

func assertType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d\n", v)
	case float64:
		fmt.Printf("%f\n", v)
	case float32, byte:
		fmt.Printf("%T  %b\n", v, v)
	}
}

func main() {

	var i interface{}
	var a int
	var b float64
	var c byte
	i = a
	assertType(i)
	i = b
	assertType(b)
	i = c
	assertType(c)
	var d float32
	i = d

	if i, ok := i.(float32); ok {
		fmt.Printf("i 是 float32类型,其值为%f\n", i)
	} else {
		fmt.Printf("i 不是float32类型")
	}

}
