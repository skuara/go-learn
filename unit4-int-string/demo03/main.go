package main

import "fmt"

type Website struct {
	Name string
}

func main() {
	// 普通占位符
	var site = Website{Name: "breserker"}
	fmt.Printf("site: %v\n", site)  // v 变量内容
	fmt.Printf("site: %#v\n", site) // # 包含类型输出
	fmt.Printf("site: %T\n", site)  // T 输出类型
	fmt.Println("%%")               // 字面上的百分号，并非值的占位符
	// 布尔占位符
	b := true
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %t\n", b)
	// 数字占位符
	i := 10
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i) // 二进制
	fmt.Printf("i: %d\n", i) // 十进制
	fmt.Printf("i: %o\n", i) // 八进制
	fmt.Printf("i: %c\n", i) // unicode码点所表示的字符
	fmt.Printf("i: %x\n", i) // 十六进制  小写字母表示
	fmt.Printf("i: %X\n", i) // 十六进制  大写字母表示
	fmt.Printf("i: %q\n", i) // 单引号围绕字符字面值，由Go语法安全转义
	fmt.Printf("i: %U\n", i) // Unicode 格式

	// 复数和浮点类型占位符（实部和虚部）
	e := 10.2
	fmt.Printf("e: %e\n", e) // 科学计数法 %e  %E
	fmt.Printf("e: %f\n", e) // 有小数点而无指数
	g := 10.20
	fmt.Printf("g: %g\n", g) // 根据情况选择%e或者%f输出以产生更紧凑的0输出
	G := 10.20 + 2i
	fmt.Printf("G: %G\n", G) // 根据情况选择%G或者%f输出以产生更紧凑的0输出

	// 字符串和字节切片
	x := 100
	p := &x
	fmt.Printf("p: %p\n", p) // 十六进制表示指针地址 0x开头

}
