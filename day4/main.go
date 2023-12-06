package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 4 Part 1")
	input := utils.ReadFileAsLines("day4/input.txt")
	cards := generateCards(&input)
	fmt.Println(totalScore(&cards))
}

func Part2() {
	fmt.Println("Day 4 Part 2")
	input := utils.ReadFileAsLines("day4/input.txt")
	cards := generateCards(&input)
	fmt.Println(totalNumberOfWins(&cards))
}

type card struct {
	id             int
	winningNumbers map[int]bool
	numbers        []int
}

func (c card) calculateScore() int {
	score := 0
	for _, num := range c.numbers {
		if c.winningNumbers[num] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func (c card) calculateNumberOfWins() int {
	numberOfWins := 0
	for _, num := range c.numbers {
		if c.winningNumbers[num] {
			numberOfWins += 1
		}
	}
	return numberOfWins
}

func generateCards(input *[]string) []card {
	cards := make([]card, 0)
	for _, line := range *input {
		state := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Trim(strings.ReplaceAll(state[0], "Card ", ""), " "))
		gameStates := strings.Split(state[1], " | ")
		winningStates := strings.Split(gameStates[0], " ")
		winningNumbers := make(map[int]bool)
		for _, winningState := range winningStates {
			num, err := strconv.Atoi(winningState)
			if err != nil {
				continue
			}
			winningNumbers[num] = true
		}
		numberStates := strings.Split(gameStates[1], " ")
		numbers := make([]int, 0)
		for _, numberState := range numberStates {
			num, err := strconv.Atoi(numberState)
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}
		cards = append(cards, card{id: id, winningNumbers: winningNumbers, numbers: numbers})
	}
	return cards
}

func totalScore(cards *[]card) int {
	sum := 0
	for _, card := range *cards {
		sum += card.calculateScore()
	}
	return sum
}

func totalNumberOfWins(cards *[]card) int {
	cardsMap := make(map[int]int)
	for _, card := range *cards {
		cardsMap[card.id] = 1
	}
	for _, card := range *cards {
		wins := card.calculateNumberOfWins() * cardsMap[card.id]
		for i := 0; i < wins; i++ {
			cardsMap[card.id+1+(i/cardsMap[card.id])] += 1
		}
	}
	sum := 0
	for _, v := range cardsMap {
		sum += v
	}
	return sum
}
