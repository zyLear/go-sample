package main

import (
	"fmt"
	"time"
)

// go run -race main.go
// go run main.go
func main() {
	var value = 0
	go func() {
		for value == 0 {

		}
		fmt.Println("break")
	}()

	go func() {
		fmt.Println("start")
		time.Sleep(5 * time.Second)
		value = 1
		fmt.Println("end")
	}()
	time.Sleep(1 * time.Minute)
}
