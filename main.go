package main

import (
	"context"
	"fmt"
	"github.com/ueumd/logger"
	"os"
	"os/signal"
	"syscall"
	"wow-admin/api"
	"wow-admin/config"
	"wow-admin/core"
	"wow-admin/global"
	"wow-admin/utils"
)

type Program struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	wait       *utils.WaitGroup
}

func (p *Program) Init() error {
	p.ctx, p.cancelFunc = context.WithCancel(context.Background())
	p.wait = &utils.WaitGroup{}

	var err error

	// config
	if err = config.Init("config.yaml"); err != nil {
		panic(err)
	}
	global.CONFIG = config.Get()

	// logger
	err = logger.Init(global.CONFIG.Log.FilePath, global.CONFIG.Log.IsStdOut, global.CONFIG.Log.LogLevel)
	if err != nil {
		return err
	}

	// API
	address := fmt.Sprintf(":%d", global.CONFIG.Port)
	api.Init(p.ctx, address, p.wait)

	// service
	err = core.InitService()
	if err != nil {
		return err
	}

	// DB
	err = core.InitDB()
	if err != nil {
		return err
	}
	return nil
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
	if err := program.Init(); err != nil {
		panic(err)
	}

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
