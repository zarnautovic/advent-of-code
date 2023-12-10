package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	filePath := os.Args[1]

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var tileMap [][]string

	for fileScanner.Scan() {
		row := fileScanner.Text()
		tileMap = append(tileMap, strings.Split(row, ""))
	}


	intMap := createIntMap(tileMap)
	start := findStartPoint(tileMap)


	getNextStep([]int{start[0], start[1]+1}, tileMap, 0, intMap, "left") 

	getNextStep([]int{start[0], start[1]-1}, tileMap, 0, intMap, "right")

	getNextStep([]int{start[0]-1, start[1]}, tileMap, 0, intMap, "down")

	getNextStep([]int{start[0]+1, start[1]}, tileMap, 1, intMap, "up")

	loop := math.Ceil(float64(getHighestNumber(intMap)) / 2)

	fmt.Println("Part 1", loop)

	grid := getLines("./input.txt")
	startingPoint := findStart(grid)
	visited := map[Point]int{startingPoint:0}
	notChecked := []Point{startingPoint}

	maxDist := 0
	for len(notChecked) > 0 {
		current := notChecked[0]
		notChecked = notChecked[1:]
		next := nextPoints(grid, current)
		for _,point := range next {
			if _,found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDist = max(maxDist, visited[current] + 1)
				notChecked = append(notChecked, point)
			}
		}
	}

	var result = maxDist
	fmt.Println("Day 10 Part 1 Result: ", result)


	countInside := 0
	for y, row := range grid {
		for x := range row {
			if isInside(grid, Point{x, y}, visited) {
				countInside++
			}
		}
	}

	var result2 = countInside
	fmt.Println("Part 2", result2)

}

func findStartPoint(tileMap [][]string) []int {
	index := []int{-1, -1}
	for i, row := range tileMap {
		for j, tile := range row {
			if tile == "S" {
				index[0] = i
        index[1] = j
				return index
			}
		}
	}
	return index
}

func createIntMap(tileMap [][]string) [][]int {
	intMap := [][]int{}
	for _, arr := range tileMap {
		intArr := make([]int, len(arr))
		for j := range arr {
			intArr[j] = 0
		}
		intMap = append(intMap, intArr)
	}
	return intMap
}

func getNextStep(start []int, tileMap [][]string, count int, intMap[][]int, previous string) [][]int {
	startValue := tileMap[start[0]][start[1]]
	count++
	if startValue == "-" {
		if previous == "right" {
			goLeft(tileMap, start, count, intMap)
		}

		if previous == "left" {
			goRight(tileMap, start, count, intMap)
		}
	}

	if startValue == "|" {
		if previous == "up" {
			goDown(tileMap, start, count, intMap)
		}

		if previous == "down" {
			goUp(tileMap, start, count, intMap)
		}
	}

	if startValue == "L" {
		if previous == "right" {
			goUp(tileMap, start, count, intMap)
		}

		if previous == "up" {
			goRight(tileMap, start, count, intMap)
		}
	}

	if startValue == "J" {
		if previous == "up" {
			goLeft(tileMap, start, count, intMap)
		}

		if previous == "left" {
			goUp(tileMap, start, count, intMap)
		}
	}

	if startValue == "F" {
		if previous == "right" {
			goDown(tileMap, start, count, intMap)
		}
		if previous == "down" {
			goRight(tileMap, start, count, intMap)
		}
	}

	if startValue == "7" {
		if previous == "down" {
			goLeft(tileMap, start, count, intMap)
		}
		if previous == "left" {
			goDown(tileMap, start, count, intMap)
		}
	}

	if startValue == "S" {
		return intMap
	}
	return [][]int{}
}

func goLeft(tileMap [][]string, start []int, count int, intMap[][]int) [][]int {
	if start[1] == 0 {
		return intMap
	}
	left := tileMap[start[0]][start[1]-1]
	if left == "L" ||  left == "F" || left == "-" {
		if (intMap[start[0]][start[1]-1]  <= count +1 ) {
			intMap[start[0]][start[1]-1] = count+1
		}
		return getNextStep([]int{start[0], start[1]-1}, tileMap, count, intMap, "right")
	}
	return intMap
}

