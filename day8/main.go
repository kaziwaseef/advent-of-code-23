package day8

import (
	"fmt"
	"regexp"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 8 Part 1")
	input := utils.ReadFileAsLines("day8/input.txt")
	instructions, nodeMap := formatInput(&input)
	fmt.Println(stepsTo(instructions, nodeMap, "AAA", "ZZZ", 0))
}

func stepsTo(instructions string, nodeMap map[string]Node, start string, end string, count int) int {
	for _, ins := range instructions {
		if start == end {
			return count
		}
		node, exists := nodeMap[start]
		if !exists {
			fmt.Println("Node", start)
			panic("Node does not exist")
		}
		count += 1
		if ins == 'L' {
			start = node.left
			continue
		}
		start = node.right
	}

	return stepsTo(instructions, nodeMap, start, end, count)
}

func formatInput(input *[]string) (string, map[string]Node) {
	instruction := (*input)[0]

	nodeMap := make(map[string]Node)

	for i := 2; i < len(*input); i++ {
		line := (*input)[i]
		regex := regexp.MustCompile(`(\w{3})`)
		matches := regex.FindAllStringSubmatch(line, -1)
		nodeMap[matches[0][0]] = Node{left: matches[1][0], right: matches[2][0]}
	}

	return instruction, nodeMap
}

type Node struct {
	left  string
	right string
}
