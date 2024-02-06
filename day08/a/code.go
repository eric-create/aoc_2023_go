package main

import (
	"eric-create/aoc_2023/utils"
)

func main() {
	lines := utils.ReadLines("input.txt")
	instructions, nodes := ProcessInput(lines)
	WalkNodes(instructions, nodes)
}

func WalkNodes(instructions []rune, nodes []*Node) {
	start := FindStart(nodes)
}

func FindStart(nodes []*Node) int {
	for i, node := range nodes {
		if node.Name == "AAA" {
			return i
		}
	}
	return -1
}

func ProcessInput(lines []string) ([]rune, []*Node) {
	parts := utils.SplitParagraphs(lines)
	instructions := GetInstructions(parts[0][0])
	nodes := GetNodeList(parts[1])
	return instructions, nodes
}

func GetInstructions(line string) []rune {
	instructions := []rune{}

	for _, r := range line {
		instructions = append(instructions, r)
	}

	return instructions
}

func GetNodeList(lines []string) []*Node {
	nodeList := []*Node{} // Ordered as seen in the input file
	allNodes := []*Node{} // Unordered

	firstNode := LineToNode(lines[0], &allNodes)
	nodeList = append(nodeList, firstNode)

	for _, line := range lines[1:] {
		node := LineToNode(line, &allNodes)
		nodeList = append(nodeList, node)
	}

	return nodeList
}

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func NewNode(name string, allNodes *[]*Node) *Node {
	node := Node{Name: name}
	*allNodes = append(*allNodes, &node)
	return &node
}

func GetNode(name string, allNodes *[]*Node) *Node {
	for _, node := range *allNodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func LineToNode(line string, allNodes *[]*Node) *Node {

	name := line[0:2]
	left := line[7:9]
	right := line[12:14]

	var newNode *Node = nil

	// Create the new node
	if node := GetNode(name, allNodes); node != nil {
		newNode = node
	} else {
		newNode = NewNode(name, allNodes)
	}

	// Left node
	if leftNode := GetNode(left, allNodes); leftNode != nil {
		newNode.Left = leftNode
	} else {
		newNode.Left = NewNode(left, allNodes)
	}

	// Right node
	if rightNode := GetNode(left, allNodes); rightNode != nil {
		newNode.Right = rightNode
	} else {
		newNode.Right = NewNode(right, allNodes)
	}

	return newNode
}
