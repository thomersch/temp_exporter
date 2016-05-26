package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type values struct {
	temperature float64
	humidity    float64
}

// sensVals is the crooked implementation of the sensor.
type sensVals struct {
	DTH22 struct {
		Temperature string `json:"Temperature:"`
		Humdity     string `json:"Humidity"`
	} `json:"dth22"`
}

func readSensor(host string, port int) (values, error) {
	var val values
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return val, err
	}
	_, err = conn.Write([]byte("THL0000000000"))
	if err != nil {
		return val, err
	}

	var (
		sens sensVals
		vals values
	)
	dec := json.NewDecoder(conn)
	err = dec.Decode(&sens)
	if err != nil {
		return val, err
	}

	vals.humidity, err = strconv.ParseFloat(sens.DTH22.Humdity, 64)
	if err != nil {
		return vals, err
	}
	vals.temperature, err = strconv.ParseFloat(sens.DTH22.Temperature, 64)
	return vals, err
}
