package funcUp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Up(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if len(slice) == 1 {
			if slice[i] == "(up)" {
				slice[i] = ""
			}
		}
		if slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			continue
		}
		if slice[i] == "(up," {
			secondPart := slice[i+1]
			stringNumber := strings.Split(secondPart, ")")[0]
			n, err := strconv.Atoi(stringNumber)
			if err != nil {
				fmt.Println("Cannot convert number or inccorect usage of cap")
				os.Exit(0)
			}
			if n <= 0 {
				fmt.Println("Number cannot be negative or zero")
				os.Exit(0)
			}
			if n <= len(slice[:i]) {
				counter := 0
				for j := i - 1; counter != n; j-- {
					slice[j] = strings.ToUpper(slice[j])
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
