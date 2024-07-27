package main

import (
	"fmt"
	"trippit/location"
	"trippit/recommendations"
	"trippit/weather"
)

func main() {
	ip, location, err := location.FindLocation()
	if err != nil {
		fmt.Println("Error finding location:", err)
		return
	}
	fmt.Printf("your ip is %s\n", ip)
	fmt.Printf("your location is %s\n", location)
	weatherInfo := weather.GetWeatherCity(location)
	placesToVisit := recommendations.RecommendPlacesToVisit(location)
	gearToCarry := recommendations.RecommendGearToCarry(weatherInfo)
	fmt.Println("Places to visit:", placesToVisit)
	fmt.Println("Gear to carry:", gearToCarry)
}
