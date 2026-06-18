//Package classification of product API
//
//Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Pratham-M-J/microservices/types"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
//
// Returns all products.
//
// responses:
//
//	200: productsResponse
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	w.Header().Set("Content-Type", "application/json")
	lp := types.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	w.Header().Set("Content-Type", "application/json")

	prod := r.Context().Value(KeyProduct{}).(*types.Product) //retrieving the product data from the context, which was set in the middleware

	//validate the product data
	err := prod.Validate()
	if err != nil {
		http.Error(w, "Unable to validate product data", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %v", prod)
	types.AddProduct(prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //mux.Vars() is a function from the Gorilla Mux package that extracts variables from the URL path of an HTTP request. It returns a map of variable names to their corresponding values.
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products")
	prod := r.Context().Value(KeyProduct{}).(*types.Product) //retrieving the product data from the context, which was set in the middleware

	w.Header().Set("Content-Type", "application/json")
	//validate the product data
	err = prod.Validate()
	if err != nil {
		http.Error(w, "Unable to validate product data", http.StatusBadRequest)
		return
	}
	types.UpdateProduct(id, prod)
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &types.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			KeyProduct{},
			prod,
		)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
