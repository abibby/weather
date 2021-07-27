package main

import (
	"fmt"
	"os"

	"github.com/abibby/weather"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	w, err := weather.Load()
	check(err)

	// fmt.Printf("%#v\n", w.ForecastGroup.Forcast[0])
	// os.Exit(1)

	fmt.Printf("%s, %s\n",
		w.ForecastGroup.Forcast[0].AbbreviatedForecast.Summary,
		w.ForecastGroup.Forcast[0].Humidex.String())

	// // http://dd.weather.gc.ca/citypage_weather/xml/ON/s0000571_e.xml
	// feed, err := gofeed.NewParser().ParseURL("https://weather.gc.ca/rss/city/on-5_e.xml")
	// check(err)

	// currentTemp := ""
	// forecasts := []string{}
	// for _, item := range feed.Items {
	// 	if item.Categories[0] == "Current Conditions" {
	// 		currentTemp = strings.Split(item.Title, ": ")[1]
	// 	}

	// 	if item.Categories[0] == "Weather Forecasts" {
	// 		forecasts = append(forecasts, strings.Split(item.Title, ": ")[1])
	// 	}
	// }
	// fmt.Printf("%s Currently %s\n", forecasts[0], currentTemp)
}
