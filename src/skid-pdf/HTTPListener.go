package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/schema"
)

// http://localhost:8080/html?grayscale=false&landscape=true&uri=developers.mindbodyonline.com
func pdfHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		pdfr := pdfRequest{}

		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&pdfr)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		decoder := schema.NewDecoder()
		// r.PostForm is a map of our POST form values
		err = decoder.Decode(&pdfr, r.PostForm)

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
