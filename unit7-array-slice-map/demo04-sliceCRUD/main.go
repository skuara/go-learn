package main

import "fmt"

// add
func add() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

//delete
func delete() {
	var s1 = []int{1, 2, 3, 4, 5}
	// i 2
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

// update
func update() {
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

// query
func query() {
	var s1 = []int{1, 2, 3, 4}
	key := 2
	for i, v := range s1 {
		if key == v {
			fmt.Printf("i: %v\n", i)
		}
	}
}

// copy
func test5() {
	var s1 = []int{1, 2, 3, 4}
	// var s2 = s1
	//s2 := []int{} 两个切片长度不同报错
	s2 := make([]int, 4)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func main() {
	add()
	delete()
	update()
	query()
	test5()
}
