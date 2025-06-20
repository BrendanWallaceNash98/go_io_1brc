package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BrendanWallaceNash98/golang-1brc-io/logger"
	"github.com/BrendanWallaceNash98/golang-1brc-io/models"
)

func main() {
	file, err := os.Open(`/Users/brendanwallace-nash/OneBillionRowChallenge/data/measurements.txt`)
	logger.PanicError(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.LogError(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	wso := models.IntialiseWeatherStation()

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		semiColonPos := strings.Index(line, ";")
		if semiColonPos == -1 {
			continue
		}
		stationName := line[:semiColonPos]
		tempStr := line[semiColonPos+1:]
		temp, err := strconv.ParseFloat(tempStr, 64)
		logger.LogError(err)
		wso.AddWeatherStation(stationName, temp)

		lineCount++
		if lineCount%10000000 == 0 {
			fmt.Printf("Processed %d lines\n\n", lineCount)
		}
	}
	outPutFile, err := os.Create("/Users/brendanwallace-nash/OneBillionRowChallenge/data/output.txt")
	logger.LogError(err)
	defer func() {
		if err := outPutFile.Close(); err != nil {
			logger.PanicError(err)
		}
	}()

	for _, weatherStation := range wso.WeatherStationsNameSorted {
		station := wso.WeatherStationsMap[weatherStation]
		station.CalAverageTemp()
		stationText := fmt.Sprintf("%s=%v/%v/%v\n", station.Name, station.Min, station.Avg, station.Max)
		if _, err := outPutFile.WriteString(stationText); err != nil {
			logger.LogError(err)
		}
	}

	if err := outPutFile.Sync(); err != nil {
		logger.PanicError(err)
	}

	logger.PanicError(scanner.Err())
}
