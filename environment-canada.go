package main

import (
	"encoding/xml"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html/charset"
)

type DateTime []struct {
	TimeZone     string     `xml:"zone,attr"`
	UTCOffset    int        `xml:"UTCOffset,attr"`
	Year         int        `xml:"year"`
	Month        time.Month `xml:"month"`
	Day          int        `xml:"day"`
	Hour         int        `xml:"hour"`
	Minute       int        `xml:"minute"`
	TextSummarty string     `xml:"textSummary"`
}

func (dt DateTime) Time() time.Time {
	d := dt[1]
	l, err := time.LoadLocation(d.TimeZone)
	if err != nil {
		panic(err)
	}
	return time.Date(d.Year, d.Month, d.Day, d.Hour, d.Minute, 0, 0, l)
}

type Location struct {
	Continent string `xml:"continent"`
	Country   string `xml:"country"`
	// CountryCode string `xml:"country > code,attr"`
	Province string `xml:"province"`
	// ProvinceCode string `xml:"province > code,attr"`
	City string `xml:"name"`
	// CityCode     string `xml:"city > code,attr"`
	// Longitude    string `xml:"city > lon,attr"`
	// Latitude     string `xml:"city > lat,attr"`
	Region string `xml:"region"`
}

type Wind struct {
	Speed     float64 `xml:"speed"`
	Gust      float64 `xml:"gust"`
	Direction string  `xml:"direction"`
	Bearing   float64 `xml:"Bearing"`
}

type Conditions struct {
	ObservationTime  DateTime `xml:"dateTime"`
	Condition        string   `xml:"condition"`
	IconCode         int      `xml:"iconCode"`
	Temperature      float64  `xml:"temperature"`
	Dewpoint         float64  `xml:"dewpoint"`
	WindChill        float64  `xml:"windChill"`
	Pressure         float64  `xml:"pressure"`
	Visibility       float64  `xml:"visibility"`
	RelativeHumidity float64  `xml:"relativeHumidity"`
	Wind             Wind     `xml:"wind"`
}

type AbbreviatedForecast struct {
	IconCode int     `xml:"iconCode"`
	POP      float64 `xml:"pop"`
	Summary  string  `xml:"textSummary"`
}

type Forecast struct {
	Period              string              `xml:"Period"`
	Summarty            string              `xml:"textSummary"`
	AbbreviatedForecast AbbreviatedForecast `xml:"abbreviatedForecast"`
}

type ForecastGroup struct {
	IssueTime DateTime   `xml:"dateTime"`
	Forcast   []Forecast `xml:"forecast"`
}

type Weather struct {
	DateCreated       DateTime      `xml:"dateTime"`
	Location          Location      `xml:"location"`
	CurrentConditions Conditions    `xml:"currentConditions"`
	ForecastGroup     ForecastGroup `xml:"forecastGroup"`
}

/*
0: sun

*/

func Load() (*Weather, error) {
	// http://dd.weather.gc.ca/citypage_weather/xml/ON/s0000571_e.xml
	responce, err := http.Get("http://dd.weather.gc.ca/citypage_weather/xml/ON/s0000571_e.xml")
	if err != nil {
		return nil, err
	}
	w := &Weather{}
	decoder := xml.NewDecoder(responce.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(w)
	if err != nil {
		return nil, err
	}
	responce.Body.Close()
	// fmt.Printf("%#v", w.CurrentConditions.Wind)
	spew.Dump(w.ForecastGroup.Forcast)
	os.Exit(1)
	return w, nil
}
