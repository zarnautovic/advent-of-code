package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	emptyRowCounter := 0
	var seeds []int
	var seedSoilMap [][]int
	var soilFertilizerMap [][]int
	var fertilizerWaterMap [][]int
	var waterLightMap [][]int
	var lightTemperatureMap [][]int
	var temperatureHumidityMap [][]int
	var humidityLocationMap [][]int

	var soils []int
	var fertilizers []int
	var waters []int
	var lights []int
	var temperatures []int
	var humidities []int
	var locations []int

	var seedWithRange [][]int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		if row == "" {
			emptyRowCounter++
			continue
		}
		if emptyRowCounter == 0 {
			seeds = stringToNumbersArray(row)
		}
		if emptyRowCounter == 1 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			seedSoilMap = append(seedSoilMap, array)
		}
		if emptyRowCounter == 2 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			soilFertilizerMap = append(soilFertilizerMap, array)
		}
		if emptyRowCounter == 3 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			fertilizerWaterMap = append(fertilizerWaterMap, array)
		}
		if emptyRowCounter == 4 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			waterLightMap = append(waterLightMap, array)
		}
		if emptyRowCounter == 5 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			lightTemperatureMap = append(lightTemperatureMap, array)
		}
		if emptyRowCounter == 6 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			temperatureHumidityMap = append(temperatureHumidityMap, array)
		}
		if emptyRowCounter == 7 {
			array := stringToNumbersArray(row)
			if len(array) == 0 {
				continue
			}
			humidityLocationMap = append(humidityLocationMap, array)
		}
	}

	soils = findInRange(seeds, seedSoilMap)
	fertilizers = findInRange(soils, soilFertilizerMap)
	waters = findInRange(fertilizers, fertilizerWaterMap)
	lights = findInRange(waters, waterLightMap)
	temperatures = findInRange(lights, lightTemperatureMap)
	humidities = findInRange(temperatures, temperatureHumidityMap)
	locations = findInRange(humidities, humidityLocationMap)

	smallestLocation := getSmallestNumberFromArray(locations)

	fmt.Println(smallestLocation)

	for i := 0; i < len(seeds)-1; i = i + 2 {
		seedWithRange = append(seedWithRange, []int{seeds[i], seeds[i] + seeds[i+1]})
	}

	i := 0
	for i < 100000000 {
		humidity := findInRangeReverse(i, humidityLocationMap)
		temp := findInRangeReverse(humidity, temperatureHumidityMap)
		light := findInRangeReverse(temp, lightTemperatureMap)
		water := findInRangeReverse(light, waterLightMap)
		fertilizer := findInRangeReverse(water, fertilizerWaterMap)
		soil := findInRangeReverse(fertilizer, soilFertilizerMap)
		seed := findInRangeReverse(soil, seedSoilMap)
		for _, item := range seedWithRange {
			if seed >= item[0] && seed < item[1] {
				fmt.Println(i)
				return
			}
		}
		i++
	}
}

func stringToNumbersArray(row string) []int {
	pattern := `\d+`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(row, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func findInRange(array []int, mapArray [][]int) []int {
	var result []int
	for _, item := range array {
		for i := 0; i < len(mapArray); i++ {
			source := mapArray[i][1]
			r := mapArray[i][2]

			if item >= source && item < source+r {
				diff := item - source
				result = append(result, mapArray[i][0]+diff)
			}
		}
	}
	return result
}

func findInRangeReverse(number int, mapArray [][]int) int {
	var result int
	for i := 0; i < len(mapArray); i++ {
		source := mapArray[i][0]
		r := mapArray[i][2]

		if number >= source && number < source+r {
			diff := number - source
			result = mapArray[i][1] + diff
		}
	}
	return result
}

func getSmallestNumberFromArray(array []int) int {
	smallest := array[0]
	for _, item := range array {
		if item < smallest {
			smallest = item
		}
	}
	return smallest
}
