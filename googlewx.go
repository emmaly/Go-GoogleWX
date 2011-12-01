package googlewx

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Data struct {
	Data string `xml:"attr"`
}

type Wx struct {
	City     Data           `xml:"weather>forecast_information>city"`
	Time     Data           `xml:"weather>forecast_information>current_date_time"`
	Current  WxConditions   `xml:"weather>current_conditions"`
	Forecast []WxConditions `xml:"weather>forecast_conditions"`
}

type WxConditions struct {
	Day_of_week    Data
	Condition      Data
	Temp_f         Data
	Low            Data
	High           Data
	Wind_condition Data
	Humidity       Data
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
	wx := new(Wx)
	weather := new(Weather)
	wxRes, err := http.Get(fmt.Sprintf("http://www.google.com/ig/api?weather=%s", query))
	if err != nil {
		return weather, err
	}
	xml.Unmarshal(wxRes.Body, wx)

	weather.City = wx.City.Data
	weather.Time = wx.Time.Data

	weather.Current.Condition = wx.Current.Condition.Data
	weather.Current.Temp = wx.Current.Temp_f.Data
	weather.Current.Wind = wx.Current.Wind_condition.Data
	weather.Current.Humidity = wx.Current.Humidity.Data

	weather.Forecast = make([]Conditions, len(wx.Forecast))
	for i := range wx.Forecast {
		weather.Forecast[i].Day = wx.Forecast[i].Day_of_week.Data
		weather.Forecast[i].Low = wx.Forecast[i].Low.Data
		weather.Forecast[i].High = wx.Forecast[i].High.Data
		weather.Forecast[i].Condition = wx.Forecast[i].Condition.Data
	}

	return weather, nil
}
