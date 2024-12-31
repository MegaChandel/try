package main

import (
	"a5_ClimateDataAnalysis/handlers"
	"a5_ClimateDataAnalysis/model"
	"fmt"
)

func main() {
	data := []model.CityData{
		{CityName: "New York", Temperature: 15.2, Rainfall: 1200.5},
		{CityName: "London", Temperature: 11.3, Rainfall: 750.0},
		{CityName: "Mumbai", Temperature: 29.1, Rainfall: 2000.0},
		{CityName: "Tokyo", Temperature: 18.0, Rainfall: 1300.4},
		{CityName: "Sydney", Temperature: 22.5, Rainfall: 850.2},
	}

	highestTempCity := handlers.FindCityWithHighestTemp(data)
	fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highestTempCity.CityName, highestTempCity.Temperature)

	lowestTempCity := handlers.FindCityWithLowestTemp(data)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowestTempCity.CityName, lowestTempCity.Temperature)

	averageRainfall := handlers.CalculateAverageRainfall(data)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	var threshold float64
	fmt.Print("Enter a rainfall threshold to filter cities: ")
	fmt.Scan(&threshold)

	filteredCities := handlers.FilterCitiesByRainfall(data, threshold)
	if len(filteredCities) > 0 {
		fmt.Println("Cities with rainfall above the threshold:")
		for _, city := range filteredCities {
			fmt.Printf("- %s (%.2f mm)\n", city.CityName, city.Rainfall)
		}
	} else {
		fmt.Println("No cities found with rainfall above the threshold.")
	}

	var cityName string
	fmt.Print("Enter a city name to search for: ")
	fmt.Scan(&cityName)

	foundCity, found := handlers.SearchCityByName(data, cityName)
	if found {
		fmt.Printf("City found: %s (%.2f°C, %.2f mm)\n", foundCity.CityName, foundCity.Temperature, foundCity.Rainfall)
	} else {
		fmt.Println("City not found.")
	}

	fmt.Println("Thank you for using the Climate Data Analysis Program!")
}
