package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/BrendanWallaceNash98/golang-1brc-io/logger"
	"github.com/BrendanWallaceNash98/golang-1brc-io/models"
)

func ProccessLine(line []byte) (string, float64) {
	lineStr := string(line)
	semiColonPos := strings.Index(lineStr, ";")
	if semiColonPos == -1 {
		return ``, 0
	}
	stationName := lineStr[:semiColonPos]
	temp, err := strconv.ParseFloat(lineStr[semiColonPos+1:len(lineStr)-1], 64)
	if err != nil {
		logger.LogError(err)
		return ``, 0
	}
	return stationName, temp
}

func processFile(file *os.File, chunkStart int64, chunkEnd int64, results chan models.WeatherStations, wg *sync.WaitGroup) {
	wso := models.IntialiseWeatherStation()
	defer wg.Done()
	if chunkStart > 0 {
		file.Seek(chunkStart-1, 0)
		reader := bufio.NewReader(file)
		if _, err := reader.ReadBytes('\n'); err != nil {
			logger.LogError(err)
			return
		}
		chunkStart, _ = file.Seek(0, io.SeekCurrent)
	} else {
		file.Seek(0, 1)
	}

	scanner := bufio.NewScanner(file)
	currentPos := chunkStart

	for scanner.Scan() {
		line := scanner.Text()

		currentPos += int64(len(line)) + 1
		if currentPos > chunkEnd {
			break
		}

		semiColonPos := strings.Index(line, ";")
		if semiColonPos == -1 {
			continue
		}
		stationName := line[:semiColonPos]
		tempStr := line[semiColonPos+1:]
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			logger.LogError(err)
		}
		wso.AddWeatherStation(stationName, temp)

	}
	results <- wso
}

func main() {
	filename := `/Users/brendanwallace-nash/OneBillionRowChallenge/data/measurements.txt`
	numWorker := runtime.NumCPU()

	fileInfo, _ := os.Stat(filename)
	fileSize := fileInfo.Size()
	chunkSize := fileSize / int64(numWorker)

	results := make(chan models.WeatherStations, numWorker)
	wg := new(sync.WaitGroup)

	for i := range numWorker {
		wg.Add(1)
		start := int64(i) * chunkSize
		end := start + chunkSize
		if i == numWorker-1 {
			end = fileSize
		}

		go func(start, end int64) {
			file, err := os.Open(filename)
			logger.LogError(err)
			defer file.Close()
			processFile(file, start, end, results, wg)
		}(start, end)

	}

	go func(wg *sync.WaitGroup, results chan models.WeatherStations) {
		wg.Wait()
		close(results)
	}(wg, results)

	finalWeatherStations := models.IntialiseWeatherStation()
	for weatherStations := range results {
		for _, name := range weatherStations.WeatherStationsName {
			if fws, exists := finalWeatherStations.WeatherStationsMap[name]; exists {
				fws.MergeStations(weatherStations.WeatherStationsMap[name])
			} else {
				finalWeatherStations.AddCalculatedWeatherStation(*weatherStations.WeatherStationsMap[name])
			}
		}
	}

	outPutFile, err := os.Create("/Users/brendanwallace-nash/OneBillionRowChallenge/data/output.txt")
	logger.LogError(err)
	defer func() {
		if err := outPutFile.Close(); err != nil {
			logger.PanicError(err)
		}
	}()

	sort.Strings(finalWeatherStations.WeatherStationsName)

	for _, weatherStation := range finalWeatherStations.WeatherStationsName {
		station := finalWeatherStations.WeatherStationsMap[weatherStation]
		station.CalAverageTemp()
		if _, err := outPutFile.WriteString(fmt.Sprintf("%s=%v/%v/%v\n", station.Name, station.Min, station.Avg, station.Max)); err != nil {
			logger.LogError(err)
		}
	}

	if err := outPutFile.Sync(); err != nil {
		logger.PanicError(err)
	}
}
