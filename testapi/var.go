package main

import (
	"fmt"
	"reflect"
)

const HelloData = 100
const HelloData2 = "world"

// 定义 string 型型 OpType
type OpType string

const (
	Login      OpType = "login"
	BindPhone  OpType = "bindPhone"
	UserName   = 12345
)

const Frame = "Vue"

func main()  {
	//fmt.Println(HelloData)
	//fmt.Println(HelloData2)
	//
	//fmt.Println(Login)
	//fmt.Println(UserName)
	//
	//fmt.Println(string(Login))

	fmt.Println("Frame 的数据类型是:",reflect.TypeOf(Frame))

	// 编译器需要推导出 name 的类型
	const hello = "Hello World"

	// 那么它是如何从无类型的常量 "hello" 中获取类型的呢？
	fmt.Printf("type is %T, value %v", hello, hello)
}