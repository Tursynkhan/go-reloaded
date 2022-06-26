package main

import (
	"fmt"
	"strings"
)

func isValid(arg []string) bool {
	if len(arg) != 2 {
		fmt.Println("Check numbers of arguments")
		return false
	}

	fNames := []string{"main.go", "go.mod"}
	if isInStrings(arg[0], fNames) || isInStrings(arg[1], fNames) {
		fmt.Println("Name of the output file should be different from existing files")
		return false
	}

	if arg[0] != "sample.txt" && arg[1] != "result.txt" {
		fmt.Println("The first should be sample.txt AND the second argument must be result.txt")
		return false
	}

	if strings.HasSuffix(arg[1], ".go") {
		fmt.Println("Do not change go files")
		return false
	}

	return true
}

func isInStrings(s string, strs []string) bool {
	for i := range strs {
		if strs[i] == s {
			return true
		}
	}
	return false
}

func isInRunes(r rune, runes []rune) bool {
	for i := range runes {
		if runes[i] == r {
			return true
		}
	}
	return false
}
