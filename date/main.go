package main

import (
	"fmt"
	"time"
)

func main() {

	parse, err := time.Parse("20060102", "20250331")
	if err != nil {
		return
	}
	fmt.Println(parse.AddDate(0, -1, 0))
	fmt.Println(parse.AddDate(0, -1, -1))
	fmt.Println(parse.AddDate(0, -1, -2))
	fmt.Println(parse.AddDate(0, -1, -3))

	//parse, err = time.Parse("0601", "2503")
	//if err != nil {
	//	return
	//}
	fmt.Println(parse.AddDate(0, -1, 1-parse.Day()))
}
