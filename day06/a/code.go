package main

import (
	"eric-create/aoc_2023/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")
	records := Records(lines)
	Compete(records)
}

func Compete(records [][2]int) {
	sum := 1
	for _, record := range records {
		maxTime := record[0]
		recordDistance := record[1]
		newRecords := NewRecords(maxTime, recordDistance)
		sum *= len(*newRecords)
	}
	fmt.Println(sum)
}

// Returns a list of "records" where the first value is the time and the second the distance.
func Records(lines []string) [][2]int {
	records := [][2]int{}

	timesString := strings.Split(lines[0], ":")[1]
	times := utils.LineToNumbers(timesString)

	distancesString := strings.Split(lines[1], ":")[1]
	distances := utils.LineToNumbers(distancesString)

	for i := 0; i < len(times); i++ {
		records = append(records, [2]int{times[i], distances[i]})
	}

	return records
}

func Distance(time int, maxTime int) int {
	return (maxTime - time) * time
}

func NewRecords(maxTime, recordDistance int) *[][2]int {
	newRecords := [][2]int{}

	for pushTime := 0; pushTime <= maxTime; pushTime++ {
		newDistance := Distance(pushTime, maxTime)

		if newDistance > recordDistance {
			newRecords = append(newRecords, [2]int{pushTime, newDistance})
		}
	}

	return &newRecords
}
