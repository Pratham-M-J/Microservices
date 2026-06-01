package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pratham-M-J/microservices/handler"
)

const port = ":8080"

func main() {

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		w.Write([]byte("Hello world"))
	})

	l := log.New(os.Stdout, "product-api", log.LstdFlags) //logger
	sm := http.NewServeMux()                              //router

	hh := handler.NewHello(l) //handlers
	gh := handler.NewGoodbye(l)
	oh := handler.NewOmg(l)

	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/omg", oh)

	slog.Info("server started", "port", port)

	s := &http.Server{ //it's a struct in Go - https://pkg.go.dev/net/http#Server
		Addr:         port,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	done := make(chan os.Signal, 1) // channel waits for shutdown signal

	signal.Notify(
		done,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done // block until signal received

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	err := s.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown the server")
		log.Fatal(err)
	}
}
