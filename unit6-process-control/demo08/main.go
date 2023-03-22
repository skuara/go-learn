package main

import "fmt"

// goto跳转到指定标签
func f1() {
	a := 2
	if a == 2 {
		goto LABEL1
	} else {
		fmt.Println("other")
	}
LABEL1:
	fmt.Println("next...")
}

// goto跳出双重循环
func f2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v,%v\n", i, j)
			if i >= 2 && j >= 2 {
				goto END
			}
		}
	}
END:
	fmt.Println("end..")
}

func main() {
	f1()
	f2()
}
