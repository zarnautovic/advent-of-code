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

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var historyOfInput [][][]int
	var sum int
	var sumPartTwo int

	for fileScanner.Scan() {
		row := fileScanner.Text()

		var rowOfInput [][]int

		arr := strings.Split(row, " ")
		intArr := make([]int, len(arr))

		for i, val := range arr {
				num, _ := strconv.Atoi(val)
				intArr[i] = num
		}
		rowOfInput = append(rowOfInput, intArr)

		historyOfInput = append(historyOfInput, rowOfInput)

	}

	for _, row := range historyOfInput {
		nextRows := calculateNextRows(row)

		nextRows[len(nextRows) - 1] = append(nextRows[len(nextRows) - 1], 0)
		reverseArray := reverse(nextRows)

		nextRows = appendToLastRow(nextRows, "+")
		nextRows2 := appendToLastRow(reverseArray, "-")

		nextHistoryValue := nextRows[0][len(nextRows[0]) - 1]
		nextHistoryValue2 := nextRows2[0][len(nextRows2[0]) - 1]
		sum += nextHistoryValue
		sumPartTwo += nextHistoryValue2
	}
	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", sumPartTwo)
}

func calculateNextRows(row [][]int) [][]int {
	lastIndex := len(row) - 1
	lastRow := row[lastIndex]
	nextRow := getNextRow(lastRow)

	// if next  row sum = 0, return row + nextRow
	newArray := append(row, nextRow)
	if checkIfAllZeroes(nextRow) {
		return newArray;
	}
	return calculateNextRows(newArray)
}

func getNextRow(row []int) []int {
	var nextRow []int
	for i := 0; i < len(row); i++ {
		if i > 0 {
			nextRow = append(nextRow, row[i] - row[i-1])
		}
	}
	return nextRow
}

func checkIfAllZeroes(arr []int) bool {
	for _, num := range arr {
			if num != 0 {
					return false
			}
	}
	return true
}

func reverse(arr [][]int) [][]int {
	reversedArr := make([][]int, len(arr))
	for i, arr := range arr {
    reversedArr[i] = make([]int, len(arr))
    for j := len(arr) - 1; j >= 0; j-- {
        reversedArr[i][len(arr) - 1 - j] = arr[j]
    }
}
	return reversedArr
}

func appendToLastRow(arr [][]int, operation string) [][]int {
	for i := len(arr) - 1; i > 0; i-- {
		lastRow := arr[i]
		secondLastRow := arr[i-1]
		if operation == "+" {
			arr[i-1]  = append(arr[i-1], lastRow[len(lastRow) - 1] + secondLastRow[len(secondLastRow) - 1])
		} else {
			arr[i-1]  = append(arr[i-1], secondLastRow[len(secondLastRow) - 1] - lastRow[len(lastRow) - 1])
		
		}
	}
	return arr
}