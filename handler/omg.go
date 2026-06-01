package handler

import (
	"log"
	"net/http"
)

type omg struct {
	l *log.Logger
}

func NewOmg(l *log.Logger) *omg {
	return &omg{l}
}

func (o *omg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.l.Print("OMG")
	w.Write([]byte("OMG"))

}
