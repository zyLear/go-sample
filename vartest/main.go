package main

import (
	"fmt"
	"time"
)

//
//import "sync"
//
//func BlockSubmitAndWaitAllFinish(tasks []func()) error {
//	if len(tasks) == 0 {
//		return nil
//	}
//	var waitGroup int
//
//	waitGroup.Add(len(tasks))
//	for i := 0; i < 10; i++ {
//
//
//
//	}
//		// 使用同名变量保留当前遍历状态
//		task := task
//		err := goroutinepool.SubmitBlock(func() {
//			defer waitGroup.Done()
//			task()
//		})
//
//	}
//	waitGroup.Wait()
//	return nil
//}
//
//
//
//func main() {
//
//}
//
//
//

func main() {
	c := make(chan int, 6)

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			c <- i
		}
		time.Sleep(7 * time.Second)
		fmt.Println("close")
		close(c)
	}()

	//for i := range c {
	//	println("range receive:", i)
	//}

	for {
		select {
		case i, ok := <-c:
			println("select receive:", i, " ok:", ok)
		}
	}

	//xx := <-c
	//fmt.Println("done..", xx)

	//select {
	//case <-c:
	//	fmt.Println("done..")
	//}

}
