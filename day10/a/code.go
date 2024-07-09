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
	field := Field(lines)

	startNode := nodes.Find(field, "S")
	startTile := NewTile(startNode, 0)
	loop := []*Tile{}
	meeting := Walk([]*Tile{startTile}, &loop)

	// PrintField(field, loop)
	fmt.Println(meeting.Distance)
}

func Field(lines []string) [][]*nodes.Node {
	stringField := utils.StringField(lines)
	nodeField := nodes.NodeField(stringField)
	return nodeField
}

func PrintField(field [][]*nodes.Node, loop []*Tile) {
	for _, row := range field {
		for _, node := range row {
			symbol := "."

			for _, tile := range loop {
				if tile.Node == node {
					symbol = strconv.Itoa(tile.Distance)
					break
				}
			}
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

type Tile struct {
	Origin   *Tile
	Vectors  []vectors.Vector
	Node     *nodes.Node
	Distance int
}

func NewTile(node *nodes.Node, distance int) *Tile {
	var tileVectors []vectors.Vector

	switch node.Symbol {
	case "|":
		tileVectors = vectors.Vertical()
	case "-":
		tileVectors = vectors.Horizontal()
	case "L":
		tileVectors = []vectors.Vector{vectors.Up(), vectors.Right()}
	case "J":
		tileVectors = []vectors.Vector{vectors.Up(), vectors.Left()}
	case "7":
		tileVectors = []vectors.Vector{vectors.Down(), vectors.Left()}
	case "F":
		tileVectors = []vectors.Vector{vectors.Down(), vectors.Right()}
	case ".":
		tileVectors = []vectors.Vector{}
	case "S":
		tileVectors = vectors.ManhattanDirections()
	}

	return &Tile{nil, tileVectors, node, distance}
}

func (t *Tile) Neighbors() []*Tile {
	neighborNodes := t.Node.RealNeighbors(vectors.ManhattanDirections())
	neighborTiles := []*Tile{}

	for _, neighborNode := range neighborNodes {
		if t.Origin == nil || (t.Origin != nil && neighborNode.Position != t.Origin.Node.Position) {
			newTile := NewTile(neighborNode, -1)
			newTile.Origin = t
			neighborTiles = append(neighborTiles, newTile)
		}
	}

	return neighborTiles
}

func (t *Tile) PointsTo(other *Tile) bool {
	for _, vector := range t.Vectors {
		if t.Node.Position.Add(vector) == other.Node.Position {
			return true
		}
	}
	return false
}

func (t *Tile) ConnectsWith(other *Tile) bool {
	return t.PointsTo(other) && other.PointsTo(t)
}

func Walk(stack []*Tile, loop *[]*Tile) *Tile {

	// This condition should never be `true`!
	if len(stack) == 0 {
		return nil
	}

	currentTile := stack[0]

	if currentTile.Node.Covered {
		return Walk(stack[1:], loop)
	}

	currentTile.Node.SetCovered()
	neighborTiles := currentTile.Neighbors()

	for _, neighborTile := range neighborTiles {
		if currentTile.ConnectsWith(neighborTile) {
			if neighborTile.Node.Covered {
				return currentTile
			}
			neighborTile.Origin = currentTile
			neighborTile.Distance = currentTile.Distance + 1
			stack = append(stack, neighborTile)
			*loop = append(*loop, neighborTile)
		}
	}

	return Walk(stack[1:], loop)
}
