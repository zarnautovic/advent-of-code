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

	sum := 0;

	powerSum := 0;


	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)


	for fileScanner.Scan() {
		oneGame := fileScanner.Text()

		before, after, found := strings.Cut(oneGame, ":")

		if (!found) {
			continue
		}

		gameNumberString:= before[strings.LastIndex(before, " ")+1:]
		gameNumber, err := strconv.Atoi(gameNumberString)

		if err != nil {
			fmt.Println(err)
		}

		cubes := strings.FieldsFunc(after, func(r rune) bool { return strings.ContainsRune(",;", r) })

		flag := true;

		red:=0;
		green :=0;
		blue:=0;

		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			numberString, color, found := strings.Cut(cube, " ");
			if(!found) {
				break
			}

			number, err := strconv.Atoi(numberString);

			if err != nil {
				fmt.Println(err)
			}

			if (color == "red") {

				if (number > 12) {
					flag = false
				}

				if (red < number) {
					red = number
				}
			}

			if (color == "green") {

				if (number > 13) {
					flag = false
				}

				if (green < number) {
					green = number
				}
			}

			if (color == "blue") {
				if (number > 14) {
					flag = false
				}

				if ( blue < number) {
					blue = number
				}
			}
		}

		if (flag) {
			sum += gameNumber
		}

		powerSum += red * green * blue;
	}
	fmt.Println(sum)
	fmt.Println(powerSum)

	readFile.Close()
}