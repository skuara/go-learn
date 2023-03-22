package main

import "fmt"

// 指向数组的指针
const MAX int = 3

func main() {
	a := []int{1, 3, 5}
	var ptr [MAX]*int
	fmt.Println(ptr)
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("a[%d]=%d\n", i, *ptr[i])
	}
}
