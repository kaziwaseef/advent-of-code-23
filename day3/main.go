package day3

import (
	"fmt"
	"strconv"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 3 Part 1")
	input := utils.ReadFileAsLines("day3/input.txt")
	fmt.Println(alg(&input, ""))
}

func Part2() {
	fmt.Println("Day 3 Part 2")
	input := utils.ReadFileAsLines("day3/input.txt")
	fmt.Println(alg(&input, "gearRatio"))
}

type gearStruct struct {
	i     int
	j     int
	left  int
	right int
}

func (g gearStruct) ratio() int {
	return g.left * g.right
}

func alg(input *[]string, constraint string) int {
	gears := make([]gearStruct, 0)
	numChar := ""
	numLen := 0
	sum := 0
	for i, line := range *input {
		numChar = ""
		numLen = 0
		for j := 0; j < len(line); j++ {
			n, err := strconv.Atoi(string(line[j]))
			if err != nil {
				if len(numChar) > 0 {
					num, err := strconv.Atoi(numChar)
					utils.CheckErr(err)
					aroundIndices := getAroundIndices(i, len(*input), j, len(line), num, numLen)
					if constraint == "gearRatio" {
						updateIfGearRatio(num, input, &aroundIndices, &gears)
					} else if isValidNumber(input, &aroundIndices) {
						sum += num
					}
				}
				numChar = ""
				numLen = 0
				continue
			}
			numChar += strconv.Itoa(n)
			numLen += 1
		}

		if len(numChar) > 0 {
			num, err := strconv.Atoi(numChar)
			utils.CheckErr(err)
			horizontal := len((*input)[0]) - 1
			aroundIndices := getAroundIndices(i, len(*input), horizontal, horizontal+1, num, numLen)
			if constraint == "gearRatio" {
				updateIfGearRatio(num, input, &aroundIndices, &gears)
			} else if isValidNumber(input, &aroundIndices) {
				sum += num
			}
		}
	}

	if constraint == "gearRatio" {
		for _, gear := range gears {
			sum += gear.ratio()
		}
	}

	return sum
}

func getAroundIndices(i, iLen, j, jLen, num, numLen int) [][]int {
	indices := make([][]int, 0)
	var left, right []int
	if (j - numLen) > 0 {
		left = []int{i, j - numLen - 1}
		indices = append(indices, left)
	}
	if j < jLen {
		right = []int{i, j}
		indices = append(indices, right)
	}

	if i > 0 {
		if left != nil {
			indices = append(indices, []int{i - 1, j - numLen - 1})
		}
		if right != nil {
			indices = append(indices, []int{i - 1, j})
		}
		for k := 0; k < numLen; k++ {
			indices = append(indices, []int{i - 1, j - numLen + k})
		}
	}

	if (i + 1) < iLen {
		if left != nil {
			indices = append(indices, []int{i + 1, j - numLen - 1})
		}
		if right != nil {
			indices = append(indices, []int{i + 1, j})
		}
		for k := 0; k < numLen; k++ {
			indices = append(indices, []int{i + 1, j - numLen + k})
		}
	}

	return indices
}

func isValidNumber(input *[]string, aroundIndices *[][]int) bool {
	invalidCharacterMap := map[string]int{
		"1": 1,
		"2": 1,
		"3": 1,
		"4": 1,
		"5": 1,
		"6": 1,
		"7": 1,
		"8": 1,
		"9": 1,
		"0": 1,
		".": 1,
	}
	for _, indices := range *aroundIndices {
		char := string((*input)[indices[0]][indices[1]])
		if _, isInvalid := invalidCharacterMap[char]; !isInvalid {
			return true
		}
	}
	return false
}

func updateIfGearRatio(num int, input *[]string, aroundIndices *[][]int, gears *[]gearStruct) {
	for _, indices := range *aroundIndices {
		char := string((*input)[indices[0]][indices[1]])
		if char == "*" {
			matched := false
			for i, gear := range *gears {
				if gear.i == indices[0] && gear.j == indices[1] {
					gear.right = num
					(*gears)[i] = gear
					matched = true
					break
				}
			}
			if !matched {
				*gears = append(*gears, gearStruct{i: indices[0], j: indices[1], left: num})
			}
		}
	}
}
