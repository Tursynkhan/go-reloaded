package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1]
	file1, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	var str string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		str += scanner.Text()
	}

	Up(&str)
	Low(&str)
	Cap(&str)
	Hex(&str)
	Bin(&str)
}

func Up(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
	}
	*str = strings.Join(word, " ")
}

func Low(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(low)" {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
	}
	*str = strings.Join(word, " ")
}

func Cap(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(cap)" {
			word[i-1] = strings.Title(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
	}
	*str = strings.Join(word, " ")
}

func Hex(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" {
			decimal, _ := strconv.ParseInt(word[i-1], 16, 64)

			word[i-1] = strconv.Itoa(int(decimal))
			word = append(word[:i], word[i+1:]...)
		}
	}
	*str = strings.Join(word, " ")
}

func Bin(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(bin)" {
			decimal, _ := strconv.ParseInt(word[i-1], 2, 64)

			word[i-1] = strconv.Itoa(int(decimal))
			word = append(word[:i], word[i+1:]...)
		}
	}
	*str = strings.Join(word, " ")
	fmt.Println(word)
}
