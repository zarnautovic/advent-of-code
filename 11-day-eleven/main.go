package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	filePath := os.Args[1]

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var tileMap []string

	for fileScanner.Scan() {
		row := fileScanner.Text()
		tileMap = append(tileMap, row)
	}

	part1(tileMap, 2)
	part1(tileMap, 1000000)
}



func part1(lines []string, factor int) {
	s := 0

	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}
	for i, l := range lines {
		if strings.Contains(l, "#") {
			continue
		}
		emptyRows[i] = true
	}
	for i := range lines[0] {
		l := ""
		for j := range lines {
			l += string(lines[j][i])
		}
		if strings.Contains(l, "#") {
			continue
		}
		emptyCols[i] = true
	}

	stars := map[int]Coord{}
	n := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				nx := 0
				ny := 0
				for k := 0; k < i; k++ {
					if emptyRows[k] {
						ny++
					}
				}
				for k := 0; k < j; k++ {
					if emptyCols[k] {
						nx++
					}
				}
				stars[n] = Coord{i + ny*(factor-1), j + nx*(factor-1)}
				n++
			}
		}
	}

	for i := range stars {
		for j := i + 1; j < len(stars); j++ {
			s += abs(stars[i].x-stars[j].x) + abs(stars[i].y-stars[j].y)
		}
	}

	fmt.Printf("%d\n", s)
}