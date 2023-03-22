package main

import "fmt"

// 遍历数组
func f1() {
	var v = [5]int{1, 2, 3, 4, 5}
	for i, v := range v {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
		fmt.Println("--------")
	}
}

// 遍历切片
func f2() {
	var a = []int{1, 2, 3}
	for _, a := range a {
		fmt.Printf("a: %v\n", a)
	}
}

// 遍历map
func f3() {
	var m = make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["gender"] = "man"
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}

// 遍历字符串
func f4() {
	s := "hello,world"
	for _, s := range s {
		fmt.Printf("s: %c\n", s)
	}
}

func main() {
	// for range 遍历数组，切片，字符串，map以及通道
	// 通过for range遍历返回有以下规律
	// 数组，切片，字符串返回索引和值
	// map返回键和值
	// 通道只返回通道的值
	f1()
	f2()
	f3()
	f4()
}
