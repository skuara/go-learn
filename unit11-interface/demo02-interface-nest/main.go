package main

import "fmt"

type TranSporter interface {
	whistle(int) int
}

type steamer interface {
	from() string
	to() string
	TranSporter // 接口嵌套
}

type train struct {
}

func (t train) from() string {
	return "beijing"
}
func (t train) to() string {
	return "hangzhou"
}
func (t train) whistle(a int) int {
	return a
}

func main() {

	var t train
	fmt.Printf("t.from(): %v\n", t.from())
	fmt.Printf("t.to(): %v\n", t.to())
	fmt.Printf("t.whistle(2): %v\n", t.whistle(2))

}
