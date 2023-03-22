package main

import (
	"fmt"
	"sync"
	"time"
)

var wt sync.WaitGroup

// 多线程取值相同，有时候同时进行加减操作会导致最后结果混乱

var lock sync.Mutex

// 给共享资源加锁

var i = 100

func add() {
	defer wt.Done()
	lock.Lock() // 加锁
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock() // 解锁
}

func sub() {
	defer wt.Done()
	lock.Lock()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}

	wt.Wait()

	fmt.Printf("end i: %v\n", i)
}
