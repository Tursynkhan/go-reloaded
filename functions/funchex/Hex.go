package funchex

import (
	"strconv"
	"strings"
)

func Hex(str string) string {
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if len(slice) == 1 {
			if slice[i] == "(hex)" {
				slice[i] = ""
			}
		}
		if slice[i] == "(hex)" {
			decimal, _ := strconv.ParseInt(slice[i-1], 16, 64)

			slice[i-1] = strconv.Itoa(int(decimal))
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return strings.Join(slice, " ")
}
