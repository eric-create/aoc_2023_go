package main

import (
	"eric-create/aoc_2023/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")
	cards := Cards(lines)
	fmt.Println(Points(cards))
}

func Points(cards [][2][]int) int {
	sum := 0

	for _, card := range cards {
		matches := 0
		winners := card[0]
		ticket := card[1]

		for _, number := range ticket {
			if slices.Contains[[]int](winners, number) {
				fmt.Print(number, " ")
				matches++
			}
		}
		fmt.Println()

		sum += int(math.Pow(2, float64(matches-1)))
	}

	return sum
}

func Cards(lines []string) [][2][]int {
	cards := [][2][]int{}

	for _, line := range lines {
		card := Card(line)
		cards = append(cards, card)
	}

	return cards
}

func Card(line string) [2][]int {
	parts := strings.Split(line, ":")
	parts = strings.Split(parts[1], "|")
	winners := Numbers(parts[0])
	ticket := Numbers(parts[1])
	return [2][]int{winners, ticket}
}

func Numbers(s string) []int {
	numbers := []int{}
	parts := strings.Split(s, " ")

	for _, part := range parts {
		if part != "" {
			number, _ := strconv.Atoi(part)
			numbers = append(numbers, number)
		}
	}

	return numbers
}