func goRight(tileMap [][]string, start []int, count int, intMap[][]int) [][]int {
	rowLength := len(tileMap[0])
	if start[1] == rowLength - 1 {
		return intMap
	}
	right := tileMap[start[0]][start[1]+1]
	if right == "J" ||  right == "7" || right == "-" {
		if (intMap[start[0]][start[1]+1]  <= count +1 ) {
			intMap[start[0]][start[1]+1] = count+1
		}
		return getNextStep([]int{start[0], start[1]+1}, tileMap, count, intMap, "left")
	}
	return intMap
}

func goUp(tileMap [][]string, start []int, count int, intMap[][]int) [][]int {
	if start[0] == 0 {
		return intMap
	}
	up := tileMap[start[0]-1][start[1]]
	if up == "7" ||  up == "F" || up == "|" {
		if (intMap[start[0]-1][start[1]]  <= count +1 ) {
			intMap[start[0]-1][start[1]] = count+1
		}
		return getNextStep([]int{start[0]-1, start[1]}, tileMap, count, intMap, "down")
	}
	return intMap
}

func goDown(tileMap [][]string, start []int, count int, intMap[][]int) [][]int {
	colLength := len(tileMap)
	if start[0] == colLength - 1 {
		return intMap
	}
	down := tileMap[start[0]+1][start[1]]
	if down == "J" ||  down == "L" || down == "|" {
		if (intMap[start[0]+1][start[1]]  <= count +1 ) {
			intMap[start[0]+1][start[1]] = count+1
		}
		return getNextStep([]int{start[0]+1, start[1]}, tileMap, count, intMap, "up")
		}
	return intMap
}

func getHighestNumber(intMap [][]int) int {
	highest := 0
	for _, arr := range intMap {
		for _, num := range arr {
			if num > highest {
				highest = num
			}
		}
	}
	return highest
}

func isInside(grid []string, p Point, theLoop map[Point]int) bool {
	if _,part := theLoop[p]; part {
		return false
	}
	count := 0
	cornerCounts := map[byte]int{}
	for y := p.y + 1; y < len(grid); y++ {
		check := Point{p.x, y}
		tile := grid[y][p.x]
		if tile == 'S' {
			tile = findStartTile(Point{p.x, y}, grid)
		}
		if _,part := theLoop[check]; part {
			if (tile == '-') {
				count++
			} else if tile != '|' && tile != '.' {
				cornerCounts[tile]++
			}
		}
	}

	count += max(cornerCounts['L'], cornerCounts['7']) - abs(cornerCounts['L'] - cornerCounts['7'])
	count += max(cornerCounts['F'], cornerCounts['J']) - abs(cornerCounts['F'] - cornerCounts['J'])
	return count % 2 == 1
}

func findStart(grid []string) Point {
	for y, row := range grid {
		for x, col := range row {
			if byte(col) == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func findStartTile(start Point, grid []string) byte {
	points := nextPoints(grid, start)
	minx, maxx, miny, maxy := min(points[0].x, points[1].x), max(points[0].x, points[1].x), min(points[0].y, points[1].y), max(points[0].y, points[1].y)
	if points[0].x == points[1].x {
		return '|'
	} else if points[0].y == points[1].y {
		return '-'
	} else if minx < start.x && miny < start.y {
		return 'J'
	} else if maxx > start.x && maxy > start.y {
		return 'F'
	} else if maxx > start.x && miny < start.y {
		return 'L'
	} else if minx < start.x && maxy > start.y {
		return '7'
	}
	return '.'
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextPoints(grid []string, p Point) []Point {
	points := []Point{}
	switch grid[p.y][p.x] {
	case '|':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case '-':
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case 'L':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case 'J':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case '7':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case 'F':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case '.':
	case 'S':
		down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		if down == '|' || down == 'L' || down == 'J' {
			points = append(points, Point{p.x, p.y + 1})
		}
		if right == '-' || right == '7' || right == 'J' {
			points = append(points, Point{p.x + 1, p.y})
		}
		if up == '|' || up == '7' || up == 'F' {
			points = append(points, Point{p.x, p.y - 1})
		}
		if left == '-' || left == 'L' || left == 'F' {
			points = append(points, Point{p.x - 1, p.y})
		}
	}
	return points
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}