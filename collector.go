package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func readSensor(hostport string) (values, error) {
	var val values

	cmd := []byte(*cmdString)
	if strings.HasPrefix(*cmdString, "http") {
		resp, err := http.Get(*cmdString)
		if err != nil {
			log.Printf("could not fetch remote command: %v", err)
			return val, err
		}
		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return val, err
		}
		cmd = buf
	}

	conn, err := net.DialTimeout("tcp", hostport, time.Second)
	if err != nil {
		return val, err
	}
	_, err = conn.Write(cmd)
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
