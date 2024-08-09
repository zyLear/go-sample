package main

import (
	_ "awesomeProject1/inittestoutside"
	"fmt"
	"time"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")

	timer := time.NewTicker(3 * time.Second)
	for {
		<-timer.C
		if err := loadCache(); err != nil {
		}
	}

}

var count = 1

func loadCache() interface{} {
	fmt.Println("start")
	if count == 2 {
		fmt.Println("===2")
		time.Sleep(15 * time.Second)
	}
	fmt.Println("end")
	count++
	return nil
}
