package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

// 等待组进行多个任务的同步，等待组可以保证在并发环境中完成指定数量的任务

func showMsg(i int) {
	defer wp.Done() //Done 方法用于将 WaitGroup 的计数值减一，可以理解为完成一个子任务
	fmt.Printf("i: %v\n", i)
}

func main() {

	for i := 0; i < 10; i++ {
		go showMsg(i)
		wp.Add(1) // Add 方法用于设置 WaitGroup 的计数值，可以理解为子任务的数量
	}

	wp.Wait() // Wait 方法用于阻塞调用者，直到 WaitGroup 的计数值为0，即所有子任务都完成
	fmt.Println("end")

}
