package main

import (
	"reflect"
	"testing"
)

func Test_makePDFFromURL(t *testing.T) {
	type args struct {
		pdfURL string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makePDFFromURL(tt.args.pdfURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makePDFFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makePDFFromImage(t *testing.T) {
	type args struct {
		pdfURL string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makePDFFromImage(tt.args.pdfURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makePDFFromImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
