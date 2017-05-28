package main

import (
	"io"
)

func pdfFromHTMLString(body string, writer io.Writer) {
	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddPage()
	// pdf.SetFont("Arial", "", 8)
	// pdf.Cell(40, 10, "Hello, world")
	// p := pdf.HTMLBasicNew()
	// pdf.AddPage()
	// p.Write(1, body)
	// pdf.Output(writer)
}
