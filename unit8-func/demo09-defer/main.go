package main

import "fmt"

// defer   延迟处理 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序执行，先被defer的语句最后执行，最后被defer的语句最先被执行

//  defer 用于注册延迟调用
//  这些调用直到return前才被执行，因此可以用来做资源清理
//  多个defer语句按照先进后出执行
//  defer语句中的变量在defer声明时就决定了

//  作用： 1. 关闭文件句柄  2. 锁资源释放  3. 数据库连接释放

// defer后不是跟func 而是跟一个表达语句

// 如果defer后跟一个func，func内部如果发生panic，会把panic暂时搁置，当其他defer执行完后再来执行这个panic

func defer_panic() {
	defer fmt.Println("AAA")
	n := 0
	defer func() {
		fmt.Println(2 / n)
		defer fmt.Println("BBB")
	}()
	defer fmt.Println("CCC")
}

func defer_exe_time() (i int) {
	i = 9
	defer func() {
		fmt.Printf("i=%d\n", i) //打印5 而非9   func里面的所有变量在注册defer时并没有被复制，defer是在函数临近返回之前才会被执行
	}()
	defer fmt.Printf("i=%d\n", i) //变量在注册defer时被拷贝或计算
	return 5
}

func main() {
	// fmt.Println("start")
	// defer fmt.Println("step1")
	// defer fmt.Println("step2")
	// defer fmt.Println("step3")
	// fmt.Println("end")

	// defer_exe_time()
	defer_panic()
}
