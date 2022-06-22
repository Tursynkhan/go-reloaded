package main

import (
	"go-reloaded/Functions/funcArticles"
	"go-reloaded/Functions/funcBin"
	"go-reloaded/Functions/funcCap"
	"go-reloaded/Functions/funcHex"
	"go-reloaded/Functions/funcLow"
	"go-reloaded/Functions/funcPunct"
	"go-reloaded/Functions/funcQuotes"
	"go-reloaded/Functions/funcUp"
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
			text = funcUp.Up(text)
		case "(up,":
			text = funcUp.Up(text)
		case "(low)":
			text = funcLow.Low(text)
		case "(low,":
			text = funcLow.Low(text)
		case "(cap)":
			text = funcCap.Cap(text)
		case "(cap,":
			text = funcCap.Cap(text)
		case "(hex)":
			text = funcHex.Hex(text)
		case "(bin)":
			text = funcBin.Bin(text)
		}

		if isInStrings(words[i], symbols) && isInRunes(rune(words[i][0]), punctuations) {
			text = funcPunct.Punct(text)
		} else if (words[i] == "a" || words[i] == "A") && isInRunes(rune(words[i+1][0]), vowels) {
			text = funcArticles.Articles(text)
		} else if words[i] == "'" {
			text = funcQuotes.SingleQuotes(text)
			text = funcQuotes.DoubleQuotes(text)
		}
	}

	return text
}
