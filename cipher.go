package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "you message goes here"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))
	for i := 0; 1 < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			// A=O, B=I, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
