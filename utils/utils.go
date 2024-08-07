package utils

import (
	"eric-create/aoc_2023/vectors"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n")
}

// Splits a multiline string by lines that are empty.
func SplitParagraphs(lines []string) [][]string {
	paragraphs := [][]string{{}}
	i := 0

	for _, line := range lines {
		if line != "" {
			paragraphs[i] = append(paragraphs[i], line)
		} else {
			paragraphs = append(paragraphs, []string{})
			i++
		}
	}

	return paragraphs
}

// Splits a string that consists of a single line that only contains numbers and white
// spaces by the white spaces and returns the numbers.
func LineToNumbers(line string) []int {
	numbers := []int{}
	parts := strings.Split(line, " ")

	for _, part := range parts {
		if part != "" {
			number, _ := strconv.Atoi(part)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func StringIsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func RuneIsNumber(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err != nil {
		return false
	}
	return true
}

func StringField(lines []string) [][]string {
	field := [][]string{}

	for y, line := range lines {
		field = append(field, []string{})
		for _, char := range line {
			field[y] = append(field[y], string(char))
		}
	}

	return field
}

// Returns `nil` if there is no neighbor in the specified `direction`, that means that an
// edge of `field` was reached.
func Navigate[T any](field [][]T, position, direction vectors.Vector) *vectors.Vector {
	xMax := len(field[0]) - 1
	yMax := len(field) - 1

	xNew := position.X + direction.X
	yNew := position.Y + direction.Y

	if xNew < 0 || xNew > xMax || yNew < 0 || yNew > yMax {
		return nil
	}

	new := vectors.Vector{X: xNew, Y: yNew}

	return &new
}

func FieldToSequence[T any](field [][]T) []T {
	sequence := []T{}

	for y := range field {
		sequence = append(sequence, field[y]...)
	}

	return sequence
}

func UniquePositions(positions []vectors.Vector) []vectors.Vector {
	uniques := []vectors.Vector{}

	for _, position := range positions {
		contains := false

		for _, unique := range uniques {
			if position.X == unique.X && position.Y == unique.Y {
				contains = true
				break
			}
		}

		if !contains {
			uniques = append(uniques, position)
		}
	}

	return uniques
}
