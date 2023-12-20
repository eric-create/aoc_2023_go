package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"
)

func main() {
	lines, err := ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	players := GetPlayers(*lines)
	players = SortPlayers(*players)
	sum := GetResult(players)

	// for _, player := range *players {
	// 	player.Print()
	// }
	// fmt.Println()
	// fmt.Println(len(*players))
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
	updatedHand := UseJoker(*hand)
	combination := GetCombination(updatedHand)
	player := Player{hand, bid, combination}
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

func GetCombination(hand []int) int {
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

func DistinctCards(hand []int) *map[int]int {
	distinctCards := map[int]int{}

	for _, card := range hand {
		distinctCards[card]++
	}

	return &distinctCards
}

func IsFourOfAKind(hand []int, set mapset.Set[int]) bool {
	distincts := DistinctCards(hand)
	return slices.Contains[[]int](maps.Values(*distincts), 4)
}

func IsThreeOfAKind(hand []int, set mapset.Set[int]) bool {
	distincts := DistinctCards(hand)
	return slices.Contains[[]int](maps.Values(*distincts), 3)
}

func UseJoker(hand []int) []int {
	distincts := DistinctCards(hand)
	bestCard := BestCard(distincts)

	newHand := []int{}

	for _, card := range hand {
		if card == 1 {
			newHand = append(newHand, bestCard)
		} else {
			newHand = append(newHand, card)
		}
	}

	return newHand
}

func BestCard(distincts *map[int]int) int {
	numOfDisticts := len(maps.Keys(*distincts))

	if numOfDisticts == 5 {
		// If five distinct cards, then the highest card is the best card.
		keys := maps.Keys(*distincts)
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		return keys[0]

	} else if numOfDisticts == 1 && (*distincts)[1] == 5 {
		// If five Jokers, then Ace is the best card.
		return 14

	} else {
		mostKey, _ := MostCard(distincts)
		return mostKey
	}
}

func MostCard(distincts *map[int]int) (int, int) {
	keys, values := HandLists(distincts)
	index := 0

	// If the most cards of the same symbol are Jokers.
	if keys[index] == 1 {
		index++
	}

	mostKey := keys[index]
	mostValue := values[index]
	return mostKey, mostValue
}

func HandLists(distincts *map[int]int) ([]int, []int) {
	keys := []int{}
	values := maps.Values(*distincts)

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, value := range values {
		for _, key := range maps.Keys(*distincts) {
			if (*distincts)[key] == value {
				keys = append(keys, key)
			}
		}
	}

	return keys, values
}

func CompareCombination(a, b int) int {
	if a < b {
		return Lower
	} else if a == b {
		return Equals
	}
	return Higher
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
				hand = append(hand, 1)
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
