package funcBin

import (
	"strconv"
	"strings"
)

func Bin(str string) string {
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(bin)" {
			decimal, _ := strconv.ParseInt(slice[i-1], 2, 64)

			slice[i-1] = strconv.Itoa(int(decimal))
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return strings.Join(slice, " ")
}
