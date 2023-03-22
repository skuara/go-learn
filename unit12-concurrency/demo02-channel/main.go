package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  go 提供了一种被称为通道的机制，用于在goroutine之间共享数据
//  通道分为无缓冲通道和缓冲通道，无缓冲通道用于执行goroutine之间同步通信，而缓冲通道用于执行异步通信
//  无缓冲通道保证在接受和发送产生的瞬间执行两个goroutine之间的交换
//  通道由make函数构建  该函数制定chan关键字和通道的元素类型
//  Unbuffered := make(chan int)
//  buffered := make(chan int, 10)

//  通道的发送和接收特性
//  1. 对于同一个通道，发送操作之间是互斥的，接受操作之间也是互斥的
//  2. 发送操作和接收操作对元素值的处理都是不可分割的
//  3. 发送操作在完全完成之前会被阻塞，接受操作也是如此

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	values <- value // 通过通道发送
}

func main() {
	defer close(values) // 关闭通道， defer作用 延迟关闭，最后关闭
	go send()
	fmt.Println("waiting...")
	value := <-values // 通过通道接收
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end")
}
