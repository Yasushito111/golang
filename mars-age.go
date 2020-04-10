package main

import(
	"fmt"
)

func main{
	age := 41
	marsAge := float64(age)
	marsDays := 678.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("私は火星では",marsAge, "歳です。")

}