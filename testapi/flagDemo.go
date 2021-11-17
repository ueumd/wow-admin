package main

import (
"flag"
"fmt"
)

func main() {
	// 参数 默认值 说明
	// 返回指向对应类型指针并分配对应指针指向的对象
	pi := flag.Int("a", 10, "apple")
	flag.Parse()
	fmt.Printf("%v\n", *pi)
	/**
	运行：
	$ go run flagDemo.go -a 20
	20

	// 默认值
	$ go run flagDemo.go
	10
	*/
}

// 等价写法
func main2() {
	// 指定指针指向的对象
	var i int
	flag.IntVar(&i, "a", 10, "apple")
	flag.Parse()
	fmt.Printf("%v\n", i)
}