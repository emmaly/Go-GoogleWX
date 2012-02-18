package googlewx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Data string `xml:"data,attr"`
}

type wx struct {
	Weather weather `xml:"weather"`
}

type weather struct {
	Forecast_information forecastinformation `xml:"forecast_information"`
	Current_conditions   wxConditions        `xml:"current_conditions"`
	Forecast_conditions  []wxConditions      `xml:"forecast_conditions"`
}

type forecastinformation struct {
	City data `xml:"city"`
	Time data `xml:"current_date_time"`
}

type wxConditions struct {
	Condition      data `xml:"condition"`
	Day_of_week    data `xml:"day_of_week"`
	Temp_f         data `xml:"temp_f"`
	Low            data `xml:"low"`
	High           data `xml:"high"`
	Wind_condition data `xml:"wind_condition"`
	Humidity       data `xml:"humidity"`
}

type Weather struct {
	City     string
	Time     string
	Current  Conditions
	Forecast []Conditions
}

type Conditions struct {
	Day       string
	Condition string
	Temp      string
	Low       string
	High      string
	Wind      string
	Humidity  string
}

func Get(query string) (*Weather, error) {
	wxRoot := new(wx)
	wxRes, err := http.Get(fmt.Sprintf("http://www.google.com/ig/api?weather=%s", query))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(wxRes.Body)
	if err != nil {
		return nil, err
	}
	xml.Unmarshal(body, wxRoot)
	wx := wxRoot.Weather

	weather := new(Weather)

	weather.City = wx.Forecast_information.City.Data
	weather.Time = wx.Forecast_information.Time.Data

	weather.Current.Condition = wx.Current_conditions.Condition.Data
	weather.Current.Temp = wx.Current_conditions.Temp_f.Data
	weather.Current.Wind = wx.Current_conditions.Wind_condition.Data
	weather.Current.Humidity = wx.Current_conditions.Humidity.Data

	weather.Forecast = make([]Conditions, len(wx.Forecast_conditions))
	for i := range wx.Forecast_conditions {
		weather.Forecast[i].Day = wx.Forecast_conditions[i].Day_of_week.Data
		weather.Forecast[i].Low = wx.Forecast_conditions[i].Low.Data
		weather.Forecast[i].High = wx.Forecast_conditions[i].High.Data
		weather.Forecast[i].Condition = wx.Forecast_conditions[i].Condition.Data
	}

	return weather, nil
}
