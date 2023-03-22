package main

import (
	"fmt"
)

func add(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		if a < b {
			max = b
		} else {
			max = a
		}
	}
	return max

}

func main() {
	fmt.Printf("add(2, 3): %v\n", add(2, 3))
	fmt.Printf("compare(2, 3): %v\n", compare(2, 3))
	r := add(4, 5)
	fmt.Printf("r: %v\n", r)

}
