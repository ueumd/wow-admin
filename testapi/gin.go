package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	log.Println(w, "Hello world!")
}
func get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	log.Println(w, "Hello world!")
}

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/", home)
	s.HandleFunc("/get", home)

	server := &http.Server{
		Addr:    ":8090",
		Handler: s,
	}
	go server.ListenAndServe()
	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-sigs:
		log.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		log.Println("http shutdown")
	}
}
