package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaziwaseef/advent-of-code-23/utils"
)

func Part1() {
	fmt.Println("Day 5 Part 1")
	input := utils.ReadFileAsLines("day5/input.txt")
	seedsString := input[0]
	seedInput := makeSeedInput(seedsString)
	mapInput := input[2:]
	alg(&mapInput, &seedInput)
}

func Part2() {
	fmt.Println("Day 5 Part 2")
	input := utils.ReadFileAsLines("day5/input.txt")
	seedsString := input[0]
	seedInput := makeRangeSeedInput(seedsString)
	fmt.Println(len(seedInput))
	mapInput := input[2:]
	alg(&mapInput, &seedInput)
}

func alg(mapInput *[]string, seedInput *[]int) {
	sdm := generateSDM(mapInput)
	values := getValues(*seedInput, "seed", "location", sdm)
	minValue := getMinValueIndex(&values)
	fmt.Println(minValue)
}

func makeSeedInput(seedString string) []int {
	seedString = strings.ReplaceAll(seedString, "seeds: ", "")
	seedStrings := strings.Split(seedString, " ")
	output := make([]int, 0)
	for _, str := range seedStrings {
		num, err := strconv.Atoi(str)
		utils.CheckErr(err)
		output = append(output, num)
	}
	return output
}

func makeRangeSeedInput(seedString string) []int {
	seedString = strings.ReplaceAll(seedString, "seeds: ", "")
	seedStrings := strings.Split(seedString, " ")
	output := make([]int, 0)
	for i := 0; i < len(seedStrings); i += 2 {
		num, err := strconv.Atoi(seedStrings[i])
		utils.CheckErr(err)
		r, err := strconv.Atoi(seedStrings[i+1])
		utils.CheckErr(err)
		for j := 0; j < r; j++ {
			output = append(output, num+j)
		}
	}
	return output
}

func getMinValueIndex(values *[]int) int {
	minValue := (*values)[0]
	for _, value := range *values {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func generateSDM(input *[]string) *SourceDestinationMappings {
	rootSdm := new(SourceDestinationMappings)
	currentSdm := rootSdm
	mapStrings := strings.Split(strings.Join(*input, "\n"), "\n\n")
	r := regexp.MustCompile(`(?s)(\w*)-to-(\w*)`)
	for _, mapString := range mapStrings {
		res := r.FindAllStringSubmatch(mapString, -1)
		source := res[0][1]
		destination := res[0][2]
		currentSdm.source = source
		currentSdm.destination = destination
		strs := strings.Split(mapString, "\n")[1:]
		for _, str := range strs {
			numChars := strings.Split(str, " ")
			destinationStart, _ := strconv.Atoi(numChars[0])
			sourceStart, _ := strconv.Atoi(numChars[1])
			sourceRange, _ := strconv.Atoi(numChars[2])
			sourceEnd := sourceStart + sourceRange - 1
			currMapping := mapping{sourceStart: sourceStart, sourceEnd: sourceEnd, destinationStart: destinationStart}
			currentSdm.mappings = append(currentSdm.mappings, currMapping)
		}
		newSdm := new(SourceDestinationMappings)
		currentSdm.destinationMap = newSdm
		currentSdm = newSdm
	}

	return rootSdm
}

func getValues(input []int, source string, destination string, sdm *SourceDestinationMappings) []int {
	if source != sdm.source {
		panic("Invalid Starting SDM")
	}
	results := make([]int, 0)
	for _, num := range input {
		results = append(results, sdm.getDestination(num))
	}
	if sdm.destination != destination {
		return getValues(results, sdm.destination, destination, sdm.destinationMap)
	}
	return results
}

type SourceDestinationMappings struct {
	source         string
	destination    string
	destinationMap *SourceDestinationMappings
	mappings       []mapping
}

func (sdm SourceDestinationMappings) getDestination(source int) int {
	for _, eachMapping := range sdm.mappings {
		if eachMapping.isInMapping(source) {
			return eachMapping.getDestination(source)
		}
	}
	return source
}

type mapping struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
}

func (m mapping) isInMapping(source int) bool {
	return source >= m.sourceStart && source <= m.sourceEnd
}

func (m mapping) getDestination(source int) int {
	return source - m.sourceStart + m.destinationStart
}
