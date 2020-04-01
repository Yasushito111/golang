package main

import (
	"fmt"
)

func main() {
	const lightSpeed = 299792
	const secondsPerDay = 86400

	var distance int64 = 413e12
	fmt.Println("アルファ•ケンタウリまでの距離は、", distance, "km。")

	days := distance
	fmt.Println("光の速度で、", days, "日かかる。")
}
