package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 2 Part 1")
	rawInput, err := os.ReadFile("day2/input.txt")
	utils.CheckErr(err)
	input := strings.Split(string(rawInput), "\n")
	gameInput := formatInput(&input)
	config := cubes{red: 12, green: 13, blue: 14}
	isValidGame(&gameInput, &config)
}

func Part2() {
	fmt.Println("Day 2 Part 2")
	rawInput, err := os.ReadFile("day2/input.txt")
	utils.CheckErr(err)
	input := strings.Split(string(rawInput), "\n")
	gameInput := formatInput(&input)
	sumOfPower(&gameInput)
}

func isValidGame(input *[]gameInput, config *cubes) {
	idSum := 0
	for _, game := range *input {
		isValidGame := true
		for _, round := range game.rounds {
			if round.red > config.red || round.green > config.green || round.blue > config.blue {
				isValidGame = false
				break
			}
		}
		if isValidGame {
			idSum += game.id
		}
	}
	fmt.Println(idSum)
}

func sumOfPower(input *[]gameInput) {
	sum := 0
	for _, game := range *input {
		minCubes := cubes{}
		for _, round := range game.rounds {
			minCubes.red = int(math.Max(float64(minCubes.red), float64(round.red)))
			minCubes.green = int(math.Max(float64(minCubes.green), float64(round.green)))
			minCubes.blue = int(math.Max(float64(minCubes.blue), float64(round.blue)))
		}
		sum += minCubes.power()
	}
	fmt.Println(sum)
}

type gameInput struct {
	id     int
	rounds []cubes
}

type cubes struct {
	red   int
	green int
	blue  int
}

func (c cubes) power() int {
	return c.red * c.green * c.blue
}

func formatInput(input *[]string) []gameInput {
	gameInputs := make([]gameInput, 0)
	for _, line := range *input {
		info := strings.Split(line, ": ")
		if len(info) != 2 {
			break
		}
		id, err := strconv.Atoi(strings.ReplaceAll(info[0], "Game ", ""))
		utils.CheckErr(err)
		gameInp := gameInput{id: id}
		gameStr := strings.Split(info[1], "; ")
		for _, round := range gameStr {
			roundInp := cubes{}
			roundInfo := strings.Split(round, ", ")
			for _, cube := range roundInfo {
				if strings.Contains(cube, "red") {
					redCount, err := strconv.Atoi(strings.ReplaceAll(cube, " red", ""))
					utils.CheckErr(err)
					roundInp.red = redCount
				}
				if strings.Contains(cube, "green") {
					greenCount, err := strconv.Atoi(strings.ReplaceAll(cube, " green", ""))
					utils.CheckErr(err)
					roundInp.green = greenCount
				}
				if strings.Contains(cube, "blue") {
					blueCount, err := strconv.Atoi(strings.ReplaceAll(cube, " blue", ""))
					utils.CheckErr(err)
					roundInp.blue = blueCount
				}
			}
			gameInp.rounds = append(gameInp.rounds, roundInp)
		}
		gameInputs = append(gameInputs, gameInp)
	}
	return gameInputs
}
