package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func init()  {
	client = redis.NewClient(&redis.Options{
		Addr: Address,
		Password: Password,
		DB: DB,			// 选择的库
		PoolSize: 20,	// 设置连接数，默认10个连接
	})
}
func main() {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	key := "account"
	field1 := "name"
	fields := map[string]interface{}{
		"addr":"beijing",
		"age":99,
		"skills":"golang",
		"demo1":"aaa",
		"demo2":"bbb",
	}
	//hash 设置一个键的field
	_ = client.HSet(key,field1,"zhangsan")

	// hash 批量设置 ,第二个参数是map类型
	status := client.HMSet(key,fields).Val()
	fmt.Println(status) //ok
	// hash 删除键的field,返回删除field的个数
	_ = client.HDel(key,"demo2").Val()
	//hash 获取field的值
	name := client.HGet(key,"name").Val()
	fmt.Println(name) //zhangsan
	//hash 获取多个field值,返回slice
	values := client.HMGet(key,"name","age").Val()
	fmt.Println(values)//[zhangsan 99]
	//hash 获取所有的值 返回map
	valueAll := client.HGetAll(key).Val()
	fmt.Println(valueAll) //map[addr:beijing age:99 demo1:aaa name:zhangsan skills:golang]
	// hash 获取所有field 返回slice
	fs := client.HKeys(key).Val()
	fmt.Println(fs) //[name addr age skills demo1]
	// hash 获取所有filed的值 返回slice
	vs := client.HVals(key).Val()
	fmt.Println(vs) //[zhangsan beijing 99 golang aaa]
	// 判断一个filed是否存在 返回bool
	e := client.HExists(key,"skills").Val()
	fmt.Println(e) //true
	// hash field自增
	n := client.HIncrBy(key,"age",1).Val()
	fmt.Println(n) //100
}