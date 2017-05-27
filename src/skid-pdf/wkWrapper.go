package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const wkhtmltopdfCmd = "wkhtmltopdf"

// WkOrientationLandscape - if passed, sets orientation to landscape
const WkOrientationLandscape = " -O Landscape "

//WkGrayscale - If passed, created PDF will be grayscale
const WkGrayscale = " -g "

// GenerateWKPDF creates PDF from target URL.  First argument is the URL to
// be converted, other arguments should be pulled from the constants above.
// Additionally, arbitrary params can be passed as long as the string is
// buffered by spaces on both sides.  The result will look something
// like this:
// $> wkhtmltopdf http://someurl/something.html - --your-arguments
func GenerateWKPDF(targetURL string, params []string) []byte {
	var result bytes.Buffer
	wkCommand := []string{}

	for _, param := range params {
		wkCommand = append(wkCommand, param)
	}

	wkCommand = append(wkCommand, targetURL)
	wkCommand = append(wkCommand, "-")

	cmd := exec.Command(wkhtmltopdfCmd, wkCommand...)

	// for testing
	fmt.Println(wkhtmltopdfCmd, wkCommand)

	cmd.Stdout = &result
	err := cmd.Run()

	if err != nil {
		log.Print(err)
	}
	return result.Bytes()
}
