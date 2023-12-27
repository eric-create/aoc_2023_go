package main

import (
	"eric-create/aoc_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")
	cards := Cards(lines)
	table := Table(cards)
	fmt.Println(Process(table))
}

func Process(table [][2]int) int {
	for i, entry := range table {
		matches := entry[0]
		numOfCard := entry[1]

		for j := 0; j < numOfCard; j++ {
			for k := i + 1; k <= (i+matches) && k < len(table); k++ {
				table[k][1]++
			}
		}
	}

	sum := 0
	for _, entry := range table {
		sum += entry[1]
	}
	return sum
}

func Table(cards [][2][]int) [][2]int {
	table := [][2]int{}

	for _, card := range cards {
		matches := 0
		winners := card[0]
		ticket := card[1]

		for _, number := range ticket {
			if slices.Contains[[]int](winners, number) {
				matches++
			}
		}

		table = append(table, [2]int{matches, 1})
	}

	return table
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
