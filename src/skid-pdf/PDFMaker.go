package main

import (
	"fmt"

	"github.com/terryh/gopdf"
)

func makePDFFromURL(body string) []byte {
	result, err := gopdf.Url2pdf("http://nvd3.org/examples/stackedArea.html")
	fmt.Println(err)
	return result
}

func makePDFFromImage(body string) []byte {
	result, err := gopdf.Url2jpeg("http://nvd3.org/examples/stackedArea.html")
	fmt.Println(err)
	return result
}
