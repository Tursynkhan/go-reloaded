package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
