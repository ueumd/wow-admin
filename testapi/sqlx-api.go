package main

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注册MySQL驱动
	"github.com/jmoiron/sqlx"
	"strings"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:Abcdef@123456@tcp(127.0.0.1:3306)/ueumd_test?charset=utf8mb4&parseTime=True"
	// Connect包含了open和ping方法，也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return db.Ping()
}

type User struct {
	ID			int				`db:"id"`
	Age 		int				`db:"age"`
	Name		string		`db:"name"`
}



//Get和Select是一个非常省时的扩展，可直接将结果赋值给结构体，其内部封装了StructScan进行转化。
//Get用于获取单个结果然后Scan，Select用来获取结果切片。

// 查询单条数据示例 结构体，Get
func getDemo()  {
	sqlStr := "select * from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 2)

	if err != nil {
		fmt.Printf("get failed err:%v\n", err)
		return
	}

	fmt.Println(u.ID, u.Name, u.Age)
}

// select查询
func selectDemo()  {
	sqlStr := "select * from user where id > ?"

	var users []User

	err := db.Select(&users, sqlStr, 0)

	if err != nil {
		fmt.Printf("get failed err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// Exec
func  execInsert()  {
	sqlStr := "insert into user (name, age) values(?, ?)"

	ret, err := db.Exec(sqlStr, "Jerry", 18)

	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	theId, err := ret.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(theId)
}

func execUpdate()  {
	sqlStr := "update user set age=? where id = ?"

	ret, err := db.Exec(sqlStr, 22, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	// 影响的行数
	count, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}

func execDelete()  {
	sqlStr := "delete from user where id=?"

	ret, err := db.Exec(sqlStr, 1)

	if err != nil {
		fmt.Println(err)
	}

	count, err := ret.RowsAffected()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}

// Query使用的connection在所有的rows通过Next()遍历完后或者调用rows.Close()后释放
func queryDemo()  {
	rows, err := db.Query("select name, age from user")
	if err != nil {
		fmt.Println("query failed, error: ", err)
		return
	}
	// //循环结果
	for rows.Next() {
		var name string
		var age int
		err = rows.Scan(&name, &age)
		fmt.Println(name, age)
	}

}
// Queryx和Query行为很相似，不过返回一个sqlx.Rows对象
//支持扩展的scan行为,同时可将对数据进行结构体转换
func queryxDemo()  {
	rows, err := db.Queryx("select id, name, age from user")
	if err != nil {
		fmt.Println("queryx failed, error: ", err)
		return
	}
	defer rows.Close()
	// //循环结果
	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		fmt.Println(u)
	}

}

//QueryRow和QueryRowx
//QueryRow和QueryRowx都是从数据库中获取一条数据，但是QueryRowx提供scan扩展，可直接将结果转换为结构体。
func queryRowAndQueryRowx()  {
	row := db.QueryRow("select id, name, age FROM user where id = ?",3) // QueryRow返回错误，错误通过Scan返回
	var id int
	var name string
	var age int
	err :=row.Scan(&id,&name,&age)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("this is QueryRow res:[%d:%s:%d]\n",id,name,age)

	var u User
	err = db.QueryRowx("select id, name, age FROM user where id = ?",2).StructScan(&u)
	if err != nil{
		fmt.Println("QueryRowx error :",err)
	}else {
		fmt.Printf("this is QueryRowx res:%v",u)
	}
}

// NamedQuery
//NamedQuery方法用来绑定SQL语句与结构体或map中的同名字段, 查询。


// Sqlx.In
func queryByIDs(ids []int) (users []User, err error)  {
	query , args, err := sqlx.In("select * from user where id in (?)", ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func testQueryByIDs()  {
	// 无法自定义顺序，默认id从小到大顺序
	users, err := queryByIDs([]int {7, 3, 6, 2})
	if err != nil {
		fmt.Printf("QueryByIDs failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}

	/**
	user:main.User{ID:2, Age:20, Name:"tom"}
	user:main.User{ID:3, Age:19, Name:"jack"}
	user:main.User{ID:6, Age:18, Name:"Jerry"}
	user:main.User{ID:7, Age:18, Name:"Jerry"}
	*/
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	// FIND_IN_SET维护顺序
	// mysql中find_in_set()函数的使用及in()用法详解 https://www.jb51.net/article/143105.htm
	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}


func testQueryAndOrderByIDs() {
	// 1. 用代码去做排序
	// 2. 让MySQL排序
	fmt.Println("----")
	users, err := QueryAndOrderByIDs([]int{2, 6, 3, 4}) // 维护id查询顺序
	if err != nil {
		fmt.Printf("QueryAndOrderByIDs failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}

	/**
	user:main.User{ID:2, Age:20, Name:"tom"}
	user:main.User{ID:6, Age:18, Name:"Jerry"}
	user:main.User{ID:3, Age:19, Name:"jack"}
	user:main.User{ID:4, Age:18, Name:"Jerry"}

	*/
}

// User 结构体实现driver.Valuer接口：
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring  // INSERT INTO user (name, age) VALUES (?, ?), (?, ?), (?, ?)
	fmt.Println(args)  // 查看生成的args // [西瓜 18 yy 23 kk 24]
	_, err := db.Exec(query, args...)
	return err
}


func testBatchInsertUsers2() {
	u1 := User{Name: "XX", Age: 18}
	u2 := User{Name: "ZZ", Age: 23}
	u3 := User{Name: "JJ", Age: 24}
	users := []interface{}{u1, u2, u3}
	err := BatchInsertUsers2(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	}
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

func testBatchInsertUsers3()  {
	u1 := User{Name: "西瓜", Age: 18}
	u2 := User{Name: "yy", Age: 23}
	u3 := User{Name: "kk", Age: 24}

	users := []*User{&u1, &u2, &u3}
	err := BatchInsertUsers3(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	}

}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("init DB success...")
	//getDemo()
	//selectDemo()
	//execInsert()
	//execUpdate()
	//execDelete()

	// queryDemo()
	// queryxDemo()

	//queryRowAndQueryRowx()

	//testQueryByIDs()
	//
	testQueryAndOrderByIDs()

	// testBatchInsertUsers2()

	// testBatchInsertUsers3()
}