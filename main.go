package main

import (
	"bufio"
	"fmt"
	"go-reloaded/Functions/funcArticles"
	"go-reloaded/Functions/funcBin"
	"go-reloaded/Functions/funcCap"
	"go-reloaded/Functions/funcHex"
	"go-reloaded/Functions/funcLow"
	"go-reloaded/Functions/funcPunct"
	"go-reloaded/Functions/funcQuotes"
	"go-reloaded/Functions/funcUp"
	"log"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Check numbers of arguments")
		return
	}
	if arg[0] == "main.go" || arg[0] == "go.mod" || arg[1] == "main.go" || arg[1] == "go.mod" {
		fmt.Println("Name of the output file should be different from existing files")
		return
	}
	if arg[0] != "sample.txt" {
		fmt.Println("The argument should be sample.txt")
		os.Exit(0)
	}
	if arg[1] != "result.txt" {
		fmt.Println("The argument should be result.txt")
		os.Exit(0)

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
	for i, w := range strSlice {
		if i >= 1 && i <= len(strSlice)-1 {
			str = str + "\n"
		}
		str = str + w
	}

	slice := strings.Split(str, " ")

	for i := 0; i < len(slice); i++ {
		switch slice[i] {
		case "(up)":
			str = funcUp.Up(str)
		case "(up,":
			str = funcUp.Up(str)
		case "(low)":
			str = funcLow.Low(str)
		case "(low,":
			str = funcLow.Low(str)
		case "(cap)":
			str = funcCap.Cap(str)
		case "(cap,":
			str = funcCap.Cap(str)
		case "(hex)":
			str = funcHex.Hex(str)
		case "(bin)":
			str = funcBin.Bin(str)
		}
		if slice[i] == "," || slice[i] == "." || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" || slice[i] == "..." || slice[i] == "?!" {
			str = funcPunct.Punct(str)
		}
		if (slice[i] == "a" || slice[i] == "A") && (slice[i+1][0] == 'a' || slice[i+1][0] == 'e' || slice[i+1][0] == 'y' || slice[i+1][0] == 'u' || slice[i+1][0] == 'i' || slice[i+1][0] == 'o' || slice[i+1][0] == 'h') {
			str = funcArticles.Articles(str)
		}
		if slice[i] == "'" {
			str = funcQuotes.SingleQuotes(str)
			str = funcQuotes.DoubleQuotes(str)
		}

	}
	file2, err := os.Create(arg[1])
	if err != nil {
		log.Fatal(err)
	}
	l, err := file2.WriteString(string(str))
	if err != nil {
		log.Fatal(err)
	}
	if l == 0 {
		return
	}
	defer file2.Close()
}
