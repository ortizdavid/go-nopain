package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/docgen"
)


func main() {
	http.HandleFunc("/customers-excel", greet)
	http.ListenAndServe(":8080", nil)
}


type customer struct {
	Name	string
	Gender	string
	Age		int
}


func getAllCustomers() []customer {
	return []customer{
		{Name: "John Doe", Gender: "Male", Age: 30},
		{Name: "Jane Smith", Gender: "Female", Age: 25},
		{Name: "Emily Davis", Gender: "Female", Age: 22},
		{Name: "Michael Brown", Gender: "Male", Age: 40},
		{Name: "Jessica Wilson", Gender: "Female", Age: 35},
		{Name: "David Johnson", Gender: "Male", Age: 28},
		{Name: "Laura Martinez", Gender: "Female", Age: 32},
		{Name: "James Rodriguez", Gender: "Male", Age: 27},
		{Name: "Maria Hernandez", Gender: "Female", Age: 29},
		{Name: "Robert Clark", Gender: "Male", Age: 45},
	}
}


func greet(w http.ResponseWriter, r *http.Request) {
	excelGenerator := docgen.NewExcelGenerator()
	excelGenerator.AddTitle("Customers")

	excelGenerator.AddHeaderRow(
		"Name",
		"Gender", 
		"Age", 
	)
	
	for _, cu := range getAllCustomers() {
		excelGenerator.AddDataRow(
			cu.Name,
			cu.Gender,
			cu.Age,
		)
	}
	
	err := excelGenerator.SaveToFile(w, "customers.xlsx")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

