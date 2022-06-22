package main

import (
	"log"
	"os"
)

func main() {
	// too many logic on main function. It must just run process or/and handle data and send to other secondary func-s.
	args := os.Args[1:]

	if !isValid(args) {
		log.Fatalln("wrong input argument")
	}

	text, err := readSample(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	text = convertText(text)

	err = writeResult(args[1], text)
	if err != nil {
		log.Fatalln(err)
	}
}
