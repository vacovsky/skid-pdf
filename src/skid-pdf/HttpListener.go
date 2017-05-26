package main

import (
	"net/http"
)

func pdfHandle(w http.ResponseWriter, r *http.Request) {
	result := makePDFFromURL("")

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(result)
}

func jpegHandle(w http.ResponseWriter, r *http.Request) {
	result := makePDFFromURL("")

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(result)
}

func landing(w http.ResponseWriter, r *http.Request) {

}

func help(w http.ResponseWriter, r *http.Request) {

}

func startRouter() {
	http.HandleFunc("/", landing)  // exlains what the service does and how to use it.
	http.HandleFunc("/help", help) // goes to source page
	http.HandleFunc("/html", pdfHandle)
	http.HandleFunc("/jpeg", jpegHandle)

	http.ListenAndServe(":8080", nil)

}
