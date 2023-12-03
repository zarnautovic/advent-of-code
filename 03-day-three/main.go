package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	var matrix [][]string;

	sum1 := 0
	sum2 := 0

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text();

		rowArray := strings.Split(row, "")

		matrix = append(matrix, rowArray)
	}

	numRows := len(matrix)
	numCols := len(matrix[0])


	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
				number, err := strconv.Atoi(matrix[i][j])
				if err != nil {
					number = 10
				}
				if ((number < 0 || number > 9) && matrix[i][j] != ".") {
					completeNumbers := traverseAround(matrix, i, j)

					symbolSum := 0
					for _, completeNumber := range completeNumbers {
						symbolSum += completeNumber
					}
					sum1 += symbolSum

					if (len(completeNumbers) == 2) {
						sum2 += completeNumbers[0] * completeNumbers[1]
					}
				}
		}
	}

	fmt.Println("Part1:", sum1)
	fmt.Println("Part2:", sum2)
}

func traverseAround(matrix[][]string, row int, col int) []int {
	numRows := len(matrix)
	numCols := len(matrix[0])

	var startIndexArray []string

	var completeNumbers []int

	for i := row - 1; i <= row + 1; i++ {
		for j := col - 1; j <= col + 1; j++ {
			if (i >= 0 && i < numRows && j >= 0 && j < numCols) {
				_, err := strconv.Atoi(matrix[i][j])

				if err == nil {
					completeNumberString, startIndex := findCompleteNumber(matrix[i], j)
					position := startIndex+strconv.Itoa(i)

					if(!numberInSlice(position, startIndexArray)) {
						startIndexArray = append(startIndexArray, position)
					
						completeNumber, _ := strconv.Atoi(completeNumberString)
							
						completeNumbers = append(completeNumbers, completeNumber)
					}
				}
			}
		}
	}
	return completeNumbers
}

func findCompleteNumber(row []string, index int) (string, string) {
	var number string
	var startIndex int

	for i := index; i >= 0; i-- {
		if _, err := strconv.Atoi(row[i]); err == nil {
			number = row[i] + number
		} else {
			startIndex = i
			break
		}
	
	}

	for i := index + 1; i < len(row); i++ {
		if _, err := strconv.Atoi(row[i]); err == nil {
			number += row[i]
		} else {
			break
		}
	}

	return number, strconv.Itoa(startIndex)
}

func numberInSlice(position string, slice []string) bool {
	for _, value := range slice {
		if value == position {
			return true
		}
	}

	return false
}