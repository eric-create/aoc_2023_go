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
		}
	}

	return numbers
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
