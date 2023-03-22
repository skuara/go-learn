package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 定义多行字符串
	s1 := `
      line1
      line2
      line3
	`
	fmt.Printf("s1: %v\n", s1)

	// 字符串连接
	// 字符串 + 字符串
	name := "tom"
	age := "20"
	msg := name + " " + age
	fmt.Printf("msg: %v\n", msg)
	msg1 := " "
	msg1 += name
	msg1 += " "
	msg1 += age
	fmt.Printf("msg: %v\n", msg1)
	// 字符串格式化输出  拼url的时候使用
	msg2 := fmt.Sprintf("%s, %s", name, age)
	fmt.Printf("msg2: %v\n", msg2)
	// strings.join
	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", s)
	// buffer.Writestring
	var buffer bytes.Buffer
	buffer.WriteString("Tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	// 转义字符
	s2 := "hello \n world"
	fmt.Printf("s2: %v\n", s2)
	// /t tab键  四个空格
	s3 := "hello\tworld"
	// /r 回车符，返回行首
	s4 := " hello\rworld"
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	// 字符串切片操作
	m := "Hello World"
	a := 2
	// %v 字符的整数表示  %c 字符本身
	fmt.Printf("s[a]: %v\n", m[a])
	fmt.Printf("s[a]: %c\n", m[a])
	b := 5
	fmt.Printf("s[a:b]: %v\n", m[a:b]) // [2:5)
	fmt.Printf("s[a:]: %v\n", m[a:])
	fmt.Printf("s[:b]: %v\n", m[:b])

	// 字符串常用函数
	fmt.Printf("len(m): %v\n", len(m))                                               // string length
	fmt.Printf("strings.Split(m, \" \"): %v\n", strings.Split(m, " "))               // 字符串分割 将字符串由某个字符分割成数组形式
	fmt.Printf("strings.Contains(m, \"hello\"): %v\n", strings.Contains(m, "hello")) // 字符串是否包含某个字符或字符串，包含返回true，不包含返回false
	fmt.Printf("strings.ToLower(m): %v\n", strings.ToLower(m))                       // 转换成小写
	fmt.Printf("strings.ToUpper(m): %v\n", strings.ToUpper(m))                       // 转换成大写
	fmt.Printf("strings.HasPrefix(m): %v\n", strings.HasPrefix(m, "Hel"))            // 是否以hel开头
	fmt.Printf("strings.HasSuffix(m, \"ld\"): %v\n", strings.HasSuffix(m, "ld"))     // 是否以ld结尾
	fmt.Printf("strings.Index(m, \"l\"): %v\n", strings.Index(m, "l"))               // 取l字符第一个索引位置
	fmt.Printf("strings.LastIndex(m, \"l\"): %v\n", strings.LastIndex(m, "l"))       // 取l字符最后一个索引位置

}
