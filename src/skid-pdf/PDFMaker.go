package main

import (
	"io"
)

func gofPDFFromURL(pdfURL string, writer io.Writer) {
	b := getBytesFromURL(pdfURL)
	pdfFromHTMLString(string(b), writer)
}

// func PDFFromURL(pdfURL string) []byte {
// 	result, err := gopdf.Url2pdf(pdfURL)
// 	fmt.Println(err)
// 	return result
// }
