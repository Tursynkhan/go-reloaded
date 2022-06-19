package main

import (
	"bufio"
	"fmt"
	"go-reloaded/funcArticles"
	"go-reloaded/funcBin"
	"go-reloaded/funcCap"
	"go-reloaded/funcHex"
	"go-reloaded/funcLow"
	"go-reloaded/funcPunct"
	"go-reloaded/funcQuotes"
	"go-reloaded/funcUp"
	"log"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Chech numbers of arguments")
		return
	}

	if arg[1] == "main.go" || arg[1] == "go.mod" {
		fmt.Println("Name of the output file should be different from existing files")
		return
	}
	if strings.HasPrefix(arg[1], ".go") {
		fmt.Println("Do not change go files")
		return
	}

	file1, err := os.Open(arg[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	var str string
	var strSlice []string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		strSlice = append(strSlice, scanner.Text())
	}
	for i, w := range strSlice {
		if i >= 1 && i <= len(strSlice)-1 {
			str = str + "\n"
		}
		str = str + w
	}
	if strings.HasPrefix(str, "(bin") || strings.HasPrefix(str, "(hex") || strings.HasPrefix(str, "(cap") || strings.HasPrefix(str, "(low") || strings.HasPrefix(str, "(up") {
		fmt.Println("Don't play with sample.txt")
		return
	}
	str = strings.Trim(str, " ")
	str = funcPunct.Punct(str)
	str = funcQuotes.SingleQuotes(str)
	str = funcQuotes.DoubleQuotes(str)
	str = funcArticles.Articles(str)
	str = funcHex.Hex(str)
	str = funcBin.Bin(str)
	str = funcUp.Up(str)
	str = funcLow.Low(str)
	str = funcCap.Cap(str)

	file2, err := os.Create(arg[1])
	if err != nil {
		log.Fatal(err)
	}
	l, err := file2.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	if l == 0 {
		return
	}
	defer file2.Close()
}
