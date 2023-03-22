package main

import "fmt"

var c = make(chan int)

// 无缓冲通道

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()
	// 协程会阻塞
	for i := 0; i < 3; i++ {
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			fmt.Println("channel is close")
		}
	}

	// for v := range c {
	// 	fmt.Printf("v: %v\n", v)
	// }

	// for i := 0; i < 3; i++ {
	// 	r := <-c
	// 	fmt.Printf("r: %v\n", r)
	// }

	// r := <-c
	// fmt.Printf("r: %v\n", r)
	// r = <-c
	// fmt.Printf("r: %v\n", r)
	// r = <-c
	// fmt.Printf("r: %v\n", r)

}
