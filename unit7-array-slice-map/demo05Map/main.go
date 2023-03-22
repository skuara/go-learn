package main

import (
	"fmt"
)

func f1() {
	// 声明一个未初始化的映射
	var m map[string]string
	// 初始化映射
	m = make(map[string]string)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("m: %T\n", m)

	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "18"
	m1["email"] = "tom@gmail.com"
	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"name":  "kaili",
		"age":   "20",
		"email": "firefox@gmail.com",
	}
	fmt.Printf("m2: %v\n", m2)
}

// 判断某个key是否存在
func f2() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "test@gmail.com"
	value, ok := m1["name"]
	fmt.Printf("value: %v\n", value)
	fmt.Printf("ok: %v\n", ok)
}

// map遍历
func f3() {
	m := map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	// for k := range m {
	// 	fmt.Printf("k: %v\n", k)
	// }
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
	// 从映射中删除一项
	delete(m, "age")
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
