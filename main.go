package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	listenAddress = flag.String("web.listen", ":9119", "Address on which to expose metrics and web interface.")
	hosts         = flag.String("hosts", "127.0.0.1:9999", "comma separated list of hosts to crawl")
)

func main() {
	flag.Parse()
	hosts := strings.Split(*hosts, ",")

	prometheus.MustRegister(NewExporter(hosts))
	log.Printf("Starting Server: %s", *listenAddress)
	handler := prometheus.Handler()
	http.Handle("/metrics", handler)

	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
