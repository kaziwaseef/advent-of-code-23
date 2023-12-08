package day8

import (
	"fmt"
	"regexp"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 8 Part 1")
	input := utils.ReadFileAsLines("day8/input.txt")
	instructions, nodeMap, _ := formatInput(&input, "AAA")
	fmt.Println(stepsTo(instructions, nodeMap, "AAA", "ZZZ", 0))
}

func Part2() {
	fmt.Println("Day 8 Part 2")
	input := utils.ReadFileAsLines("day8/input.txt")
	instructions, nodeMap, startArr := formatInput(&input, "..A")
	results := make([]int, 0)
	for _, start := range startArr {
		results = append(results, stepsTo(instructions, nodeMap, start, "..Z", 0))
	}
	fmt.Println(LCM(results...))
}

func stepsTo(instructions string, nodeMap map[string]Node, start string, end string, count int) int {
	endRegex := regexp.MustCompile(end)
	for _, ins := range instructions {
		if endRegex.Match([]byte(start)) {
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

func formatInput(input *[]string, start string) (string, map[string]Node, []string) {

	startArr := make([]string, 0)
	startRegex := regexp.MustCompile(start)

	instruction := (*input)[0]

	nodeMap := make(map[string]Node)

	for i := 2; i < len(*input); i++ {
		line := (*input)[i]
		regex := regexp.MustCompile(`(\w{3})`)
		matches := regex.FindAllStringSubmatch(line, -1)
		nodeMap[matches[0][0]] = Node{left: matches[1][0], right: matches[2][0]}
		if startRegex.Match([]byte(matches[0][0])) {
			startArr = append(startArr, matches[0][0])
		}
	}

	return instruction, nodeMap, startArr
}

type Node struct {
	left  string
	right string
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
