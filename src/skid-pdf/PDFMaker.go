package main

import (
	"fmt"

	"github.com/terryh/gopdf"
)

func getBytesFromURL(pdfURL string) []byte {
	return []byte{}
}

func makePDFFromURL(pdfURL string) []byte {
	result, err := gopdf.Url2pdf(pdfURL)
	fmt.Println(err)
	return result
}

func makePDFFromImage(pdfURL string) []byte {
	result, err := gopdf.Url2jpeg(pdfURL)
	fmt.Println(err)
	return result
}
