package main

import "fmt"

func f1() {
	var a1 [2]int
	a1[0] = 20
	a1[1] = 30
	fmt.Printf("a1[1]: %v\n", a1[1])
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Println("------------")
	a1[0] = 200
	a1[1] = 300
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
}

// 数组的遍历
func f2() {
	// 1. 通过索引的方式遍历
	var a = [...]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]: %v\n", i, a[i])
	}
	// 2. for range
	for i, v := range a {
		fmt.Printf("a[%v]: %v\n", i, v)
	}
}

// 数组的赋值操作
func copy() {
	var array1 [5]string
	array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
	// 只有数组长度和数组类型相同的数组才能互相赋值
	array1 = array2
	fmt.Printf("array1: %v\n", array1)
}

// 指针数组赋值
func pointerArrayCopy() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	*array2[1] = "Yellow"
	*array2[2] = "Blue"
	array1 = array2
	fmt.Printf("array1: %v\n", array1)
}

// 多维数组
func multArray() {
	var array [4][2]int
	array = [4][2]int{{20, 30}, {21, 22}, {31, 32}, {30, 31}}
	fmt.Printf("array: %v\n", array)
	array = [4][2]int{1: {20, 30}, 2: {30, 40}}
	fmt.Printf("array: %v\n", array)
	array = [4][2]int{1: {0: 30}, 2: {1: 20}}
	fmt.Printf("array: %v\n", array)

	// 同类型多维数组赋值
	var array1 [2][2]int
	var array2 [2][2]int
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	array1 = array2
	fmt.Printf("array1: %v\n", array1)

	// 使用索引为多维数组赋值
	var array3 [2]int = array1[1]
	fmt.Printf("array3: %v\n", array3)
	var value int = array1[1][0]
	fmt.Printf("value: %v\n", value)
}

// func foo(array *[1e6]int){
// 	...
// }

// // 在函数里传递数组
// // 传入数组的指针会节省资源，但是如果改变指针指向的值，会改变共享的内存
// var array [1e6]int
// foo(&array)

func main() {
	f1()
	f2()
	copy()
	pointerArrayCopy()
	multArray()
}
