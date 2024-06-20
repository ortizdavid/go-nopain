package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", listProductHandler)
	http.ListenAndServe(":4000", mux)

}

func listProductHandler(w http.ResponseWriter, r *http.Request) {
	pageIndex := r.URL.Query().Get("page_index")
	pageSize := r.URL.Query().Get("page_size")

	if pageIndex == "" {
		pageIndex = "0"
	}

	if pageSize == "" {
		pageSize = "1"
	}

	GenerateProducts()

	index := conversion.StringToInt(pageIndex)
	size := conversion.StringToInt(pageSize)
	products := getProductsLimit(size, index)
	count := len(productList)

	pagination, _ := httputils.NewPagination(r, products, count, index, size)
	httputils.WriteJsonSimple(w, 200, pagination)
}

type Product struct {
	Id int
	Name string
	Price float32
}

var productList []Product

func GenerateProducts() error {
	rand.Seed(time.Now().UnixNano())
	var products []Product
	for i := 1; i <= 30; i++ {
		product := Product{
			Id:    i,
			Name:  fmt.Sprintf("Product %d", i),
			Price: rand.Float32() * 100,
		}
		productList = append(products, product)
	}
	return nil
}

func getProductsLimit(limit int, offset int) []Product {
	// Validate offset
	if offset < 0 {
		offset = 0
	}

	// Apply offset
	if offset >= len(productList) {
		return []Product{}
	}
	limitedProducts := productList[offset:]

	// Apply limit
	if limit > 0 && limit < len(limitedProducts) {
		limitedProducts = limitedProducts[:limit]
	}

	return limitedProducts
}