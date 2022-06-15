package funcPunct

import (
	"strings"
)

func Punct(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "," || slice[i] == "." || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" {
			slice[i-1] = slice[i-1] + slice[i]
			slice[i] = ""
			continue
		}
		if (strings.HasPrefix(slice[i], ".") || strings.HasPrefix(slice[i], ",") || strings.HasPrefix(slice[i], "!") || strings.HasPrefix(slice[i], "?") ||
			strings.HasPrefix(slice[i], ":") || strings.HasPrefix(slice[1], ";")) && (slice[i][1] == 46 || slice[i][1] == 44 || slice[i][1] == 33 || slice[i][1] == 63 || slice[i][1] == 58 ||
			slice[i][1] == 59) {
			slice[i-1] = slice[i-1] + slice[i]
			slice[i] = ""
			continue
		}
		if (strings.HasPrefix(slice[i], ".") || strings.HasPrefix(slice[i], ",") || strings.HasPrefix(slice[i], "!") || strings.HasPrefix(slice[i], "?") ||
			strings.HasPrefix(slice[i], ":") || strings.HasPrefix(slice[1], ";")) && !(slice[i][1] == 46 || slice[i][1] == 44 || slice[i][1] == 33 || slice[i][1] == 63 || slice[i][1] == 58 ||
			slice[i][1] == 59) {
			slice[i-1] = slice[i-1] + string(slice[i][0])
			slice[i] = slice[i][1:]
			continue
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return strings.Join(slice, " ")
}
