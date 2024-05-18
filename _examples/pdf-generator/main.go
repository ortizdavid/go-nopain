package main

import (
	"net/http"
	"github.com/ortizdavid/go-nopain/docgen"
	"github.com/ortizdavid/go-nopain/httputils"
)


func main() {
	http.HandleFunc("/simple-pdf", simplePdfHandler)
	http.ListenAndServe(":8080", nil)
}


func simplePdfHandler(w http.ResponseWriter, r *http.Request) {
	var pdfGen docgen.HtmlPdfGenerator

	data := map[string]interface{}{
		"Title": "Simple PDF",
	}
	pdfBytes, err := pdfGen.GeneratePDF("templates/simple-pdf.html", data)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pdfGen.SetOutput(w, pdfBytes, "simple.pdf")
}


func tableReportHandler(w http.ResponseWriter, r *http.Request) {
}