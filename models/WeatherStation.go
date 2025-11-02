package models

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

func (ws *WeatherStation) CalAverageTemp() {
	ws.Avg = ws.Sum / float64(ws.Count)
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
