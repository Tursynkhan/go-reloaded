package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]
	file1, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		str := (scanner.Text())
		fmt.Println(str)
	}
}
