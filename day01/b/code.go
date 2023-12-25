package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadLines("./input.txt")

	sum := 0

	for _, line := range lines {
		numbers := GetNumbers(line)
		first := strconv.Itoa(numbers[0])
		last := strconv.Itoa(numbers[len(numbers)-1])
		number, _ := strconv.Atoi(first + last)
		// fmt.Println(number)
		sum += number
	}

	fmt.Println(sum)
	// fmt.Println(lines)
}

func GetNumbers(line string) []int {
	numbers := []int{}

	for i := 0; i < len(line); i++ {
		if number := Number(string(line[i])); number > -1 {
			numbers = append(numbers, number)
		} else {
			if number := LiteralNumber(line[i:]); number > -1 {
				numbers = append(numbers, number)
			}
		}
	}

	return numbers
}

func LiteralNumber(s string) int {
	for _, word := range []string{ZERO, ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE} {
		wordLen := len(word)

		if wordLen > len(s) {
			continue
		}

		if s[:wordLen] == word {
			switch word {
			case ZERO:
				return 0
			case ONE:
				return 1
			case TWO:
				return 2
			case THREE:
				return 3
			case FOUR:
				return 4
			case FIVE:
				return 5
			case SIX:
				return 6
			case SEVEN:
				return 7
			case EIGHT:
				return 8
			case NINE:
				return 9
			}
		}
	}
	return -1
}

func Number(s string) int {
	if number, err := strconv.Atoi(s); err != nil {
		return -1
	} else {
		return number
	}
}

const (
	ZERO  = "zero"
	ONE   = "one"
	TWO   = "two"
	THREE = "three"
	FOUR  = "four"
	FIVE  = "five"
	SIX   = "six"
	SEVEN = "seven"
	EIGHT = "eight"
	NINE  = "nine"
)

func ReadLines(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n")
}
