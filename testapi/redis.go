package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	Address  = "127.0.0.1:6379"
	Password = "Abcdef@123456"
	DB 		 = 3
)

var client *redis.Client

func init()  {
	client = redis.NewClient(&redis.Options{
		Addr: Address,
		Password: Password,
		DB: DB,			// 选择的库
		PoolSize: 20,	// 设置连接数，默认10个连接
	})
}

func main()  {
	// 获取所有健
	//global()

	// 设置
	// redisSet()

	// hash类型操作
	//hash()

	list()

	set()
}

func global()  {
	keys := client.Keys("*").Val()
	fmt.Println(keys)

	size := client.DbSize().Val()
	fmt.Println(size)

	exist := client.Exists("age", "name").Val()

	fmt.Println(exist)

	// 删除键,删除成功返回删除的数,删除失败返回0
	del := client.Del("unknownKey").Val()
	fmt.Println(del)

	// 查看键的有效时间
	ttl := client.TTL("age").Val()
	fmt.Println(ttl)

	// 给键设置有效时间,设置成功返回true,失败返回false
	expire := client.Expire("age",time.Second*86400).Val()
	fmt.Println(expire)

	// 查看键的类型(string,hash,list,set,zset...)
	Rtype := client.Type("store:finish:bill:list").Val()
	fmt.Println(Rtype)

	// 给键重命令,成功返回true,失败false
	Rname := client.RenameNX("age","newAge").Val()
	fmt.Println(Rname)

	// 从redis中随机返回一个键
	key := client.RandomKey().Val()
	fmt.Println(key)

}

func redisSet() {
	// 设置键值对, 10秒后消失
	set1 := client.Set("age", 18, time.Second*10).Val()
	fmt.Println(set1)  // OK

	// key不存在 则设置成功
	set2 := client.SetNX("age", 18, time.Hour * 1).Val()
	fmt.Println(set2) // false

	// key存在则设置成功
	set3 := client.SetXX("age", 19, time.Hour*1).Val()
	fmt.Println(set3) // true

	// 批量设置
	set4 := client.MSet("age1", 17, "name", "Tom").Val()
	fmt.Println(set4) //OK

	// 获取一个键的值 不存在返回空，存在返回值
	res1 := client.Get("age111").Val()
	fmt.Println("res:", res1)

	// 批量获取,获取成功返回slice类型的结果数据
	get2 := client.MGet("age", "age1", "age2").Val()
	fmt.Println(get2) // [30 40 50]

	// 对指定的键进行自增操作
	incr1 := client.Incr("age").Val()
	fmt.Println(incr1) // 31

	// 对指定键进行自减操作
	decr1 := client.Decr("age1").Val()
	fmt.Println(decr1) //39

	// 自增指定的值
	incr2 := client.IncrBy("age", 10).Val()
	fmt.Println(incr2) // 41

	// 自减指定的值
	decr2 := client.DecrBy("age1", 10).Val()
	fmt.Println(decr2) // 29

	// 在key后面追加指定的值,返回字符串的长度
	append1 := client.Append("age2", "abcd").Val()
	fmt.Println(append1) // 6

	// 获取一个键的值得长度
	strlen1 := client.StrLen("age2").Val()
	fmt.Println(strlen1) //6

	// 设置一个键的值,并返回原有的值
	getset1 := client.GetSet("age2", "hello golang").Val()
	fmt.Println(getset1) // 50abcd

	// 设置键的值,在指定的位置
	_ = client.SetRange("age2", 0, "H")
	fmt.Println(client.Get("age2").Val()) // Hello golang

	// 截取一个键的值的部分,返回截取的部分
	newStr := client.GetRange("age2", 6, 11).Val()
	fmt.Println(newStr) //golang
}

func hash()  {
	key := "account"
	field1 := "name"
	fields := map[string]interface{}{
		"addr":"Shanghai",
		"age": 19,
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

func list()  {
	key := "demo"

	client.Del(key)
	for i := 0; i < 5; i++ {
		client.LPush(key, "e-"+strconv.Itoa(i))
	}

	// 获取list 长度
	length := client.LLen(key).Val()
	fmt.Println(length) //5

	// 获取指定下标元素
	value1 := client.LIndex(key, 0).Val()
	fmt.Println(value1) //e-4

	// 获取所有元素
	vs := client.LRange(key, 0, -1).Val()
	fmt.Println(vs) //[e-4 e-3 e-2 e-1 e-0]

	// 修改指定下标的元素值
	status := client.LSet(key, 0, "golang").Val()
	fmt.Println(status) //ok

	// 从右边插入元素
	client.RPush(key, "e-right")
	// 从左边插入元素

	client.LPush(key, "e-left")
	// 从list最右边弹出元素

	v1 := client.RPop(key).Val()
	fmt.Println(v1) // e-right

	// 从list最左边弹出元素
	v2 := client.LPop(key).Val()
	fmt.Println(v2) // e-left

	// 删除指定元素
	n := client.LRem(key, 0, "e-3").Val()
	fmt.Println(n) //1

	status1 := client.LTrim(key, 0, 1).Val()
	fmt.Println(status1) //ok

}

func set()  {
	key := "sdemo"
	key1 := "sdemo1"
	client.Del(key)
	client.Del(key1)
	for i := 0; i < 6; i++ {
		// set 类型添加元素
		client.SAdd(key, "ele-"+strconv.Itoa(i))
	}
	for i := 3; i < 9; i++ {
		// set 类型添加元素
		client.SAdd(key1, "ele-"+strconv.Itoa(i))
	}
}