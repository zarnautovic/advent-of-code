package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath);

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var readLines []string

	for fileScanner.Scan() {
		readLines = append(readLines, fileScanner.Text())
	}

	readFile.Close()

	re := regexp.MustCompile(`[^0-9]+`)

	result1 := part1(readLines, re);
	result2 := part2(readLines, re);

	fmt.Println("Result of part 1:", result1)
	fmt.Println("Result of part 2:", result2)
}

func part1(input []string, re *regexp.Regexp ) int {
	sum := 0
	for _, line := range input {
		sum += getFirstAndLastNumber(line, re)
	}
	return sum
}

func part2(input []string, re *regexp.Regexp) int {
	sum := 0
	for _, line := range input {
		replaced := applyReplacementRules(line)
		sum += getFirstAndLastNumber(replaced, re)
	}
	return sum
}

func getFirstAndLastNumber(text string, re *regexp.Regexp) int {
		spliced := strings.Join(strings.Split(text, ""), " ")
		numberOnly := re.ReplaceAllString(spliced, "")
		firstAndLast := string(numberOnly[0]) + string(numberOnly[len(numberOnly)-1])
		number, err := strconv.Atoi(firstAndLast)

		if (err != nil) {
			fmt.Println(err)
		}

		return number;
	}
	
func applyReplacementRules(text string) string {
	replacementRules := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for key, value := range replacementRules {
		text = strings.Replace(text, key, value, -1)
	}

	return text
} 