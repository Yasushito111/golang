package main

import (
	"fmt"
	"math/big"
)

func maina() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("2400000000000000000000", 10)
	fmt.Println("アンドロメダ銀河までの距離は、", distance, "km。")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("光の速度で、", days, "日かかる。")
}
