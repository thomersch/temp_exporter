# Temperature Exporter

This application collects data from a Snaptekk STTTH222 v03 sensor via TCP and makes it accessible in Prometheus.

## Building

For Go 1.6 and above: `go build`

## Usage

	./temp_exporter -hosts "10.0.0.1:9999,10.0.0.2:9999"

