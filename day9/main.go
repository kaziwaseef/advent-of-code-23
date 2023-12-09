package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 9 Part 1")
	input := utils.ReadFileAsLines("day9/input.txt")
	fmt.Println(alg(&input, "future"))
}

func Part2() {
	fmt.Println("Day 9 Part 2")
	input := utils.ReadFileAsLines("day9/input.txt")
	fmt.Println(alg(&input, "past"))
}

func alg(input *[]string, direction string) int {
	value := 0
	for _, line := range *input {
		numStrs := strings.Split(line, " ")
		values := make([]int, 0)
		for _, num := range numStrs {
			value, err := strconv.Atoi(num)
			utils.CheckErr(err)
			values = append(values, value)
		}
		var initialFinalValues []int
		if direction == "future" {
			initialFinalValues = []int{values[len(values)-1]}
		} else {
			initialFinalValues = []int{values[0]}
		}
		sequence := Sequence{values: values, finalValues: initialFinalValues}
		sequence.resolvefinalValues(direction)
		sequence.resolvePrediction(direction)
		value += sequence.prediction
	}
	return value
}

type Sequence struct {
	finalValues []int
	values      []int
	prediction  int
}

func (s *Sequence) resolvefinalValues(direction string) {
	diff := make([]int, 0)
	finished := true
	for i := 1; i < len(s.values); i++ {
		thisDiff := s.values[i] - s.values[i-1]
		diff = append(diff, thisDiff)
		if thisDiff != 0 {
			finished = false
		}
	}
	finalValue := 0
	if direction == "future" {
		finalValue = diff[len(diff)-1]
	} else {
		finalValue = diff[0]
	}
	s.finalValues = append(s.finalValues, finalValue)
	if finished {
		return
	}
	s.values = diff
	s.resolvefinalValues(direction)
}

func (s *Sequence) resolvePrediction(direction string) int {
	prediction := 0
	for i := len(s.finalValues) - 1; i >= 0; i-- {
		if direction == "future" {
			prediction += s.finalValues[i]
		} else {
			prediction = s.finalValues[i] - prediction
		}
	}
	s.prediction = prediction
	return prediction
}
