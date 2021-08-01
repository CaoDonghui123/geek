package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var server *http.Server

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	server = &http.Server{
		Addr:    ":8001",
		Handler: mux,
	}

	group.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(ctx)
	})

	fmt.Println("Starting  httpserver")

	group.Go(func() error {
		return server.ListenAndServe()
	})

	//signal
	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			fmt.Printf("signal:%v\n", sig)
			return server.Shutdown(ctx)
		}
	})

	err := group.Wait()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println("Server exited")

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
