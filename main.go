package main

import (
	"fmt"
)

func main() {
	ip, location := FindLocation()
	fmt.Printf("your ip is %s\n", ip)
	fmt.Printf("your location is %s\n", location)
	weather := getWeatherCity(location)
	fmt.Println("weather:", weather)
	placesToVisit := recommendPlacesToVisit(weather)
	gearToCarry := recommendGearToCarry(weather)
	fmt.Println("Places to visit:", placesToVisit)
	fmt.Println("Gear to carry:", gearToCarry)
}

func recommendPlacesToVisit(weather string) string {
	return "Central Park"
}

func recommendGearToCarry(weather string) string {
	return "Coat"
}
