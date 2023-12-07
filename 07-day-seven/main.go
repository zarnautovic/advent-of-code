package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var StrengthMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	totalWinnings := 0
	totalWinningsJokers := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var handsWithBids [][]string
	var fiveOfAKind [][]string
	var fourOfAKind [][]string
	var fullHouse [][]string
	var threeOfAKind [][]string
	var twoPairs [][]string
	var onePair [][]string
	var highCard [][]string

	var fiveOfAKindJokers [][]string
	var fourOfAKindJokers [][]string
	var fullHouseJokers [][]string
	var threeOfAKindJokers [][]string
	var twoPairsJokers [][]string
	var onePairJokers [][]string
	var highCardJokers [][]string

	var handWithBidsSorted [][]string
	var handWithBidsSortedJokers [][]string

	for fileScanner.Scan() {
		row := fileScanner.Text()

		words := strings.Fields(row)
		var temp []string
		for index, word := range words {

			if index%2 == 0 {
				temp = strings.Split(word, "")
			} else {
				temp = append(temp, word)
			}
		}
		handsWithBids = append(handsWithBids, temp)

	}

	for _, hand := range handsWithBids {
		numberOfJokers := countJokers(hand)
		if isFiveOfAKind(hand) {
			fiveOfAKind = append(fiveOfAKind, hand)
			fiveOfAKindJokers = append(fiveOfAKindJokers, hand)
		} else if isForOfAKind(hand) {
			fourOfAKind = append(fourOfAKind, hand)

			if numberOfJokers == 1 || numberOfJokers == 4 {
				fiveOfAKindJokers = append(fiveOfAKindJokers, hand)
			} else {
				fourOfAKindJokers = append(fourOfAKindJokers, hand)
			}

		} else if isFullHouse(hand) {
			fullHouse = append(fullHouse, hand)

			if numberOfJokers >= 1 {
				fiveOfAKindJokers = append(fiveOfAKindJokers, hand)
			} else {
				fullHouseJokers = append(fullHouseJokers, hand)
			}
		} else if isThreeOfAKind(hand) {
			threeOfAKind = append(threeOfAKind, hand)
			if numberOfJokers == 1 || numberOfJokers == 3 {
				fourOfAKindJokers = append(fourOfAKindJokers, hand)
			} else {
				threeOfAKindJokers = append(threeOfAKindJokers, hand)
			}
		} else if isTwoPairs(hand) {
			twoPairs = append(twoPairs, hand)
			if numberOfJokers == 2 {
				fourOfAKindJokers = append(fourOfAKindJokers, hand)
			} else if numberOfJokers == 1 {
				fullHouseJokers = append(fullHouseJokers, hand)
			} else {
				twoPairsJokers = append(twoPairsJokers, hand)
			}
		} else if isOnePair(hand) {
			onePair = append(onePair, hand)
			if numberOfJokers == 1 || numberOfJokers == 2 {
				threeOfAKindJokers = append(threeOfAKindJokers, hand)
			} else {
				onePairJokers = append(onePairJokers, hand)
			}
		} else {
			highCard = append(highCard, hand)
			if numberOfJokers == 1 {
				onePairJokers = append(onePairJokers, hand)
			} else {
				highCardJokers = append(highCardJokers, hand)
			}
		}
	}

	SortByFirstDiffChar(fiveOfAKind)
	SortByFirstDiffChar(fourOfAKind)
	SortByFirstDiffChar(fullHouse)
	SortByFirstDiffChar(threeOfAKind)
	SortByFirstDiffChar(twoPairs)
	SortByFirstDiffChar(onePair)
	SortByFirstDiffChar(highCard)

	handWithBidsSorted = append(handWithBidsSorted, fiveOfAKind...)
	handWithBidsSorted = append(handWithBidsSorted, fourOfAKind...)
	handWithBidsSorted = append(handWithBidsSorted, fullHouse...)
	handWithBidsSorted = append(handWithBidsSorted, threeOfAKind...)
	handWithBidsSorted = append(handWithBidsSorted, twoPairs...)
	handWithBidsSorted = append(handWithBidsSorted, onePair...)
	handWithBidsSorted = append(handWithBidsSorted, highCard...)

	length := len(handWithBidsSorted)

	for index, hand := range handWithBidsSorted {
		bid, _ := strconv.Atoi(hand[5])
		rank := length - index
		totalWinnings += rank * bid
	}

	fmt.Println(totalWinnings)

	StrengthMap["J"] = 1

	SortByFirstDiffChar(fiveOfAKindJokers)
	SortByFirstDiffChar(fourOfAKindJokers)
	SortByFirstDiffChar(fullHouseJokers)
	SortByFirstDiffChar(threeOfAKindJokers)
	SortByFirstDiffChar(twoPairsJokers)
	SortByFirstDiffChar(onePairJokers)
	SortByFirstDiffChar(highCardJokers)

	handWithBidsSortedJokers = append(handWithBidsSortedJokers, fiveOfAKindJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, fourOfAKindJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, fullHouseJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, threeOfAKindJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, twoPairsJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, onePairJokers...)
	handWithBidsSortedJokers = append(handWithBidsSortedJokers, highCardJokers...)

	length2 := len(handWithBidsSortedJokers)

	for index, hand := range handWithBidsSortedJokers {
		bid, _ := strconv.Atoi(hand[5])
		rank := length2 - index
		totalWinningsJokers += rank * bid
	}

	fmt.Println(totalWinningsJokers)
	fmt.Println((length2))

}

