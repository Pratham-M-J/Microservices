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
	"github.com/gorilla/mux"
)

const port = ":8000"

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags) //logger
	sm := mux.NewRouter()                                 //router

	ph := handler.NewProducts(l)

	getRouter := sm.Methods(http.MethodGet).Subrouter() //subrouter for GET requests
	getRouter.HandleFunc("/products", ph.GetProducts)   //handle GET requests to /products
	//sm.Handle("/products", ph)
	// sm.Handle("/products/", ph)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation) //middleware for validating product data in PUT requests
	slog.Info("server started", "port", port)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation) //middleware for validating product data in POST requests

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

	sig := <-done // block until signal received

	l.Println("Received terminate, graceful shutdown", sig)
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
