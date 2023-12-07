package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 7 Part 1")
	input := utils.ReadFileAsLines("day7/input.txt")
	hands := formatInput(&input)
	slices.SortFunc(hands, comparator)
	fmt.Println(calculateScore(&hands))
}

func calculateScore(hands *[]hand) int {
	score := 0

	for i, hand := range *hands {
		score += hand.getScore(i + 1)
	}
	return score
}

func comparator(left hand, right hand) int {
	return left.compare(right)
}

func formatInput(input *[]string) []hand {
	hands := make([]hand, 0)
	for _, line := range *input {
		handInput := strings.Split(line, " ")
		bid, err := strconv.Atoi(handInput[1])
		utils.CheckErr(err)
		hands = append(hands, hand{cards: handInput[0], bid: bid})
	}
	return hands
}

type hand struct {
	cards string
	bid   int
	power int
}

func (h hand) getScore(rank int) int {
	return h.bid * rank
}

func (h *hand) getPower() int {
	if h.power > 0 {
		return h.power
	}
	return h.calculatePower()
}

func (h *hand) calculatePower() int {
	cardMap := make(map[string]int)
	for _, card := range h.cards {
		cardStr := string(card)
		if v, exists := cardMap[cardStr]; exists {
			cardMap[cardStr] = v + 1
			continue
		}
		cardMap[cardStr] = 1
	}
	cardMapLength := len(cardMap)
	if cardMapLength == 1 {
		// 5 of a kind
		return 7
	}
	if cardMapLength == 2 {
		max := 0
		for _, num := range cardMap {
			if num > max {
				max = num
			}
		}
		if max == 4 {
			// 4 of a kind
			return 6
		}
		// Full house
		return 5
	}
	if cardMapLength == 3 {
		max := 0
		for _, num := range cardMap {
			if num > max {
				max = num
			}
		}
		if max == 3 {
			// 3 of a kind
			return 4
		}
		// 2 pair
		return 3
	}
	if cardMapLength == 4 {
		// 1 pair
		return 2
	}
	return 1
}

func (left hand) compare(right hand) int {
	// left < right = -1
	// left > right = 1
	// left == right = 0
	if left.cards == right.cards {
		return 0
	}
	if left.getPower() > right.getPower() {
		return 1
	}
	if left.getPower() < right.getPower() {
		return -1
	}
	return left.compareCards(right)
}

var cardOrder = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

func (left hand) compareCards(right hand) int {
	// left < right = -1
	// left > right = 1
	// left == right = 0
	for i := 0; i < len(left.cards); i++ {
		leftOrder := cardOrder[string(left.cards[i])]
		rightOrder := cardOrder[string(right.cards[i])]
		if leftOrder > rightOrder {
			return 1
		}
		if leftOrder < rightOrder {
			return -1
		}
	}

	return 0
}
