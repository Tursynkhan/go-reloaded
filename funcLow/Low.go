package funcLow

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Low(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(low," {
			num := slice[i+1][:len(slice[i+1])-1]
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Cannot convert number")
				os.Exit(0)
			}
			if n <= 0 {
				fmt.Println("Number cannot be negative or zero")
				os.Exit(0)
			}
			counter := 0
			for j := i - 1; counter != n; j-- {
				slice[j] = strings.ToLower(slice[j])
				counter++
			}
			slice = append(slice[:i], slice[i+2:]...)
		}
	}
	return strings.Join(slice, " ")
}
