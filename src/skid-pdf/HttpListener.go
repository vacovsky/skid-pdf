package main

import (
	"fmt"
	"net/http"
)

func gofPDFHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))
	w.Header().Set("Content-Type", "application/pdf")

	gofPDFFromURL(pdfURL, w)
	// w.Write(result)
}

func pdfHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))
	result := PDFFromURL(pdfURL)
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(result)
}

func landing(w http.ResponseWriter, r *http.Request) {

}

func help(w http.ResponseWriter, r *http.Request) {

}

func startHTTPListener() {
	http.HandleFunc("/", landing)  // exlains what the service does and how to use it.
	http.HandleFunc("/help", help) // goes to source page
	http.HandleFunc("/html", pdfHandle)
	http.HandleFunc("/gof", gofPDFHandle)

	// http.HandleFunc("/jpeg", jpegHandle)

	http.ListenAndServe(":8080", nil)

}
