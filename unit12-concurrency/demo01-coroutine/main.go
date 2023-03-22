package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//  goroutines 是并发运行的函数  golang提供了goroutines作为并发处理操作的一种方式

// func show(msg string) {
// 	for i := 1; i < 10; i++ {
// 		fmt.Printf("msg: %v\n", msg)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

func responseSize(url string) {
	fmt.Println("step1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step2:", url)
	defer response.Body.Close()

	fmt.Println("step3:", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step4:", len(body))
}

func main() {
	// // goroutines  协程
	// go show("go") // 协程1
	// show("test")
	// // go show("test")    // 协程2
	// fmt.Println("end") // 主协程   主函数推出 程序就结束了， 其他协程也会被杀死

	go responseSize("http://baidu.com")
	go responseSize("http://jd.com")
	time.Sleep(10 * time.Second)

}
