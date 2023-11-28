package main

import (
	"context"
	"github.com/ueumd/logger"
	"os"
	"os/signal"
	"syscall"
	"wow-admin/api"
	"wow-admin/dao"
	"wow-admin/utils"
)

type Program struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	wait       *utils.WaitGroup
}

func (p *Program) Init() {
	p.ctx, p.cancelFunc = context.WithCancel(context.Background())
	p.wait = &utils.WaitGroup{}

	utils.InitViper()
	utils.InitLogger()
	utils.InitRedis()

	//if err := utils.InitSqlxDB(); err != nil {
	//	fmt.Printf("init DB failed, err:%v\n", err)
	//	return
	//}

	//fmt.Println("init DB succeeded")

	dao.DB = utils.InitMyQLDB()

	api.Init(p.ctx, p.wait)
}

func (p *Program) Start() error {
	return nil
}

func (p *Program) Stop() error {
	p.cancelFunc()
	p.wait.Wait()
	logger.Infoln("client-config-server program exit....")
	return nil
}

func main() {
	program := &Program{}
	program.Init()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGHUP:
		case syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM:
			program.Stop()
			os.Exit(0)
		}
	}
}
