package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IPInfo struct {
	IP string `json:"ip"`
}

type GeoInfo struct {
	City string `json:"city"`
}

func FindLocation() (string, string) {
	fmt.Println(" Finding your location")
	// Step 1: Get IP address
	ipURL := "https://api.ipify.org?format=json"
	ipResponse, err := http.Get(ipURL)
	if err != nil {
		log.Fatal(err)
	}
	defer ipResponse.Body.Close()

	ipBody, err := io.ReadAll(ipResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(ipBody, &ipInfo); err != nil {
		log.Fatal(err)
	}

	// Step 2: Get city information using the IP address
	geoURL := fmt.Sprintf("http://ip-api.com/json/%s", ipInfo.IP)
	geoResponse, err := http.Get(geoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer geoResponse.Body.Close()

	geoBody, err := io.ReadAll(geoResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var geoInfo GeoInfo
	if err := json.Unmarshal(geoBody, &geoInfo); err != nil {
		log.Fatal(err)
	}

	return ipInfo.IP, geoInfo.City
}
