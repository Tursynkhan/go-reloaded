package main

import (
	"go-reloaded/functions/funcarticles"
	"go-reloaded/functions/funcbin"
	"go-reloaded/functions/funccap"
	"go-reloaded/functions/funchex"
	"go-reloaded/functions/funclow"
	"go-reloaded/functions/funcpunct"
	"go-reloaded/functions/funcquotes"
	"go-reloaded/functions/funcup"
	"strings"
)

func convertText(text string) string {
	words := strings.Split(text, " ")
	symbols := []string{",", ".", "!", "?", ":", ";", "...", "?!"}
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}
	vowels := []rune{'a', 'e', 'y', 'u', 'i', 'o', 'h'}
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(up)":
			text = funcup.Up(text)
		case "(up,":
			text = funcup.Up(text)
		case "(low)":
			text = funclow.Low(text)
		case "(low,":
			text = funclow.Low(text)
		case "(cap)":
			text = funccap.Cap(text)
		case "(cap,":
			text = funccap.Cap(text)
		case "(hex)":
			text = funchex.Hex(text)
		case "(bin)":
			text = funcbin.Bin(text)
		}

		if isInStrings(words[i], symbols) && isInRunes(rune(words[i][0]), punctuations) {
			text = funcpunct.Punct(text)
		} else if (words[i] == "a" || words[i] == "A") && isInRunes(rune(words[i+1][0]), vowels) {
			text = funcarticles.Articles(text)
		} else if words[i] == "'" {
			text = funcquotes.SingleQuotes(text)
			text = funcquotes.DoubleQuotes(text)
		}
	}

	return text
}
