package main

import (
	"fmt"
	"time"
)

func main55() {
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c: //gopherの起床待ち
			fmt.Println("gopher ", gopherID, "はスリープを終えました")
		case <-timeout:
			fmt.Println("忍耐の限界に達した！")
			return
		}
	}
}
