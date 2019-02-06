package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	feed, err := gofeed.NewParser().ParseURL("https://weather.gc.ca/rss/city/on-5_e.xml")
	check(err)

	currentTemp := ""
	todaysForecast := ""
	for _, item := range feed.Items {
		if item.Categories[0] == "Current Conditions" {
			currentTemp = strings.Split(item.Title, ": ")[1]
		}

		if item.Categories[0] == "Weather Forecasts" && todaysForecast == "" {
			todaysForecast = strings.Split(item.Title, ": ")[1]
		}
	}
	fmt.Printf("%s Currently %s\n", todaysForecast, currentTemp)
}
