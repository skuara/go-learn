package main

import (
	"fmt"
	"time"
)

// var wp sync.WaitGroup

// Timer 是 go 中 time 包里的一种一次性计时器，它的作用是定时触发事件，在触发之后这个 Timer 就会失效

func main() {
	// timer1 := time.NewTimer(time.Second * 2)
	// // 创建 timer，时间到了后就会返回创建的 Timer 的指针
	// t1 := time.Now()
	// fmt.Printf("t1: %v\n", t1)
	// t2 := <-timer1.C
	// // 这个 Timer 会在经过 d 这么长的时间时间之后，向 Timer 中的 channel 发送当前时间
	// fmt.Printf("t2: %v\n", t2)

	// timer2 := time.NewTimer(time.Second * 2)
	// <-timer2.C
	// fmt.Println("2s后")
	// time.Sleep(time.Second * 2)
	// fmt.Println("+2s后")
	// <-time.After(time.Second * 2)  // time.After 返回值是chan time
	// fmt.Println("++2s后"
	timer3 := time.NewTimer(time.Second * 2)
	// // wp.Add(1)
	// go func() {
	// 	<-timer3.C
	// 	// defer wp.Done()
	// 	fmt.Println("Timer 3 expired")
	// }()
	// // wp.Wait()

	s := timer3.Stop()
	// 调用stop终止timer。使其失效
	// 调用 Stop() 方法之后，会将这个 Timer 从时间堆里移除，如果这个 Timer 还没超时，依然在时间堆中，那么就会被成功移除并且返回 true；如果这个 Timer 不在时间堆里，说明已经超时了或者已经被 stop 了，这个时候就会返回 false。
	if s {
		fmt.Println("stop...")
	}
	timer3.Reset(time.Second * 2)
	fmt.Println(time.Now())
	<-timer3.C
	t := time.Now()
	fmt.Printf("t: %v\n", t)

	// 问题记录
	// all goroutines are asleep - deadlock

}
