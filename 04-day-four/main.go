package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	sum := 0

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cards := []string{}

	for fileScanner.Scan() {
		card := fileScanner.Text();

		cards = append(cards, card)
		// copyOfCards := cards
	}

	distinctNumbers := make([]int, len(cards))
	for i := range distinctNumbers {
		distinctNumbers[i] = 1
	}

	for index, card := range cards {
		winingString := substringBetweenChars(card, ':', '|')
		winningNumbers := getNumbersFromString(winingString)

		playedString := substringBetweenChars(card, '|', '\n')
		playedNumbers := getNumbersFromString(playedString)

		count := countOfWinnings(winningNumbers, playedNumbers)

		if (count > 0) {
			result := math.Pow(2, count-1)
			sum += int(result)

			for i := 1; i <= int(count); i++ {
				if(index+i >= len(distinctNumbers)) {
					continue
				}
				distinctNumbers[index+i] = distinctNumbers[index+i] + distinctNumbers[index]
			}
		}
	}

	sum2 := 0

	for _, num := range distinctNumbers {
		sum2 += num
	}
	
	fmt.Println("sum", sum)
	fmt.Println("sum2", sum2)
}

func substringBetweenChars(input string, start byte, end byte) string {
	startIndex := strings.IndexByte(input, start)
	endIndex := strings.IndexByte(input, end)

	if (endIndex == -1) {
		endIndex = len(input)
	
	}
	return input[startIndex+1:endIndex]
}

func getNumbersFromString(input string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(input, -1)
}

func countOfWinnings(winningNumbers []string, playedNumbers []string) float64 {
	var count float64 = 0

	for _, winningNumber := range winningNumbers {
		for _, playedNumber := range playedNumbers {
			if (winningNumber == playedNumber) {
				count++
			}
		}
	}
	return count
}