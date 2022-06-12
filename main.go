package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
}

func Up(str *string) {
	word := strings.Fields(*str)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
	}
	fmt.Println(word)
	*str = strings.Join(word, " ")
}
