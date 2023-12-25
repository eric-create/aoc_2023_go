package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadLines("input.txt")

	games := Games(lines)

	redLim := 12
	greenLim := 13
	blueLim := 14

	possibleGames := []*Game{}

	sum := 0

	for _, game := range *games {
		if game.maxRed <= redLim && game.maxGreen <= greenLim && game.maxBlue <= blueLim {
			possibleGames = append(possibleGames, game)
			sum += game.ID
		}
	}
	fmt.Println(len(possibleGames))
	fmt.Println(sum)
}

func Games(lines []string) *[]*Game {
	games := []*Game{}
	for _, line := range lines {
		game := NewGame(line)
		games = append(games, game)
	}

	return &games
}

func NewGame(line string) *Game {
	id, processedLine := GetID(line)
	roundStrings := strings.Split(processedLine, ";")

	rounds := [][3]int{}
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, roundString := range roundStrings {
		values := strings.Split(roundString, ",")
		round := [3]int{}

		for _, value := range values {
			valueElements := strings.Split(value, " ")
			valueElements = valueElements[1:]
			count, _ := strconv.Atoi(valueElements[0])

			switch color := valueElements[1]; color {
			case "red":
				round[0] = count
				if count > maxRed {
					maxRed = count
				}
			case "green":
				round[1] = count
				if count > maxGreen {
					maxGreen = count
				}
			case "blue":
				round[2] = count
				if count > maxBlue {
					maxBlue = count
				}
			}
		}

		rounds = append(rounds, round)
	}

	game := Game{id, rounds, maxRed, maxGreen, maxBlue}
	return &game
}

func GetID(line string) (int, string) {
	parts := strings.Split(line, ":")
	firstParts := strings.Split(parts[0], " ")
	id, _ := strconv.Atoi(firstParts[1])
	return id, parts[1]
}

type Game struct {
	ID       int
	rounds   [][3]int // red, green, blue
	maxRed   int
	maxGreen int
	maxBlue  int
}

func ReadLines(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n")
}
