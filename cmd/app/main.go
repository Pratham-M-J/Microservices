package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
)

const port = ":8080"

func main() {

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("GET /school", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello I'm in school")

		d, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error reading body", "error", err)
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		log.Printf("Data: %s", d)

		w.Write(d)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Pratham is cool")
		w.Write([]byte("Pratham is cool"))
	})

	slog.Info("server started", "port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
