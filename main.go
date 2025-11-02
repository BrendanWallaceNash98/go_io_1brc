package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/BrendanWallaceNash98/golang-1brc-io/logger"
	"github.com/BrendanWallaceNash98/golang-1brc-io/models"
)

func processFile(file *os.File, chunkStart int64, chunkEnd int64, waitgroup *sync.WaitGroup) {
	scanner := bufio.NewScanner(file)
	wso := models.IntialiseWeatherStation()

	for scanner.Scan() {
		line := scanner.Text()

		semiColonPos := strings.Index(line, ";")
		if semiColonPos == -1 {
			continue
		}
		stationName := line[:semiColonPos]
		tempStr := line[semiColonPos+1 : len(line)-1]
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			logger.LogError(err)
		}
		wso.AddWeatherStation(stationName, temp)

	}
}

func main() {
	filename := `/Users/brendanwallace-nash/OneBillionRowChallenge/data/measurements.txt`
	file, err := os.Open(filename)
	logger.PanicError(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.LogError(err)
		}
	}(file)

	numWorker := runtime.NumCPU()

	fileInfo, _ := os.Stat(filename)
	fileSize := fileInfo.Size()
	chunkSize := fileSize / int64(numWorker)

	results := make(chan map[string]*models.WeatherStation, numWorker)
	var wg sync.WaitGroup

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		start := int64(i) * chunkSize
		end := start + chunkSize
		if i == numWorker-1 {
			end = fileSize
		}

		go processFile(file, start, end, &wg)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	outPutFile, err := os.Create("/Users/brendanwallace-nash/OneBillionRowChallenge/data/output.txt")
	logger.LogError(err)
	defer func() {
		if err := outPutFile.Close(); err != nil {
			logger.PanicError(err)
		}
	}()

	sort.Strings(wso.WeatherStationsName)

	for _, weatherStation := range wso.WeatherStationsName {
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
