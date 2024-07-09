package main

import (
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadLines("input.txt")
	field := utils.StringsField(lines)
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

func UniqueNumbers(numbers [][]vectors.Vector) [][]vectors.Vector {
	uniques := [][]vectors.Vector{numbers[0]}

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

func NumberEquals(a, b []vectors.Vector) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].X != b[i].X || a[i].Y != b[i].Y {
			return false
		}
	}

	return true
}

func AllPositions(structure [][]vectors.Vector) []vectors.Vector {
	flattened := []vectors.Vector{}

	for i := range structure {
		flattened = append(flattened, structure[i]...)
	}

	return flattened
}

func NumberStarts(field [][]string, adjacentNumbers []vectors.Vector) []vectors.Vector {
	allNumbers := []vectors.Vector{}

	for _, adjacentNumber := range adjacentNumbers {
		numberStart := NumberStart(field, &adjacentNumber)
		allNumbers = append(allNumbers, numberStart)
	}

	return allNumbers
}

func NumberStart(field [][]string, position *vectors.Vector) vectors.Vector {
	nextPosition := utils.Navigate[string](field, *position, vectors.Vector{X: -1, Y: 0})

	// The edge of the board was reached.
	if nextPosition == nil {
		return *position
	}

	char := field[(*nextPosition).Y][(*nextPosition).X]
	if !utils.StringIsNumber(char) {
		return *position
	}

	return NumberStart(field, nextPosition)
}

func Numbers(field [][]string, numberStarts []vectors.Vector) []int {
	numberInts := []int{}
	numbers := ParseNumbers(field, numberStarts)

	for _, number := range numbers {
		integer := NumberToInt(field, number)
		numberInts = append(numberInts, integer)
	}

	return numberInts
}

func NumberToInt(field [][]string, number []vectors.Vector) int {
	characters := []string{}

	for _, position := range number {
		characters = append(characters, field[position.Y][position.X])
	}

	word := ""
	for _, character := range characters {
		word += string(character)
	}

	integer, _ := strconv.Atoi(word)
	return integer
}

func ParseNumbers(field [][]string, numberStarts []vectors.Vector) [][]vectors.Vector {
	numbers := [][]vectors.Vector{}

	for _, numberStart := range numberStarts {
		number := ParseNumber(field, &numberStart)
		numbers = append(numbers, number)
	}

	return numbers
}

func ParseNumber(field [][]string, position *vectors.Vector) []vectors.Vector {
	numbers := []vectors.Vector{*position}

	for {
		position = utils.Navigate[string](field, *position, vectors.Vector{X: 1, Y: 0})
		// The edge of the board was reached.
		if position == nil {
			break
		}

		char := field[(*position).Y][(*position).X]
		if !utils.StringIsNumber(char) {
			break
		}

		numbers = append(numbers, *position)
	}

	return numbers
}

func PrintField(field [][]rune, positions []vectors.Vector) {
	for y := range field {
		for x := range field[y] {
			matches := false

			for _, position := range positions {
				if position.X == x && position.Y == y {
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

func AllAdjacentNumbers(field [][]string, symbols []vectors.Vector) []vectors.Vector {
	allNumbers := []vectors.Vector{}
	directions := vectors.AllDirections()

	for _, symbol := range symbols {
		if adjacentNumbers := AdjacentNumbers(field, symbol, directions); adjacentNumbers != nil {
			allNumbers = append(allNumbers, adjacentNumbers...)
		}
	}

	return allNumbers
}

func AdjacentNumbers(field [][]string, position vectors.Vector, directions []vectors.Vector) []vectors.Vector {
	numbers := []vectors.Vector{}

	for _, direction := range directions {
		newPosition := utils.Navigate[string](field, position, direction)

		if newPosition != nil {
			symbol := field[(*newPosition).Y][(*newPosition).X]

			if utils.StringIsNumber(symbol) {
				numbers = append(numbers, position.Add(direction))
			}
		}
	}

	return numbers
}

func SymbolLocations(field [][]string) []vectors.Vector {
	locations := []vectors.Vector{}

	for y := range field {
		for x := range field[y] {
			if utils.StringIsNumber(field[y][x]) || field[y][x] == "." {
				continue
			} else {
				locations = append(locations, vectors.Vector{X: x, Y: y})
			}
		}
	}

	return locations
}
