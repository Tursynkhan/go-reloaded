package main

import (
	"bufio"
	"os"
	"strings"
)

func readSample(fName string) (string, error) {
	file1, err := os.Open(fName)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	defer file1.Close()

	var str string
	var strSlice []string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		strSlice = append(strSlice, scanner.Text())
	}

	str = strings.Join(strSlice, "\n")

	return str, nil
}

func writeResult(fName, text string) error {
	file2, err := os.Create(fName)
	if err != nil {
		return err
	}
	defer file2.Close()

	_, err = file2.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}
