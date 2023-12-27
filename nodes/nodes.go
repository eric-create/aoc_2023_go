package nodes

import (
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
	"slices"
)

type Node struct {
	Position    [2]int
	Symbol      rune
	Covered     bool
	Neighbors   [3][3]*Node
	Connections []*Node
}

func (n *Node) IsSymbol() bool {
	return n.Symbol != '.'
}

func (n *Node) IsNumber() bool {
	return utils.IsNumber(n.Symbol)
}

// Sets the attribute `Covered` to `true`.
func (n *Node) SetCovered() {
	n.Covered = true
}

func (n *Node) String() string {
	return string(n.Symbol)
}

func (n *Node) SetNeighbor(neighbor *Node, direction [2]int) {
	x := direction[vectors.X()] + 1
	y := direction[vectors.Y()] + 1
	n.Neighbors[y][x] = neighbor
}

func (n *Node) GetNeighbor(direction [2]int) *Node {
	x := direction[vectors.X()] + 1
	y := direction[vectors.Y()] + 1
	return n.Neighbors[y][x]
}

// Gets all neighbor nodes that actually exist.
func (n *Node) RealNeighbors(directions [][2]int) []*Node {
	neighbors := []*Node{}

	for _, direction := range directions {
		if neighbor := n.GetNeighbor(direction); neighbor != nil {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func NodeField(runeField [][]rune) [][]*Node {
	nodeField := [][]*Node{}

	for y := range runeField {
		nodeField = append(nodeField, []*Node{})

		for x := range runeField[y] {
			node := Node{
				Position: [2]int{x, y},
				Symbol:   runeField[y][x],
			}

			nodeField[y] = append(nodeField[y], &node)
		}
	}

	DetermineNeighbors(nodeField)

	return nodeField
}

func DetermineNeighbors(field [][]*Node) {
	for y := range field {
		for x := range field[y] {
			node := field[y][x]

			for _, direction := range vectors.AllDirections() {
				neighborPosition := utils.Navigate[*Node](field, node.Position, direction)

				if neighborPosition != nil {
					neighborX := neighborPosition[vectors.X()]
					neighborY := neighborPosition[vectors.Y()]
					neighbor := field[neighborY][neighborX]
					node.SetNeighbor(neighbor, direction)
				}
			}
		}
	}
}

func PrintNodeField(field [][]*Node, selection []*Node) {
	for y := range field {
		for x := range field[y] {
			node := field[y][x]

			if len(selection) == 0 {
				fmt.Print(node.String())

			} else {
				if slices.Contains[[]*Node](selection, node) {
					fmt.Print(node.String())
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func SequenceToString(sequence []*Node) string {
	word := ""

	for _, node := range sequence {
		word += node.String()
	}

	return word
}

func SortHorizontallyAscending(sequence []*Node) []*Node {
	sortedNodes := []*Node{}

	for _, node := range sequence {
		if len(sortedNodes) == 0 {
			sortedNodes = append(sortedNodes, node)

		} else {
			for i, sorted := range sortedNodes {
				if node.Position[vectors.X()] < sorted.Position[vectors.X()] {
					sortedNodes = slices.Insert[[]*Node](sortedNodes, i, node)
					break
				} else if i == len(sortedNodes)-1 {
					sortedNodes = append(sortedNodes, node)
				}
			}
		}
	}

	return sortedNodes
}
