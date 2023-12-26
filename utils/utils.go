package utils

import (
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

func IsNumber(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err != nil {
		return false
	}
	return true
}

func RuneField(lines []string) [][]rune {
	field := [][]rune{}

	for y, line := range lines {
		field = append(field, []rune{})
		for _, char := range line {
			field[y] = append(field[y], char)
		}
	}

	return field
}

func AddVectors(a [2]int, b [2]int) *[2]int {
	newVector := [2]int{a[0] + b[0], a[1] + b[1]}
	return &newVector
}

func AllDirections() [][2]int {
	return [][2]int{
		{0, -1},  // up
		{1, -1},  // up-right
		{1, 0},   // right
		{1, 1},   // down-right
		{0, 1},   // down
		{-1, 1},  // down-left
		{-1, 0},  // left
		{-1, -1}, // up-left
	}
}

func Navigate[T any](field [][]T, position *[2]int, direction [2]int) *[2]int {
	xMax := len(field[0]) - 1
	yMax := len(field) - 1

	xPos := position[0]
	yPos := position[1]

	xDir := direction[0]
	yDir := direction[1]

	xNew := xPos + xDir
	yNew := yPos + yDir

	if xNew < 0 || xNew > xMax || yNew < 0 || yNew > yMax {
		return nil
	}

	return &[2]int{xNew, yNew}
}

func UniquePositions(positions [][2]int) [][2]int {
	uniques := [][2]int{}

	for _, position := range positions {
		contains := false

		for _, unique := range uniques {
			if position[0] == unique[0] && position[1] == unique[1] {
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
