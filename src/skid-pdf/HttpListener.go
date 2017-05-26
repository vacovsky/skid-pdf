package main

import (
	"fmt"
	"net/http"
)

func pdfHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))
	result := makePDFFromURL(pdfURL)
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(result)
}

func jpegHandle(w http.ResponseWriter, r *http.Request) {

	result := makePDFFromImage("")

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(result)
}

func landing(w http.ResponseWriter, r *http.Request) {

}

func help(w http.ResponseWriter, r *http.Request) {

}

func startHTTPRouter() {
	http.HandleFunc("/", landing)  // exlains what the service does and how to use it.
	http.HandleFunc("/help", help) // goes to source page
	http.HandleFunc("/html", pdfHandle)
	http.HandleFunc("/jpeg", jpegHandle)

	http.ListenAndServe(":8080", nil)

}
