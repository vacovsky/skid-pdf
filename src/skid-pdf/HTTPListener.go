package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func gofPDFHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))
	w.Header().Set("Content-Type", "application/pdf")

	gofPDFFromURL(pdfURL, w)
	// w.Write(result)
}

func pdfHandle(w http.ResponseWriter, r *http.Request) {

	useGrayscal, err := strconv.ParseBool(r.Form.Get("grayscale"))
	useLandscape, err := strconv.ParseBool(r.Form.Get("landscape"))

	if err != nil {
		print("Unable to parse all params from query string")
	}

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))

	extraParams := []string{}
	if useGrayscal {
		extraParams = append(extraParams, WkGrayscale)
	}
	if useLandscape {
		extraParams = append(extraParams, WkOrientationLandscape)
	}

	result := GenerateWKPDF(pdfURL)
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
