package funccap

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cap(str string) string {
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if len(slice) == 1 {
			if slice[i] == "(cap)" {
				slice[i] = ""
			}
		}
		if slice[i] == "(cap)" {
			slice[i-1] = strings.Title(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			continue
		}
		if slice[i] == "(cap," {
			secondPart := slice[i+1]
			stringNumber := strings.Split(secondPart, ")")[0]
			n, err := strconv.Atoi(stringNumber)
			if err != nil {
				fmt.Println("Cannot convert number")
				os.Exit(0)
			}
			if n <= 0 {
				slice = append(slice[:i], slice[i+2:]...)
				break
			}
			if n <= len(slice[:i]) {
				counter := 0
				for j := i - 1; counter != n; j-- {
					slice[j] = strings.Title(slice[j])
					counter++
				}
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				fmt.Println("Wrond numbers to Upper, check it after word: " + slice[i-1])
				os.Exit(0)
			}
		}
	}

	return strings.Join(slice, " ")
}
