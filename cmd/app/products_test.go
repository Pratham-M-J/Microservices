package main

import (
	"testing"

	"github.com/Pratham-M-J/microservices/types"
)

func TestMain(t *testing.T) {
	p := &types.Product{
		Name:        "Brysta",
		Description: "Brown creamy coffee",
		SKU:         "fgg-fees-uhd",
		Price:       3.50,
	}

	err := p.Validate()
	if err != nil {
		t.Error("Expected error, got nil")
	}

}
