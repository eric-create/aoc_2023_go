package main

import (
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadLines("input.txt")
	field := utils.RuneField(lines)
	symbolLocations := SymbolLocations(field)
	adjacentNumbers := AllAdjacentNumbers(field, symbolLocations)
	// PrintField(field, adjacentNumbers)
	numberStarts := NumberStarts(field, adjacentNumbers)
	numberStarts = utils.UniquePositions(numberStarts)
	numbers := Numbers(field, numberStarts)
	fmt.Println(numbers)

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)
}

func UniqueNumbers(numbers [][][2]int) [][][2]int {
	uniques := [][][2]int{numbers[0]}

	for _, number := range numbers[1:] {

		contains := false
		for _, unique := range uniques {
			if NumberEquals(number, unique) {
				contains = true
				break
			}
		}

		if !contains {
			uniques = append(uniques, number)
		}
	}

	return uniques
}

func NumberEquals(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i][X] != b[i][X] || a[i][Y] != b[i][Y] {
			return false
		}
	}

	return true
}

func AllPositions(structure [][][2]int) [][2]int {
	flattened := [][2]int{}

	for i := range structure {
		flattened = append(flattened, structure[i]...)
	}

	return flattened
}

func NumberStarts(field [][]rune, adjacentNumbers [][2]int) [][2]int {
	allNumbers := [][2]int{}

	for _, adjacentNumber := range adjacentNumbers {
		numberStart := NumberStart(field, &adjacentNumber)
		allNumbers = append(allNumbers, numberStart)
	}

	return allNumbers
}

func NumberStart(field [][]rune, position *[2]int) [2]int {
	nextPosition := utils.Navigate[rune](field, *position, [2]int{-1, 0})

	// The edge of the board was reached.
	if nextPosition == nil {
		return *position
	}

	char := field[(*nextPosition)[Y]][(*nextPosition)[X]]
	if !utils.IsNumber(char) {
		return *position
	}

	return NumberStart(field, nextPosition)
}

func Numbers(field [][]rune, numberStarts [][2]int) []int {
	numberInts := []int{}
	numbers := ParseNumbers(field, numberStarts)

	for _, number := range numbers {
		integer := NumberToInt(field, number)
		numberInts = append(numberInts, integer)
	}

	return numberInts
}

func NumberToInt(field [][]rune, number [][2]int) int {
	characters := []rune{}

	for _, position := range number {
		characters = append(characters, field[position[Y]][position[X]])
	}

	word := ""
	for _, character := range characters {
		word += string(character)
	}

	integer, _ := strconv.Atoi(word)
	return integer
}

func ParseNumbers(field [][]rune, numberStarts [][2]int) [][][2]int {
	numbers := [][][2]int{}

	for _, numberStart := range numberStarts {
		number := ParseNumber(field, &numberStart)
		numbers = append(numbers, number)
	}

	return numbers
}

func ParseNumber(field [][]rune, position *[2]int) [][2]int {
	numbers := [][2]int{*position}

	for {
		position = utils.Navigate[rune](field, *position, [2]int{1, 0})
		// The edge of the board was reached.
		if position == nil {
			break
		}

		char := field[(*position)[Y]][(*position)[X]]
		if !utils.IsNumber(char) {
			break
		}

		numbers = append(numbers, *position)
	}

	return numbers
}

func PrintField(field [][]rune, positions [][2]int) {
	for y := range field {
		for x := range field[y] {
			matches := false

			for _, position := range positions {
				if position[X] == x && position[Y] == y {
					matches = true
					break
				}
			}

			if matches {
				fmt.Print(string(field[y][x]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

const (
	X = 0
	Y = 1
)

func AllAdjacentNumbers(field [][]rune, symbols [][2]int) [][2]int {
	allNumbers := [][2]int{}
	directions := vectors.AllDirections()

	for _, symbol := range symbols {
		if adjacentNumbers := AdjacentNumbers(field, symbol, directions); adjacentNumbers != nil {
			allNumbers = append(allNumbers, adjacentNumbers...)
		}
	}

	return allNumbers
}

func AdjacentNumbers(field [][]rune, position [2]int, directions [][2]int) [][2]int {
	numbers := [][2]int{}

	for _, direction := range directions {
		newPosition := utils.Navigate[rune](field, position, direction)

		if newPosition != nil {
			symbol := field[(*newPosition)[Y]][(*newPosition)[X]]

			if utils.IsNumber(symbol) {
				numbers = append(numbers, *vectors.Add(position, direction))
			}
		}
	}

	return numbers
}

func SymbolLocations(field [][]rune) [][2]int {
	locations := [][2]int{}

	for y := range field {
		for x := range field[y] {
			if utils.IsNumber(field[y][x]) || field[y][x] == '.' {
				continue
			} else {
				locations = append(locations, [2]int{x, y})
			}
		}
	}

	return locations
}
