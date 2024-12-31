package handlers

import (
	"a5_ClimateDataAnalysis/model"
	"strings"
)

func FindCityWithHighestTemp(data []model.CityData) model.CityData {
	highest := data[0]
	for _, city := range data {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest

}

func FindCityWithLowestTemp(data []model.CityData) model.CityData {
	lowest := data[0]
	for _, city := range data {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

func CalculateAverageRainfall(data []model.CityData) float64 {
	totalRainfall := 0.0
	for _, city := range data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(data))
}

func FilterCitiesByRainfall(data []model.CityData, threshold float64) []model.CityData {
	var filteredCities []model.CityData
	for _, city := range data {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func SearchCityByName(data []model.CityData, name string) (model.CityData, bool) {
	for _, city := range data {
		if strings.EqualFold(city.CityName, name) {
			return city, true
		}
	}
	return model.CityData{}, false
}
