package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func checkPermit(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		return
	}
	if ok == true {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)

	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
func TestCasBin(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	//基本权限设置
	checkPermit(e, "demo", "/user", "read")
	checkPermit(e, "demo", "/order", "write")
	checkPermit(e, "demo1", "/user/userList", "read")
	checkPermit(e, "demo1", "/order/orderList", "write")
}
