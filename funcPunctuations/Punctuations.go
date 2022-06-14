package funcPunctuations

import (
	"fmt"
	"strings"
)

func Punctuations(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "," || slice[i] == "." || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" {
			slice[i-1] = slice[i-1] + slice[i]
			slice[i] = ""
			continue
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	fmt.Println(slice)
	return strings.Join(slice, " ")
}
