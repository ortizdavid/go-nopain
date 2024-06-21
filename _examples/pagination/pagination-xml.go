package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
)

type Product struct {
	Id    int
	Name  string
	Price float32
}

var productList []Product

func GenerateProducts(qtd int) error {
	rand.Seed(time.Now().UnixNano())
	var products []Product
	for i := 1; i <= qtd; i++ {
		product := Product{
			Id:    i,
			Name:  fmt.Sprintf("Product %d", i),
			Price: rand.Float32() * 100,
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

func listProductHandler(w http.ResponseWriter, r *http.Request) {
	currentPage := r.URL.Query().Get("current_page")
	limit := r.URL.Query().Get("limit")

	if currentPage == "" { currentPage = "1" }
	if limit == "" { limit = "5" }

	GenerateProducts(20)

	index := conversion.StringToInt(currentPage)
	size := conversion.StringToInt(limit)

	start := (index - 1) * size
	end := start + size
	count := len(productList)

	products, err := getProductsLimit(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.WriteXmlPaginated(w, r, http.StatusOK, products, count, index, size)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", listProductHandler)
	http.ListenAndServe(":4000", mux)
}
