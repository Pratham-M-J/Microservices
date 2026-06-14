package handler

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pratham-M-J/microservices/types"
)

func TestGetproducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rr := httptest.NewRecorder()

	logger := log.New(io.Discard, "", 0)
	handler := NewProducts(logger)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 got %d", rr.Code)
	}
}

func TestAddProduct(t *testing.T) {
	payload := `{
		"Name": "Brysta",
		"Description": "Brown creamy coffee",
		"SKU": "fgg-fees-uhd",
		"Price": 3.50
	}`
	req := httptest.NewRequest(
		http.MethodPost,
		"/products",
		strings.NewReader(payload),
	)

	rr := httptest.NewRecorder()

	logger := log.New(io.Discard, "", 0)
	handler := NewProducts(logger)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 got %d", rr.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	logger := log.New(io.Discard, "", 0)
	handler := NewProducts(logger)

	payload := `{
		"Name": "Testing",
		"Description": "Brown creamy coffee",
		"SKU": "fgg-fees-uhd",
		"Price": 3.50
	}`
	req := httptest.NewRequest(
		http.MethodPut,
		"/products/1",
		strings.NewReader(payload),
	)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 got %d", rr.Code)
	}

	//verify update happened
	products := types.GetProducts()

	if products[0].Name != "Testing" {
		t.Errorf("Expected product name 'Testing' got '%s'", products[0].Name)
	}
}

func TestUpdateProductInvalidID(t *testing.T) {
	logger := log.New(io.Discard, "", 0)

	handler := NewProducts(logger)
	req := httptest.NewRequest(
		http.MethodPut,
		"/products/abc",
		strings.NewReader(`{}`),
	)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 got %d", rr.Code)
	}
}

func TestUpdateProductValidationFailure(t *testing.T) {
	logger := log.New(io.Discard, "", 0)
	handler := NewProducts(logger)

	payload := `{
        "name":"",
        "price":-10
    }`

	req := httptest.NewRequest(
		http.MethodPut,
		"/products/1",
		strings.NewReader(payload),
	)

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %d got %d",
			http.StatusBadRequest,
			rr.Code,
		)
	}
}
