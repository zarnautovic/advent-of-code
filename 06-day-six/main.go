package main

import (
	"bufio"
	"fmt"
	"math"
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

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	pattern := `\d+`
	re := regexp.MustCompile(pattern)

	var rows []string
	var times []float64
	var distances []float64
	var part1Result float64 = 1
	var timeString string
	var distanceString string

	for fileScanner.Scan() {
		row := fileScanner.Text()
		rows = append(rows, row)
	}

	fmt.Println(rows)

	for index, row := range rows {

		matches := re.FindAllString(row, -1)

		for _, match := range matches {
			num, err := strconv.ParseFloat(match, 64)
			if err == nil && index%2 == 0 {
				times = append(times, num)
				timeString += match
			} else if err == nil && index%2 != 0 {
				distances = append(distances, num)
				distanceString += match
			}
		}
	}

	for i := 0; i < len(times); i++ {
		insideSquare := math.Pow(times[i], 2) - (4 * 1 * distances[i])
		lowerRange := (times[i] - math.Sqrt(insideSquare)) / (2 * 1)
		upperRange := (times[i] + math.Sqrt(insideSquare)) / (2 * 1)
		numOfSuccessful := math.Floor(upperRange) - math.Floor(lowerRange)

		part1Result *= numOfSuccessful
	}

	time, _ := strconv.ParseFloat(timeString, 64)
	distance, _ := strconv.ParseFloat(distanceString, 64)

	fmt.Println(part1Result)

	insideSquare := math.Pow(time, 2) - (4 * 1 * distance)
	lowerRange := (time - math.Sqrt(insideSquare)) / (2 * 1)
	upperRange := (time + math.Sqrt(insideSquare)) / (2 * 1)
	part2Result := math.Floor(upperRange) - math.Floor(lowerRange)

	fmt.Println(part2Result)

}
