package utils

import (
	"os"
	"strings"
)

func ReadFileAsLines(fileName string) []string {
	rawInput, err := os.ReadFile(fileName)
	CheckErr(err)
	input := strings.Split(strings.Trim(string(rawInput), "\n"), "\n")
	return input
}
