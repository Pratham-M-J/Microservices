package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Pratham-M-J/microservices/handler"
)

const port = ":8080"

func main() {

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		w.Write([]byte("Hello world"))
	})

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	sm := http.NewServeMux()

	hh := handler.NewHello(l)
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
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
