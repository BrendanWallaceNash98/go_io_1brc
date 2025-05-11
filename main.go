package main 

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func check(e error) {
  if e !- nil {
    panic(e)
  }
}

type WeatherStation struct {
  Name string
  Min float64
  Max float64
  Sum float64
  Count int
}

func (ws *WeatherStation) InialiseWeatherStation(name string, temp float64) {
  ws.Name = name
  ws.Min = temp
  ws.Max = temp
  ws.Sum = temp
  ws.Count = 1
}

func (ws *WeatherStation) AddNewValue(temp float64) {
  ws.sum = ws.sum + temp 
  ws.Count = ws.Count + 1

  if temp < ws.Min {
    ws.Min = temp
    return
  }

  if temp > ws.Max {
    ws.Max  = temp
    return
  }
}

func main() {

}
