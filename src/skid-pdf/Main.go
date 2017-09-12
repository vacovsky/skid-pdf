package main

import (
	"log"
	"net/http"
	"sync"
	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	appname  = "skid-pdf"
	version  = "1.1.2"
	settings = Settings{}
	promAddr = ":8080"
	httpReqs *prometheus.CounterVec
	pdfTime *prometheus.HistogramVec
)

func main() {
	// load settings from file
	settings.load()
	wg := &sync.WaitGroup{}

	// initialize prometheus metrics for export
	// set up monitoring
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
	
	pdfTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "pdf_seconds",
			Help: "Time taken to create pdf",
		}, 
		[]string{"code"},
	)
	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(pdfTime)
	
	// listen for inbound http-originating requests for PDFs
	go startHTTPListener()
	wg.Add(1)

	// listen on queue for inbound PDF generation messages
	if settings.UseQueue {
		go startQueueListener(wg)
		wg.Add(1)
	}
	wg.Wait()
}
