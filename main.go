package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"wow-admin/api"
	"wow-admin/dao"
	"wow-admin/utils"
)

type mainAppBoot struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	wait       *utils.WaitGroup
}

func (mainApp *mainAppBoot) Init() {
	mainApp.ctx, mainApp.cancelFunc = context.WithCancel(context.Background())
	mainApp.wait = &utils.WaitGroup{}

	utils.InitViper()
	utils.InitLogger()
	utils.InitRedis()

	//if err := utils.InitSqlxDB(); err != nil {
	//	fmt.Printf("init DB failed, err:%v\n", err)
	//	return
	//}

	//fmt.Println("init DB succeeded")

	dao.DB = utils.InitMyQLDB()

	api.Init(mainApp.ctx, mainApp.wait)
}

func (mainApp *mainAppBoot) Stop() {
	mainApp.cancelFunc()
	mainApp.wait.Wait()
	utils.Logger.Info("client-config-server app exit....")
}

func main() {
	mainApp := &mainAppBoot{}
	mainApp.Init()

	signalChan := make(chan os.Signal, 1)

	// SIGHUP   终止进程(终端线路挂断) 终端控制进程结束(终端连接断开)
	// SIGTERM  终止进程(软件终止信号) 结束程序(可以被捕获、阻塞或忽略)
	// SIGINT   终止进程(中断进程)     用户发送INTR字符(Ctrl+C)触发
	// SIGKILL  终止进程(杀死进程)     无条件结束程序(不能被捕获、阻塞或忽略)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGHUP:
		case syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM:
			mainApp.Stop()
			os.Exit(0)
		}
	}
}
