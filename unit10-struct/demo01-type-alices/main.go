package main

import "fmt"

// 类型定义   类型别名
// 类型别名只会在代码中存在，编译完成后就不存在了
// 类型别名可以延续使用原类型的方法
// type NewType Type

func type_alices() {
	type MyInt2 = int
	var i MyInt2
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i)
}

func main() {
	type MyInt int
	// i 为MyInt类型
	var i MyInt
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i)
	type_alices()

}
