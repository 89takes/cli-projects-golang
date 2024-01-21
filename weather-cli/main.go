package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const weatherAPIURL = "https://api.openweathermap.org/data/2.5/weather"

var apiKey string
var city string

func init() {
	flag.StringVar(&apiKey, "api-key", "", "Your OpenWeatherMap API key")
	flag.StringVar(&city, "city", "", "Name of the city")
	flag.Parse()

	if apiKey == "" {
		apiKey = "81bead084b5dfe3432009ae4f0f16753"
		fmt.Println("Warning: Using default api-key, Use -api-key flag to enter your own key.")
	}
}

func main() {
	weather, err := getWeather(city)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Weather in %s:\n", city)
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp-273.15)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
}

func getWeather(city string) (*WeatherData, error) {
	// Example coordinates for New York City
	latitude := 40.7128
	longitude := -74.0060
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", weatherAPIURL, latitude, longitude, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with error code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil

}


type WeatherData struct {
	Main    MainInfo   `json:"main"`
	Weather []Weather  `json:"weather"`
}

type MainInfo struct {
	Temp float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}