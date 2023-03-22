package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)
	fmt.Printf("%T %dB %v~%v\n", u8, unsafe.Sizeof(u8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", u16, unsafe.Sizeof(u16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", u32, unsafe.Sizeof(u32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", u64, unsafe.Sizeof(u64), 0, uint64(math.MaxUint64)) //uint64(math.MaxUint64)  ???

	// 十进制
	var a int = 10
	fmt.Printf("a: %v\n", a)
	// %b 二进制
	fmt.Printf("%b \n", a)

	// 八进制  0开头
	var b int = 077
	fmt.Printf("%o \n", b)

	// 十六进制 0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	// 浮点型
	fmt.Printf("%f \n", math.Pi)
	// %.2f  保留两位小数
	fmt.Printf("%.2f \n", math.Pi)

}
