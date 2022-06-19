package funcUp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Up(str string) string {
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(up," {
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
				slice[j] = strings.ToUpper(slice[j])
				counter++
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		}
	}
	return strings.Join(slice, " ")
}
