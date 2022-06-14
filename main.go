package main

import (
	"bufio"
	"fmt"
	"go-reloaded/funcBin"
	"go-reloaded/funcCap"
	"go-reloaded/funcHex"
	"go-reloaded/funcLow"
	"go-reloaded/funcPunctuations"
	"go-reloaded/funcUp"
	"log"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Chech numbers of arguments")
		return
	}
	if arg[0] != "sample.txt" {
		fmt.Println("the name of the input file that should be sample.txt")
		return
	}
	if arg[1] != "result.txt" {
		fmt.Println("the name of the input file that should be result.txt")
		return
	}

	file1, err := os.Open(arg[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	var str string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		str += scanner.Text()
	}

	str = funcUp.Up(str)
	str = funcLow.Low(str)
	str = funcCap.Cap(str)
	str = funcHex.Hex(str)
	str = funcBin.Bin(str)
	str = funcPunctuations.Punctuations(str)

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
