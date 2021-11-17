package main
import (
"log"
"os"
"os/signal"
"syscall"
)
func main() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	// 监听所有信号
	log.Println("listen sig")
	signal.Notify(sigs)
	// 打印进程id
	log.Println("PID:", os.Getppid())

	s := <-sigs
	log.Println("退出信号", s)
}