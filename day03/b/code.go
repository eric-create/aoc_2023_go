package main

import (
	"eric-create/aoc_2023/nodes"
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
)

func main() {
	lines := utils.ReadLines("input.txt")
	field := nodes.NodeField(utils.RuneField(lines))
	nodes.PrintNodeField(field, []*nodes.Node{})
	fmt.Println()
	symbols := FindSymbols(field)
	nodes.PrintNodeField(field, symbols)
	fmt.Println()
	numbers := AllNumbers(symbols)

	for _, number := range numbers {
		fmt.Println(nodes.SequenceToString(number))
	}
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

func AllNumbers(symbols []*nodes.Node) [][]*nodes.Node {
	allNumbers := [][]*nodes.Node{}

	for _, symbol := range symbols {
		adjacentNumbers := AdjacentNumbers(symbol)
		allNumbers = append(allNumbers, adjacentNumbers...)
	}

	return allNumbers
}

func AdjacentNumbers(start *nodes.Node) [][]*nodes.Node {
	numbers := [][]*nodes.Node{}

	for _, neighbor := range start.RealNeighbors(vectors.AllDirections()) {
		if neighbor.IsNumber() {
			neighbor.SetCovered()
			number := []*nodes.Node{neighbor}
			TrackNumber(neighbor, &number)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func TrackNumber(current *nodes.Node, number *[]*nodes.Node) {
	for _, neighbor := range current.RealNeighbors(vectors.Horizontal()) {
		if neighbor.IsNumber() && !neighbor.Covered {
			neighbor.SetCovered()
			*number = append(*number, neighbor)
			TrackNumber(neighbor, number)
		} else {
			return
		}
	}
}
