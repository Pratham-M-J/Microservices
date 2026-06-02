package types

import (
	"encoding/json"
	"io"
	"time"
)

// Struct tags define how a struct field maps to JSON keys during encoding and decoding.
// Struct tags tell Go which JSON key corresponds to each struct field.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name`
	Description string  `json:"description`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` //ignores this field from the output
	DeletedOn   string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Flat white coffee",
		Price:       3.50,
		SKU:         "fgd-343-ghh",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Mocha",
		Description: "Chocolate coffee",
		Price:       4.00,
		SKU:         "fgd-343-ghj",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Cappuccino",
		Description: "Coffee with foam",
		Price:       3.00,
		SKU:         "fgd-343-ghk",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          4,
		Name:        "Espresso",
		Description: "espresso",
		Price:       2.50,
		SKU:         "fgd-343-ghl",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// this is done, just to abstract the functionalities and handle the encoding here and not in the handler
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}
