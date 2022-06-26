package main

import (
	"log"
	"os"
)

func main() {
	// too many logic on main function. It must just run process or/and handle data and send to other secondary func-s.
	args := os.Args[1:]

	if !isValid(args) {
		log.Fatalln("wrong input argument")
	}

	text, err := readSample(args[0])
	if err != nil {
		log.Fatalln(err)
	}
<<<<<<< HEAD
	defer file1.Close()
	var str string
	var strSlice []string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		strSlice = append(strSlice, scanner.Text())
	}
	str = strings.Join(strSlice, "\n")
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		switch slice[i] {
		case "(up)":
			str = funcUp.Up(str)
		case "(up,":
			str = funcUp.Up(str)
		case "(low)":
			str = funcLow.Low(str)
		case "(low,":
			str = funcLow.Low(str)
		case "(cap)":
			str = funcCap.Cap(str)
		case "(cap,":
			str = funcCap.Cap(str)
		case "(hex)":
			str = funcHex.Hex(str)
		case "(bin)":
			str = funcBin.Bin(str)
		}
		if slice[i] == "," || slice[i] == "." || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" || slice[i] == "..." || slice[i] == "?!" && (slice[i][0] == 46 || slice[i][0] == 44 || slice[i][0] == 33 || slice[i][0] == 63 || slice[i][0] == 58 ||
			slice[i][0] == 59) {
			str = funcPunct.Punct(str)
		}
=======
>>>>>>> audit-d0ssan

	text = convertText(text)

	err = writeResult(args[1], text)
	if err != nil {
		log.Fatalln(err)
	}
}
