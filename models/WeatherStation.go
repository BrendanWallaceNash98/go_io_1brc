package models

import (
	"fmt"
	"math"
)

type WeatherStation struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int
	Avg   float64
}

type WeatherStations struct {
	WeatherStationsMap  map[string]*WeatherStation
	WeatherStationsName []string
}

func IntialiseWeatherStation() WeatherStations {
	wsMap := make(map[string]*WeatherStation)
	var wsSlice []string
	WeatherStationsObj := WeatherStations{
		WeatherStationsMap:  wsMap,
		WeatherStationsName: wsSlice,
	}

	return WeatherStationsObj
}

func (ws *WeatherStations) AddCalculatedWeatherStation(w WeatherStation) error {
	if _, exists := ws.WeatherStationsMap[w.Name]; exists {
		return fmt.Errorf("weatherstation already exists, cannot add to map")
	}
	ws.WeatherStationsMap[w.Name] = &w
	ws.WeatherStationsName = append(ws.WeatherStationsName, w.Name)

	return nil
}

func (ws *WeatherStation) InitialiseWeatherStation(name string, temp float64) {
	ws.Name = name
	ws.Min = temp
	ws.Max = temp
	ws.Sum = temp
	ws.Count = 1
}

func (ws *WeatherStation) AddNewValue(temp float64) {
	ws.Sum = ws.Sum + temp
	ws.Count = ws.Count + 1

	if temp < ws.Min {
		ws.Min = temp
		return
	}

	if temp > ws.Max {
		ws.Max = temp
		return
	}
}

func (wso *WeatherStation) MergeStations(_wso *WeatherStation) error {
	if wso.Name != _wso.Name {
		return fmt.Errorf("not the same weatherstation, should not merge")
	}
	wso.Count += _wso.Count
	wso.Sum += _wso.Sum
	if wso.Min > _wso.Min {
		wso.Min = _wso.Min
	}
	if wso.Max < _wso.Max {
		wso.Max = _wso.Max
	}
	return nil
}

func (ws *WeatherStation) CalAverageTemp() {
	avg := ws.Sum / float64(ws.Count)
	ratio := math.Pow(10, 1.0)
	ws.Avg = math.Round(avg*ratio) / ratio
}

func (wso *WeatherStations) AddWeatherStation(name string, tmp float64) {
	val, exists := wso.WeatherStationsMap[name]
	if exists {
		val.AddNewValue(tmp)
	} else {
		ws := &WeatherStation{}
		ws.InitialiseWeatherStation(name, tmp)
		wso.WeatherStationsMap[name] = ws
		wso.WeatherStationsName = append(wso.WeatherStationsName, name)
	}
}
