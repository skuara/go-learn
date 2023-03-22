package main

import (
	"errors"
	"fmt"
	"time"
)

// panic
// when happen ?   1. 运行错误时会导致Panic，比如数组越界，除0   2. 程序主动调用panic(error)
// what will do when panic happen? 1. 逆序执行当前goroutine的defer链（recover从这里介入） 2. 打印错误信息和调用堆栈 3. 调用exit(2)结束整个进程

// recover函数

func soo(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("soo函数中发生了panic:%s\n", err)
		}
	}()
	panic(errors.New("my error"))
}

// error  自定义error
type PathError struct {
	path       string
	op         string
	createtime string
	message    string
}

func (err PathError) Error() string {
	return err.createtime + " " + err.op + " " + err.path + " " + err.message
}

// 异常机制  error

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	var e1 PathError = PathError{
		path:       "/root/test",
		op:         "delete",
		createtime: time.Now().Format("2006-01-02"),
		message:    "目录不存在",
	}
	fmt.Printf("e1.Error(): %v\n", e1.Error())

	if c, err := divide(3, 0); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(c)
	}

	soo(1, 2)
	fmt.Printf("end\n")
}
