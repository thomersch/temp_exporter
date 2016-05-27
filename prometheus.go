package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	temperature = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "sensor",
		Name:      "temperature",
		Help:      "Current temperature in °C",
	}, []string{"host"})

	humidity = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "sensor",
		Name:      "humidity",
		Help:      "Current humidity in %",
	}, []string{"host"})
)

type Exporter struct {
	// List of host:port
	targets []string
}

func NewExporter(targets []string) *Exporter {
	return &Exporter{
		targets: targets,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	temperature.Describe(ch)
	humidity.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	for _, target := range e.targets {
		vals, err := readSensor(target)
		if err != nil {
			log.Printf("could not read from sensor: %s", err)
			return
		}

		temperature.WithLabelValues(target).Set(vals.temperature)
		humidity.WithLabelValues(target).Set(vals.humidity)
	}
	temperature.Collect(ch)
	humidity.Collect(ch)
}
