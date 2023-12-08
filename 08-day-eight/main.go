package main

import (
	"bufio"
	"fmt"
	"os"
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

	var rows []string
	var lrinstruction []string
	network := map[string][]string{}
	startingPoint := "AAA"
	count := 0
	finishPoint := "ZZZ"
	allCounts := []int{}

	for fileScanner.Scan() {
		row := fileScanner.Text()

		rows = append(rows, row)
	}

	for index, row := range rows {
		if index == 0 {
			lrinstruction = strings.Split(row, "")
		}

		if index > 1 {
			network[row[0:3]] = []string{row[7:10], row[12:15]}
		}
	}

	step := network[startingPoint]

	length := len(lrinstruction)
	for i := 0; i < 100000; i++ {
		index := i
		if i > length-1 {
			index = i % length
		}

		instruction := lrinstruction[index]

		if (instruction == "L" && step[0] == finishPoint) || (instruction == "R" && step[1] == finishPoint) {
			count++
			break
		}
		step = getNextStep(network, step, instruction)
		count++
	}

	steps := findValuesWithLastElementA(network)

	for _, s := range steps {
		counter := 0
		for i := 0; i < 100000; i++ {
			index := i
			if i > length-1 {
				index = i % length
			}
			instruction := lrinstruction[index]

			if (instruction == "L" && s[0][len(s[0])-1] == 'Z') || (instruction == "R" && s[1][len(s[1])-1] == 'Z') {
				counter++
				allCounts = append(allCounts, counter)
				break
			}
			s = getNextStep(network, s, instruction)
			counter++
		}
	}

	fmt.Println("Part 1 result:", count)
	fmt.Println("Part 2 result:", lcmN(allCounts))
}

func getNextStep(network map[string][]string, currentPoint []string, instruction string) []string {
	var nextStep []string
	if instruction == "L" {
		nextStep = network[currentPoint[0]]
	} else {
		nextStep = network[currentPoint[1]]
	}
	return nextStep
}

func findValuesWithLastElementA(network map[string][]string) [][]string {
	valuesWithLastA := [][]string{}
	for key, values := range network {
		if key[len(key)-1] == 'A' {
			valuesWithLastA = append(valuesWithLastA, values)
		}
	}
	return valuesWithLastA
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	// a * b = lcm(a, b) * gcd(a, b)
	return (a * b) / gcd(a, b)
}

func lcmN(n []int) int {
	if len(n) == 2 {
		return lcm(n[0], n[1])
	}
	return lcm(n[0], lcmN(n[1:]))
}
