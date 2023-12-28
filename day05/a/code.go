package main

import (
	"eric-create/aoc_2023/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")
	paragraphs := utils.SplitParagraphs(lines)
	seeds, tables := ProcessParagraphs(paragraphs)
	locations := ProcessSeeds(seeds, tables)

	sort.Ints(locations)
	fmt.Println(locations[0])

	// for _, location := range locations {
	// 	fmt.Println(location)
	// }

}

type Instruction struct {
	Destination int
	Source      int
	Length      int
	Diff        int
	Index       int
}

func NewInstruction(details []int, i int) *Instruction {
	instruction := Instruction{
		Destination: details[0],
		Source:      details[1],
		Length:      details[2],
		Index:       i,
	}
	instruction.Diff = instruction.Destination - instruction.Source
	return &instruction
}

func (i *Instruction) Matches(value int) bool {
	return i.Source <= value && value < i.Source+i.Length
}

func (i *Instruction) Translate(value int) int {
	newValue := value + i.Diff
	return newValue
}

type Table struct {
	Name         string
	Index        int
	Instructions []*Instruction
}

func NewTable(paragraph []string, index int) *Table {
	header := strings.Split(paragraph[0], " ")
	nameParts := strings.Split(header[0], "-")
	name := nameParts[len(nameParts)-1]
	table := Table{Name: name, Index: index}

	for j, line := range paragraph[1:] {
		details := utils.LineToNumbers(line)
		table.Instructions = append(table.Instructions, NewInstruction(details, j))
	}

	return &table
}

func ProcessParagraphs(paragraphs [][]string) ([]int, []*Table) {
	seeds := utils.LineToNumbers(paragraphs[0][1])
	tables := []*Table{}

	for i, paragraph := range paragraphs[1:] {
		tables = append(tables, NewTable(paragraph, i))
	}

	return seeds, tables
}

func ProcessSeeds(seeds []int, tables []*Table) []int {
	locations := []int{}

	for _, seed := range seeds {
		location := ProcessSeed(seed, tables)
		locations = append(locations, location)
	}

	return locations
}

func ProcessSeed(value int, tables []*Table) int {
	newValue := value

	fmt.Println(value, "seed")
	for _, table := range tables {
		for _, instruction := range table.Instructions {
			if instruction.Matches(newValue) {
				newValue = instruction.Translate(newValue)
				break
			}
		}
		fmt.Print(newValue, " ", table.Name)
		fmt.Println()
	}
	fmt.Println()

	return newValue
}
