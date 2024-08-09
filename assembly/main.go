package main

import (
	"fmt"
	"runtime"
	"time"
)

func Add(x, y int64) int64

func main() {
	a := int64(10)
	b := int64(20)
	result := add(a, b)
	print(result)

	fmt.Println(result) // 输出：30

	runtime.Gosched()
	timer := time.NewTimer(time.Second * 2)
	<-timer.C

	time.Sleep(1)
	//var ab *uintptr

	//cc := &ab
	//s := *(**uintptr)( unsafe.Pointer(&a))
}

func print(result int64) {

}

func add(a, b int64) int64 {
	return a + b
}
