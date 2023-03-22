package main

import "fmt"

// 接口是一组行为规范的集合

// 定义
type TranSporter interface { //定义接口，通常结尾名以er结尾
	// 接口里面只定义方法，不定义变量
	move(src, dest string) (int, error) // 方法名（参数列表和返回值列表）
	whistle(int) int                    // 参数列表和返回值列表的变量名可以省略
}

type person interface {
	say() string
}

// 实现
type Car struct { //定义结构体时不需要显式的声明他要实现什么接口
	price int
}

// 只要结构体拥有接口里的所有方法，就说明该结构体实现了接口， 一个结构体可以实现多个接口
func (car Car) move(src, dest string) (int, error) {
	car.price = 3
	return car.price, nil
}

// 方法的接受者是值
func (car Car) whistle(a int) int {
	return a
}

// 方法接受者用指针，则实现接口的类型是指针类型
func (car *Car) say() string {
	car.price = 4
	return "hello"
}

// 使用
func transport_1(src, dest string, transporter TranSporter) error {
	_, err := transporter.move(src, dest)
	return err
}

func main() {
	// 接口值由两部分组成，第一个是指向该接口的具体类型的指针， 另外一个指向该具体类型的数据
	car := Car{2}
	var transport TranSporter
	// 接口的赋值   值赋值= 值+指针  都实现了方法
	transport = car
	fmt.Printf("transport.whistle(3): %v\n", transport.whistle(3))
	fmt.Println(transport.move("北京", "天津"))
	fmt.Printf("transport: %v\n", transport) // 2  外面的接口的变量的值没有改变

	fmt.Printf("transport_1(\"北京\", \"天津\", car): %v\n", transport_1("北京", "天津", car))

	var p person
	// 接口的赋值
	p = &Car{2} // car 的指针实现了person接口，所以赋值要赋指针。  对应的值没有实现接口
	p.say()
	fmt.Printf("p: %v\n", p) // 4  外面的接口的变量的值改变了
	fmt.Printf("p.say(): %v\n", p.say())

}
