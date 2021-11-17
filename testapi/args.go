package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	args := os.Args
	arg1 := os.Args[1]

	// 你可以使用索引的方式来获取单个参数
	arg3 := os.Args[3]

	fmt.Println(args) // [./args a b c d]
	fmt.Println(arg1) // a
	fmt.Println(arg3) // c

	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}

	/**
	[./args a b c d]
	a
	c
	参数0: ./args
	参数1: a
	参数2: b
	参数3: c
	参数4: d
	 */

	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Println(os.Args[1:])

	// 去掉第一个参数
	for idx, args := range os.Args[1:] {
		fmt.Println("参数：" + strconv.Itoa(idx) + ":", args)
	}

	/**
	hsd:testapi daysun$ go build args.go
	hsd:testapi daysun$ ./args a b c d
	[./args a b c d]
	a
	c
	参数0: ./args
	参数1: a
	参数2: b
	参数3: c
	参数4: d
	a
	b
	c
	d
	[a b c d]
	参数：0: a
	参数：1: b
	参数：2: c
	参数：3: d
	 */
}