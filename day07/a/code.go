package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	lines, err := ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	players := GetPlayers(*lines)
	players = SortPlayers(*players)
	sum := GetResult(players)

	for _, player := range *players {
		player.Print()
	}
	fmt.Println()
	fmt.Println(len(*players))
	fmt.Println(sum)
}

func GetResult(players *[]*Player) int {
	sum := 0

	for i, player := range *players {
		sum += (i + 1) * player.Bid
	}
	return sum
}

func SortPlayers(players []*Player) *[]*Player {
	sortedPlayers := []*Player{}

	for _, player := range players {

		if len(sortedPlayers) == 0 {
			sortedPlayers = append(sortedPlayers, player)

		} else {
			for j, sortedPlayer := range sortedPlayers {

				if player.Compare(sortedPlayer) == Lower {
					sortedPlayers = slices.Insert(sortedPlayers, j, player)
					break

				} else if j == len(sortedPlayers)-1 {
					sortedPlayers = append(sortedPlayers, player)
				}
			}
		}
	}

	return &sortedPlayers
}

type Player struct {
	Hand        *[]int
	Bid         int
	Combination int
}

func NewPlayer(hand *[]int, bid int) *Player {
	player := Player{hand, bid, -1}
	combination := GetCombination(&player)
	player.Combination = combination

	return &player
}

const (
	Higher int = iota
	Lower
	Equals
)

func (p *Player) Compare(other *Player) int {
	if p.Combination < other.Combination {
		return Lower
	} else if p.Combination > other.Combination {
		return Higher
	}

	for i := range [5]int{} {
		if (*p.Hand)[i] < (*other.Hand)[i] {
			return Lower
		} else if (*p.Hand)[i] > (*other.Hand)[i] {
			return Higher
		}
	}
	return Equals
}

func (p *Player) Print() {
	fmt.Println((*p).Combination, (*p).Bid, *((*p).Hand))
}

const (
	HighCard     int = iota // 5 distinct
	OnePair                 // 4 distinct
	TwoPair                 // 3 distinct
	ThreeOfAKind            // 3 distinct
	FullHouse               // 2 distinct
	FourOfAKind             // 2 distinct
	FiveOfAKind             // 1 distinct
)

func GetCombination(player *Player) int {
	hand := *((*player).Hand)
	set := mapset.NewSet(hand...)

	switch len(set.ToSlice()) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		if IsThreeOfAKind(hand, set) {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	case 2:
		if IsFourOfAKind(hand, set) {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 1:
		return FiveOfAKind
	}
	panic("No no.")
}

func IsFourOfAKind(hand []int, set mapset.Set[int]) bool {
	distincts := set.ToSlice()
	counts := []int{}

	for _, distinct := range distincts {
		count := 0
		for _, card := range hand {
			if card == distinct {
				count++
			}
		}
		counts = append(counts, count)
	}

	return slices.Contains(counts, 4)
}

func IsThreeOfAKind(hand []int, set mapset.Set[int]) bool {
	distincts := set.ToSlice()
	counts := []int{}

	for _, distinct := range distincts {
		count := 0
		for _, card := range hand {
			if card == distinct {
				count++
			}
		}
		counts = append(counts, count)
	}

	return slices.Contains(counts, 3)
}

func GetPlayers(lines []string) *[]*Player {
	players := []*Player{}

	for _, line := range lines {
		players = append(players, GetPlayer(line))
	}

	return &players
}

func GetPlayer(line string) *Player {
	parts := strings.Split(line, " ")
	hand := GetHand((parts[0]))
	bid, _ := strconv.Atoi(parts[1])
	player := NewPlayer(hand, bid)

	return player
}

func GetHand(handStr string) *[]int {
	hand := []int{}

	for _, cardRune := range handStr {
		if card, err := strconv.Atoi(string(cardRune)); err != nil {
			switch cardRune {
			case 'T':
				hand = append(hand, 10)
			case 'J':
				hand = append(hand, 11)
			case 'Q':
				hand = append(hand, 12)
			case 'K':
				hand = append(hand, 13)
			case 'A':
				hand = append(hand, 14)
			}
		} else {
			hand = append(hand, card)
		}
	}

	return &hand
}

func ReadLines(path string) (*[]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return &lines, nil
}
