package funcCap

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cap(str string) string {
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(cap)" {
			slice[i-1] = strings.Title(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(cap," {

			num := string(slice[i+1][0])
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
				slice[j] = strings.Title(slice[j])
				counter++
				i--
			}
			slice = append(slice[:i], slice[i+2:]...)

		}
	}

	return strings.Join(slice, " ")
}
