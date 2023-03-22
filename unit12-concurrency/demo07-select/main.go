package main

import (
	"fmt"
	"time"
)

// select  （switch）
// Go中的一个控制结构，用于处理异步IO操作，select会监听case语句中的channel的读写操作，当case语句中的channel非阻塞状态（可以读写）时，将会触发相应的动作

// select中的case必须是一个channel操作    default语句总是可运行的

// 多个case都可以执行  随机选择一个执行
// 没有case可以执行 有default  执行default
// 没有case也没有default   select将阻塞，直到某个case通信可以运行

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {

	go func() {
		chanInt <- 2
		chanStr <- "hello"
		close(chanInt)
		close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Println("default....")
		}
		time.Sleep(time.Second)
	}

}
