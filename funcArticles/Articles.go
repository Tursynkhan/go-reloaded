package funcArticles

import "strings"

func Articles(str string) string {
	slice := strings.Fields(str)

	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[len(slice)-1] {
			break
		}
		if (slice[i] == "a" || slice[i] == "A") && (slice[i+1][0] == 'a' || slice[i+1][0] == 'e' || slice[i+1][0] == 'y' || slice[i+1][0] == 'u' || slice[i+1][0] == 'i' || slice[i+1][0] == 'o' || slice[i+1][0] == 'h') {
			slice[i] = slice[i] + "n"
		}
	}
	return strings.Join(slice, " ")
}
