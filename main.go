package main

import (
	"fmt"
	"github.com/BrendanWallaceNash98/golang-1brc-io/logger"
	"github.com/BrendanWallaceNash98/golang-1brc-io/models"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(`/Users/brendanwallace-nash/OneBillionRowChallenge/data/measurements.txt`)
	logger.PanicError(err)
	var wso models.WeatherStations
	for d := range data {
		dStr := string(d)
		fmt.Println(dStr)
		dList := strings.Split(dStr, `;`)
		num, err := strconv.ParseFloat(dList[1], 64)
		logger.PanicError(err)
		wso.AddWeatherStation(dList[0], num)
	}
}
