package types

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var ProductList = []*Product{
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
