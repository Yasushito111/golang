package main

import "fmt"

func maina() {
	message := "L fdph, L vdz, L frqtxhuhb."

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
}