func GetStrength(char string) int {
	if strength, ok := StrengthMap[char]; ok {
		return strength
	}
	// Convert string to integer for numeric characters.
	num, err := strconv.Atoi(char)
	if err == nil {
		return num
	}
	// Return 0 for unrecognized characters.
	return 0
}

func CompareFirstDiffChar(subSlice1, subSlice2 []string) bool {
	for i := 0; i < len(subSlice1) && i < len(subSlice2); i++ {
		if subSlice1[i] != subSlice2[i] {
			return GetStrength(subSlice1[i]) > GetStrength(subSlice2[i])
		}
	}
	return len(subSlice1) > len(subSlice2)
}

func SortByFirstDiffChar(data [][]string) {
	sort.Slice(data, func(i, j int) bool {
		return CompareFirstDiffChar(data[i], data[j])
	})
}

func isFiveOfAKind(hand []string) bool {
	for i := 1; i < 5; i++ {
		if hand[0] != hand[i] {
			return false
		}
	}
	return true
}

func isForOfAKind(hand []string) bool {
	elementCounts := make(map[string]int)
	for _, element := range hand[:5] {
		elementCounts[element]++
	}

	foundFour := false
	for _, count := range elementCounts {
		if count == 4 {
			foundFour = true
		}
	}
	return foundFour
}

func isFullHouse(hand []string) bool {
	elementCounts := make(map[string]int)
	for _, element := range hand[:5] {
		elementCounts[element]++
	}

	foundThree := false
	foundTwo := false
	for _, count := range elementCounts {
		if count == 3 {
			foundThree = true
		}
		if count == 2 {
			foundTwo = true
		}
	}
	return foundThree && foundTwo
}

func isThreeOfAKind(hand []string) bool {
	elementCounts := make(map[string]int)
	for _, element := range hand[:5] {
		elementCounts[element]++
	}

	foundThree := false
	for _, count := range elementCounts {
		if count == 3 {
			foundThree = true
		}
	}
	return foundThree
}

func isTwoPairs(hand []string) bool {
	seenChars := make(map[string]bool)
	pairsFound := 0
	for _, char := range hand {
		if seenChars[char] {
			pairsFound++
			delete(seenChars, char)
		} else {
			seenChars[char] = true
		}
	}
	return pairsFound == 2
}

func isOnePair(hand []string) bool {
	elementCounts := make(map[string]int)
	for _, element := range hand[:5] {
		elementCounts[element]++
	}

	foundTwo := false
	for _, count := range elementCounts {
		if count == 2 {
			foundTwo = true
		}
	}
	return foundTwo
}

func countJokers(hand []string) int {
	count := 0
	for _, item := range hand {
		if item == "J" {
			count++
		}
	}
	return count
}

// if one J in poker then if 5, if 4 J then 5

// already full house, if one J then 5

// three of a kind, if one J then poker, if 2 J then 5, if 3 J then poker

// two pairs, if one J then full house, if 2 J then poker

// one pair, if one J then three of a kind, if 2 J then three of a kind

// high card, if one J then one pair
