package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/random"
)

type Product struct {
	Id    int
	Name  string
	Price float32
}

var productList []Product

func GenerateProducts(qtd int) error {
	var products []Product
	for i := 1; i <= qtd; i++ {
		product := Product{
			Id:    i,
			Name:  fmt.Sprintf("Product %d", i),
			Price: float32(random.Float64(100.5, 12977.99)),
		}
		products = append(products, product)
	}
	productList = products
	return nil
}

func getProductsLimit(start int, end int) ([]Product, error) {
	if start < 0 || end > len(productList) || start >= end {
		return nil, fmt.Errorf("invalid range: start=%d, end=%d", start, end)
	}
	return productList[start:end], nil
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	currentPage := r.URL.Query().Get("current_page")
	limit := r.URL.Query().Get("limit")

	if currentPage == "" { currentPage = "1" }
	if limit == "" { limit = "5" }

	GenerateProducts(20)

	index := conversion.StringToInt(currentPage)
	size := conversion.StringToInt(limit)

	start := (index - 1) * size
	end := start + size

	products, err := getProductsLimit(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count := len(productList)
	httputils.WriteJsonPaginated(w, r, http.StatusOK, products, count, index, size)
}

// Start the server
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProductsHandler)

	http.ListenAndServe(":4000", mux)
}
