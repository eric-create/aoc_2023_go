package main

import (
	"eric-create/aoc_2023/utils"
	"fmt"
	"math"
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
		time := record[0]
		recordDistance := record[1]
		highPoints := HighPoints(time)
		idealTime := highPoints[0]
		myDistance := Distance(idealTime, time)
		sum *= myDistance
		fmt.Println("Time frame: ", time)
		fmt.Println("Record distance: ", recordDistance)
		fmt.Println("My distance: ", myDistance)
		fmt.Println()
	}
	fmt.Println(sum)
}

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

func Distance(time int, max_time int) int {
	return (max_time - time) * time
}

func HighPoints(time int) []int {
	highPoint := time / 2
	remainder := math.Remainder(float64(time), 2)

	if remainder == float64(0.5) {
		return []int{int(highPoint), int(highPoint + 1)}
	} else if remainder < float64(0.5) {
		return []int{int(highPoint)}
	}
	return []int{int(highPoint) + 1}
}
