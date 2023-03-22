package main

import "fmt"

// make函数创造切片，length是数组的长度，也是切片的初始长度
func f1() {
	var s = make([]int, 2)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", s)
}

// 切片的容量（cap）和长度（len）
func f2() {
	var s = []int{1, 2, 3}
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("cap(s): %v\n", cap(s))
	s1 := make([]int, 2, 4)
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
}

// 切片
func f3() {
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := s1[0:3] //[)
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
}

//数组
func f4() {
	s1 := [...]int{1, 2, 3, 4, 5, 6}
	s2 := s1[2:5]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[0:3]
	fmt.Printf("s4: %v\n", s4)
}

// 切片遍历

func f5() {
	s1 := []int{1, 2, 3, 4, 5}
	l := len(s1)
	for i := 0; i < l; i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	}
	for _, v := range s1 {
		fmt.Printf("v: %v\n", v)
	}
}

func f6() {
	slice := []string{"apple", "orange", "banana", "plum"}
	slice1 := slice[1:2:4] // slice[i:j:k]  设置容量大于已有容量的语言运行时错误
	l := len(slice1)       // l = j - i
	c := cap(slice1)       // c = k - i
	fmt.Printf("l: %v\n", l)
	fmt.Printf("c: %v\n", c)
}

func f7() {
	source := []string{"apple", "orange", "banana", "plum"}
	// 设置容量和长度一样的好处，当我们对slice调用append的时候，容量不够，会创建一个新的底层数组，这个数组包括两个元素 （banana, kiwi)
	slice := source[2:3:3]
	slice = append(slice, "kiwi")
	// 创建新的一个底层数组不会修改原底层数组（source） 的值
	fmt.Printf("source: %v\n", source)
	fmt.Printf("slice: %v\n", slice)
	slice2 := source[2:3:4]
	// 容量足够时，会修改原来底层数组（source）的值
	slice2 = append(slice2, "kiwi")
	fmt.Printf("source: %v\n", source)
	fmt.Printf("slice2: %v\n", slice2)
}

// 多维切片
func multSlice() {
	slice := [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
	fmt.Printf("slice: %v\n", slice)
}

func main() {
	// 可变长度的数组 相同类型可变长度的序列
	// 空切片在未初始化之前默认为nil，长度为0，容量为0
	var names []string
	var numbers []int
	fmt.Printf("names: %T\n", names)
	fmt.Printf("names: %v\n", names)
	fmt.Printf("len(names): %v\n", len(names))
	fmt.Printf("numbers: %T\n", numbers)
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("len(numbers): %v\n", len(numbers))
	f1()
	f2()
	f3()
	f4()
	f5()
	f6()
	f7()
	multSlice()
}
