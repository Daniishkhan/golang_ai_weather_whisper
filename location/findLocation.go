package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IPInfo struct {
	IP string `json:"ip"`
}

type GeoInfo struct {
	City string `json:"city"`
}

func FindLocation() (string, string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	ip, err := getIP(client)
	if err != nil {
		return "", "", fmt.Errorf("failed to get IP: %w", err)
	}

	city, err := getCity(client, ip)
	if err != nil {
		return "", "", fmt.Errorf("failed to get city: %w", err)
	}

	return ip, city, nil
}

func getIP(client *http.Client) (string, error) {
	resp, err := client.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ipInfo IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "", err
	}

	return ipInfo.IP, nil
}

func getCity(client *http.Client, ip string) (string, error) {
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var geoInfo GeoInfo
	if err := json.NewDecoder(resp.Body).Decode(&geoInfo); err != nil {
		return "", err
	}

	return geoInfo.City, nil
}
