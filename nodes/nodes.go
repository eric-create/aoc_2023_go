package nodes

import (
	"eric-create/aoc_2023/utils"
	"eric-create/aoc_2023/vectors"
	"fmt"
	"slices"
)

type Node struct {
	Position    vectors.Vector
	Symbol      string
	Covered     bool
	Neighbors   [3][3]*Node
	Connections []*Node
}

func (n *Node) IsSymbol() bool {
	return n.Symbol != "."
}

func (n *Node) IsNumber() bool {
	return utils.StringIsNumber(n.Symbol)
}

// Sets the attribute `Covered` to `true`.
func (n *Node) SetCovered() {
	n.Covered = true
}

func (n *Node) String() string {
	return string(n.Symbol)
}

func (n *Node) SetNeighbor(neighbor *Node, direction vectors.Vector) {
	n.Neighbors[direction.Y+1][direction.X+1] = neighbor
}

func (n *Node) GetNeighbor(direction vectors.Vector) *Node {
	return n.Neighbors[direction.Y+1][direction.X+1]
}

// Gets all neighbor nodes that actually exist.
func (n *Node) RealNeighbors(directions []vectors.Vector) []*Node {
	neighbors := []*Node{}

	for _, direction := range directions {
		if neighbor := n.GetNeighbor(direction); neighbor != nil {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func NodeField(stringField [][]string) [][]*Node {
	nodeField := [][]*Node{}

	for y := range stringField {
		nodeField = append(nodeField, []*Node{})

		for x := range stringField[y] {
			node := Node{
				Position: vectors.Vector{X: x, Y: y},
				Symbol:   stringField[y][x],
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
				neighborPosition := utils.Navigate(field, node.Position, direction)

				if neighborPosition != nil {
					neighbor := field[neighborPosition.Y][neighborPosition.X]
					node.SetNeighbor(neighbor, direction)
				}
			}
		}
	}
}

// If `selection` is empty, then prints all nodes.
func PrintNodeField(field [][]*Node, selection []*Node) {
	for y := range field {
		for x := range field[y] {
			node := field[y][x]

			if len(selection) == 0 {
				fmt.Print(node.String())

			} else {
				if slices.Contains(selection, node) {
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
				if node.Position.X < sorted.Position.X {
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

func Find(field [][]*Node, symbol string) *Node {
	for _, row := range field {
		for _, node := range row {
			if node.Symbol == symbol {
				return node
			}
		}
	}
	return nil
}
