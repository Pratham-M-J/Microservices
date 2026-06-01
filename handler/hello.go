package handler

import (
	"io"
	"log"
	"log/slog"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello I'm in school")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading body", "error", err)
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data: %s", d)

	w.Write([]byte("Schoolll"))
}
