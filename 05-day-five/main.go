package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	emptyRowCounter := 0
	var seedsString string
	var seedSoilMap []string
	var soilFertilizerMap []string
	var fertilizerWaterMap []string
	var waterLightMap []string
	var lightTemperatureMap []string
	var temperatureHumidityMap []string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		if row == "" {
			emptyRowCounter++
			continue
		}
		if emptyRowCounter == 0 {
			seedsString += row
		}
		if emptyRowCounter == 1 {
			seedSoilMap = append(seedSoilMap, row)
		}
		if emptyRowCounter == 2 {
			soilFertilizerMap = append(soilFertilizerMap, row)
		}
		if emptyRowCounter == 3 {
			fertilizerWaterMap = append(fertilizerWaterMap, row)
		}
		if emptyRowCounter == 4 {
			waterLightMap = append(waterLightMap, row)
		}
		if emptyRowCounter == 5 {
			lightTemperatureMap = append(lightTemperatureMap, row)
		}
		if emptyRowCounter == 6 {
			temperatureHumidityMap = append(temperatureHumidityMap, row)
		}
	}

	fmt.Println(seedsString)
	fmt.Println(lightTemperatureMap)
}
