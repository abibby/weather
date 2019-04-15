package weather

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

type Unit struct {
	Type  string `xml:"unitType,attr"`
	Units string `xml:"units,attr"`
	Class string `xml:"class,attr"`
	Value string `xml:",innerxml"`
}

func (u *Unit) Float64() float64 {
	f, _ := strconv.ParseFloat(u.Value, 64)
	return f
}

func (u *Unit) String() string {
	switch u.Units {
	case "C":
		return fmt.Sprintf("%0.0f°C", u.Float64())
	default:
		return fmt.Sprintf("%0.1f %s", u.Float64(), u.Units)
	}
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
	Temperature      Unit     `xml:"temperature"`
	Dewpoint         Unit     `xml:"dewpoint"`
	WindChill        Unit     `xml:"windChill"`
	Pressure         Unit     `xml:"pressure"`
	Visibility       Unit     `xml:"visibility"`
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
	// fmt.Printf("%+v", w)
	// spew.Dump(w.ForecastGroup.Forcast)
	// os.Exit(1)
	return w, nil
}
