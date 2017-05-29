package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// http://localhost:8080/html?grayscale=false&landscape=true&uri=developers.mindbodyonline.com
func pdfHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// advanced - builds URL using POST data to more easily transmit query
		//string params.  We build those into the wkhtmltopdf http endpoint.
		var err error
		pdfr := pdfRequest{}

		pdfr.Action = r.FormValue("action")
		pdfr.Data = r.FormValue("data")
		pdfr.URL = r.FormValue("url")

		gs := r.FormValue("grayscale")
		pdfr.Grayscale, err = strconv.ParseBool(gs)
		if err != nil {
			fmt.Println("Unable to parse grayscale from query string")
			fmt.Println(err)
		}
		landscapeForm := r.FormValue("landscape")
		pdfr.Landscape, err = strconv.ParseBool(landscapeForm)
		if err != nil {
			fmt.Println("Unable to parse landscape from query string")
			fmt.Println(err)
		}
		result := generateFromPDFRequest(&pdfr)
		w.Header().Set("Content-Type", "application/pdf")
		w.Write(result)

	case "GET":
		// simple
		r.ParseForm()
		grayscaleForm := r.Form.Get("grayscale")
		useGrayscal, err := strconv.ParseBool(grayscaleForm)
		if err != nil {
			fmt.Println("Unable to parse grayscale from query string")
			fmt.Println(err)
		}

		landscapeForm := r.Form.Get("landscape")
		useLandscape, err := strconv.ParseBool(landscapeForm)
		if err != nil {
			fmt.Println("Unable to parse landscape from query string")
			fmt.Println(err)
		}

		pdfURL := fmt.Sprintf("%s", r.Form.Get("uri")) //, r.Form.Get("sid"))

		extraParams := []string{}
		if useGrayscal {
			extraParams = append(extraParams, WkGrayscale...)
		}
		if useLandscape {
			extraParams = append(extraParams, WkOrientationLandscape...)
		}

		result := generateWKPDF(pdfURL, extraParams)
		w.Header().Set("Content-Type", "application/pdf")
		w.Write(result)
	}
}

func landing(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/vacoj/skid-pdf", 301)
}

func help(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/vacoj/skid-pdf", 301)
}

func gofPDFHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pdfURL := fmt.Sprintf("http://%s?sid=%s", r.Form.Get("uri"), r.Form.Get("sid"))
	w.Header().Set("Content-Type", "application/pdf")

	gofPDFFromURL(pdfURL, w)
	// w.Write(result)
}
func startHTTPListener() {

	http.HandleFunc("/", landing)  // exlains what the service does and how to use it.
	http.HandleFunc("/help", help) // goes to source page
	http.HandleFunc("/html", pdfHandle)
	http.HandleFunc("/gof", gofPDFHandle)

	// http.HandleFunc("/jpeg", jpegHandle)

	s := &http.Server{
		Addr:           ":" + settings.HTTPPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
