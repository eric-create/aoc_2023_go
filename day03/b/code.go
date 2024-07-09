package main

import (
	"eric-create/aoc_2023/nodes"
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadLines("input.txt")
	field := nodes.NodeField(utils.StringField(lines))
	// nodes.PrintNodeField(field, []*nodes.Node{})
	fmt.Println()
	symbols := FindSymbols(field)
	// nodes.PrintNodeField(field, symbols)
	fmt.Println()
	numbers, gears := AllNumbers(symbols)

	sum := 0
	for _, number := range numbers {
		n, _ := strconv.Atoi(nodes.SequenceToString(number))
		sum += n
	}

	fmt.Println("a", sum)

	sum = 0
	for _, gear := range gears {
		if gear[0] != nil && gear[1] != nil {
			gear0, _ := strconv.Atoi(nodes.SequenceToString(gear[0]))
			gear1, _ := strconv.Atoi(nodes.SequenceToString(gear[1]))
			sum += gear0 * gear1
		}
	}
	fmt.Println("b", sum)
}

func FindSymbols(field [][]*nodes.Node) []*nodes.Node {
	sequence := utils.FieldToSequence[*nodes.Node](field)
	symbols := []*nodes.Node{}

	for _, node := range sequence {
		if node.IsSymbol() && !node.IsNumber() {
			symbols = append(symbols, node)
		}
	}

	return symbols
}

func AllNumbers(symbols []*nodes.Node) ([][]*nodes.Node, [][2][]*nodes.Node) {
	allNumbers := [][]*nodes.Node{}
	gears := [][2][]*nodes.Node{}

	for _, symbol := range symbols {
		adjacentNumbers, gear := AdjacentNumbers(symbol)
		allNumbers = append(allNumbers, adjacentNumbers...)
		gears = append(gears, gear)
	}

	return allNumbers, gears
}

func AdjacentNumbers(start *nodes.Node) ([][]*nodes.Node, [2][]*nodes.Node) {
	numbers := [][]*nodes.Node{}
	gear := [2][]*nodes.Node{}

	for _, neighbor := range start.RealNeighbors(vectors.AllDirections()) {
		if neighbor.IsNumber() && !neighbor.Covered {
			neighbor.SetCovered()
			number := []*nodes.Node{neighbor}
			TrackNumber(neighbor, &number)
			number = nodes.SortHorizontallyAscending(number)
			numbers = append(numbers, number)
		}
	}

	if len(numbers) == 2 {
		gear = [2][]*nodes.Node{numbers[0], numbers[1]}
	}

	return numbers, gear
}

func TrackNumber(current *nodes.Node, number *[]*nodes.Node) {
	for _, neighbor := range current.RealNeighbors(vectors.Horizontal()) {
		if neighbor.IsNumber() && !neighbor.Covered {
			neighbor.SetCovered()
			*number = append(*number, neighbor)
			TrackNumber(neighbor, number)
		}
	}
}
