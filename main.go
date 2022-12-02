package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var kata = "Kemarin Shopia per[gi ke mall."
	var string = strings.Split(kata, " ")

	var regex, _ = regexp.Compile(`[a-z,A-Z,?,-.]+`)
	var jumlah = 0

	for i := 0; i < len(string); i++ {
		string[i] = regex.ReplaceAllString(string[i], "*")
		if string[i] == "*" {
			jumlah += 1
		}
	}
	fmt.Println(jumlah)
}
