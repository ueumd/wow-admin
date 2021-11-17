package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)

	// 1e6 = 1*10^6 = 100000

	//时间戳（秒）：1636616300;
	//时间戳（纳秒）：1636616300046251000;
	//时间戳（毫秒）：1636616300046;
	//时间戳（纳秒转换为秒）：1636616300;
}


func main()  {
	fmt.Println(time.Hour)
	fmt.Println(time.Second)
	fmt.Println(time.Now().Hour())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2021-11-11 15:32:09
	fmt.Println(time.Now().Format(time.UnixDate))     // Thu Nov 11 15:32:09 CST 2021


	// 秒 1636616165
	fmt.Println(time.Now().Unix())

	// 获取指定日期的时间戳
	dt, _ := time.Parse("2006-01-02 15:04:05", "2021-11-11 15:32:09")

	// 1636644729
	fmt.Println(dt.Unix())

	// 1636702329
	fmt.Println(time.Date(2021, 11,12,15,32,9,0, time.Local).Unix())

}
