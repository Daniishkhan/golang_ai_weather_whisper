package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

func GetWeatherCity(city string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var openWeatherMapAPI = "http://api.openweathermap.org/data/2.5/weather?units=metric&appid=" + os.Getenv("OPEN_WEATHER_API_KEY")

	resp, err := http.Get(openWeatherMapAPI + "&q=" + city)
	if err != nil {
		return "Error getting weather"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return fmt.Sprintf("Error parsing weather data: %v", err)
	}

	if len(weather.Weather) == 0 {
		return "No weather data available"
	}

	return fmt.Sprintf("It's %.1f°C in %s. Conditions: %s",
		weather.Main.Temp,
		city,
		weather.Weather[0].Main)
}
