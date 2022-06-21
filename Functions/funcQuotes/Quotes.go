package funcQuotes

import (
	"strings"
)

func SingleQuotes(str string) string {
	var r []rune = []rune(str)
	var s string = ""
	for i := 0; i < len(str); i++ {
		if r[i] == '\'' {
			s = s + " " + "'" + " "
			continue
		}
		s = s + string(r[i])
	}
	str = s

	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	var flag bool
	for i := 0; i < len(slice); i++ {
		if slice[i] == "'" && !flag {
			slice[i+1] = "'" + slice[i+1]
			slice[i] = ""
			flag = true
		} else if slice[i] == "'" && flag {
			slice[i-1] = slice[i-1] + "'"
			slice[i] = ""
			flag = false
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

func DoubleQuotes(str string) string {
	var r []rune = []rune(str)
	var s string = ""
	for i := 0; i < len(str); i++ {
		if r[i] == '"' {
			s = s + " " + "\"" + " "
			continue
		}
		s = s + string(r[i])
	}
	str = s

	slice := strings.Split(str, " ")

	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	var flag bool
	for i := 0; i < len(slice); i++ {
		if slice[i] == "\"" && !flag {
			slice[i+1] = "\"" + slice[i+1]
			slice[i] = ""
			flag = true
		} else if slice[i] == "\"" && flag {
			slice[i-1] = slice[i-1] + "\""
			slice[i] = ""
			flag = false
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
