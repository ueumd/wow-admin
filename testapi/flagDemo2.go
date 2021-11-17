package main

import (
	"flag"
	"fmt"
)

func main() {
	namePtr := flag.String("name", "username", "姓名")
	agePtr := flag.Int("age", 18, "年龄")
	musclePtr := flag.Bool("muscle", true, "是否有肌肉")

	var email string
	flag.StringVar(&email, "email", "chenqionghe@sina.com", "邮箱")


	flag.Parse()

	args := flag.Args()
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("muscle:", *musclePtr)
	fmt.Println("email:", email)
	fmt.Println("args:", args)
}