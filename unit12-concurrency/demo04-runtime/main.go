package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func showExit() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit() // 退出当前协程
		}
		fmt.Printf("i: %v\n", i)
	}
}

func showA() {
	for i := 0; i < 20; i++ {
		fmt.Printf("A: %v\n", i)
	}
}

func showB() {
	for i := 0; i < 20; i++ {
		fmt.Printf("B: %v\n", i)
	}
}

func main() {
	// go show("java")
	// // 主协程
	// for i := 0; i < 2; i++ {
	// 	runtime.Gosched() // 有权利执行其他任务，让给你来执行（让出cpu时间片，重新等待任务安排）
	// 	fmt.Println("golang")
	// }

	// go showExit()
	// time.Sleep(time.Second)

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2) // 调整运行并发性能  1 单核心执行 >1 多核心执行  <1 不修改任何数值
	go showA()
	go showB()
	time.Sleep(time.Second)
	fmt.Println("end")
}
