package day6

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 6 Part 1")
	input := utils.ReadFileAsLines("day6/input.txt")
	races := formatInput(&input)
	fmt.Println(alg(&races))
}

func alg(races *[]race) int {
	product := 1
	for _, race := range *races {
		product *= race.numberOfWinsPossible()
	}
	return product
}

func formatInput(input *[]string) []race {
	races := make([]race, 0)
	regex := regexp.MustCompile(`(\d+)`)
	timeMatches := regex.FindAllStringSubmatch((*input)[0], -1)
	recordMatches := regex.FindAllStringSubmatch((*input)[1], -1)
	for i := 0; i < len(timeMatches); i++ {
		time, _ := strconv.Atoi(timeMatches[i][0])
		record, _ := strconv.Atoi(recordMatches[i][0])
		races = append(races, race{time: time, record: record})
	}
	return races
}

type race struct {
	time   int
	record int
}

func (r race) numberOfWinsPossible() int {
	determinant := (math.Pow(float64(r.time), 2)) - float64(4*r.record)
	determinant = math.Sqrt(determinant)

	firstRoot := math.Floor((float64(r.time) - determinant) / 2)
	secondRoot := math.Ceil((float64(r.time) + determinant) / 2)

	return int(secondRoot - firstRoot - 1)
}
