package main

import "fmt"

// go函数 都是值传递，改变变量的值，可以创建一个指向该变量地址的指针变量，传递数据使用指针，无需拷贝数据
// 类型指针不能进行偏移和运算

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	var sp *string
	var s string = "hello"
	sp = &s
	fmt.Printf("sp: %v\n", sp)
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %v\n", *sp)

	var bp *bool
	var b bool = true
	fmt.Printf("bp: %v\n", bp)
	bp = &b
	fmt.Printf("bp: %v\n", bp)
	fmt.Printf("bp: %T\n", bp)
	fmt.Printf("*bp: %v\n", *bp)

	// 假设i为指针，那&i为“指针地址”，i为“值地址”或“指针”，*i为“值地址指向的值”，三者以指针i为中枢，构成了一个前中后的关系
	// 定义一个指针，才最先有指针地址，然后初始化值地址，才有了值地址，这时候才能对值地址进行赋值
	var np *int
	// new()函数 创建一个int类型的变量，初始值是0 返回值是np指向int类型变量的指针
	np = new(int)
	fmt.Printf("np: %v\n", *np)

}
