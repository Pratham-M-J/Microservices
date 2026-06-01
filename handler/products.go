package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pratham-M-J/microservices/types"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle PUT Products")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(types.ProductList)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}
