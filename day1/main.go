package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 1 Part 1")
	rawInput, err := os.ReadFile("day1/input.txt")
	utils.CheckErr(err)
	input := strings.Split(string(rawInput), "\n")
	alg(&input, nil)
}

func Part2() {
	fmt.Println("Day 1 Part 2")
	replacementMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	rawInput, err := os.ReadFile("day1/input.txt")
	utils.CheckErr(err)
	input := strings.Split(string(rawInput), "\n")
	alg(&input, &replacementMap)
}

func alg(input *[]string, replacementMap *map[string]int) {
	var sum int
	for _, line := range *input {
		chars := strings.Split(line, "")
		var first, last int
		for i := 0; i < len(chars); i++ {
			num, err := strconv.Atoi(chars[i])
			if err != nil {
				if replacementMap != nil {
					// fmt.Println("Forward")
					matched := false
					for k, v := range *replacementMap {
						if i+len(k) > len(chars) {
							continue
						}
						lookup := strings.Join(chars[i:i+len(k)], "")
						// fmt.Println(lookup, k)
						if k == lookup {
							// fmt.Println("Matched", k, lookup, v)
							first = v
							matched = true
							break
						}
					}
					if matched {
						break
					}
				}
				continue
			}
			first = num
			break
		}
		for i := len(chars) - 1; i >= 0; i-- {
			num, err := strconv.Atoi(chars[i])
			if err != nil {
				if replacementMap != nil {
					// fmt.Println("Reverse")
					matched := false
					for k, v := range *replacementMap {
						if i-len(k)+1 < 0 {
							continue
						}
						lookup := strings.Join(chars[i-len(k)+1:i+1], "")
						// fmt.Println(lookup, k)
						if k == lookup {
							// fmt.Println("Matched", k, lookup, v)
							last = v
							matched = true
							break
						}
					}
					if matched {
						break
					}
				}
				continue
			}
			last = num
			break
		}
		sum += (first * 10) + last
	}

	fmt.Println(sum)
}
