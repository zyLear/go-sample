package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	ready := false // cond所等待的条件

	go func() {
		fmt.Println("doing some work...")
		time.Sleep(time.Second)
		ready = true
		cond.Signal()
	}()

	cond.L.Lock() // 检查目标条件时先加锁
	for !ready {
		cond.Wait()
	}
	fmt.Println("got work done signal!")
	cond.L.Unlock() // Wait返回并跳出for循环后需要解锁
}

// OUTPUT:
// doing some work...
// got work done signal!

//func main() {
//	var ch = make(chan struct{})
//	ready := false // cond所等待的条件
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		fmt.Println("doing some work...")
//		time.Sleep(time.Second)
//		ready = true
//		close(ch)
//	}()
//
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		i := i
//		go func() {
//			defer wg.Done()
//			for !ready {
//				<-ch
//			}
//			fmt.Println("got work done signal! ", i)
//		}()
//	}
//
//	wg.Wait()
//}

//func main() {
//	var ch = make(chan struct{})
//	ready := false // cond所等待的条件
//
//	go func() {
//		fmt.Println("doing some work...")
//		time.Sleep(time.Second)
//		ready = true
//		ch <- struct{}{}
//	}()
//
//	for !ready {
//		<-ch
//	}
//	fmt.Println("got work done signal!")
//}

//var (
//	buffer      = make([]int, 0, 10)
//	bufferMutex sync.Mutex
//	bufferCond  = sync.NewCond(&bufferMutex)
//)
//
//func producer() {
//	for i := 0; i < 20; i++ {
//		bufferMutex.Lock()
//		for len(buffer) == cap(buffer) {
//			bufferCond.Wait()
//		}
//		buffer = append(buffer, i)
//		fmt.Printf("Produced: %d\n", i)
//		bufferCond.Signal()
//		bufferMutex.Unlock()
//		time.Sleep(time.Millisecond * 100)
//	}
//}
//
//func consumer() {
//	for i := 0; i < 20; i++ {
//		bufferMutex.Lock()
//		for len(buffer) == 0 {
//			bufferCond.Wait()
//		}
//		item := buffer[0]
//		buffer = buffer[1:]
//		fmt.Printf("Consumed: %d\n", item)
//		bufferCond.Signal()
//		bufferMutex.Unlock()
//		time.Sleep(time.Millisecond * 150)
//	}
//}
//
//func main() {
//	go producer()
//	go consumer()
//
//	time.Sleep(time.Second * 5)
//}

//func main() {
//	cond := sync.NewCond(new(sync.Mutex))
//	ready := false // cond所等待的条件
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		fmt.Println("doing some work...")
//		time.Sleep(time.Second)
//		ready = true
//		cond.Broadcast() // 通知多个被阻塞的goroutine
//	}()
//
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		i := i
//		go func() {
//			defer wg.Done()
//			cond.L.Lock()
//			for !ready {
//				cond.Wait()
//			}
//			fmt.Println("got work done signal! ", i)
//			cond.L.Unlock()
//		}()
//	}
//
//	wg.Wait()
//}

//got work done signal!  3
//got work done signal!  1
//got work done signal!  4
//got work done signal!  0
//got work done signal!  2

//var data atomic.Int64
//
//func read(index int, c *sync.Cond) {
//	c.L.Lock()
//	for data.Load() == 0 {
//		c.Wait()
//	}
//	log.Println("goroutine:", index, "开始读:", data.Add(-1)+1)
//	c.L.Unlock()
//}
//
//func writeSignal(c *sync.Cond) {
//	log.Println("开始写")
//	time.Sleep(time.Second)
//	log.Println("唤醒其中一个goroutine")
//	data.Add(1)
//	c.Signal()
//}
//
//func writeBroadcast(c *sync.Cond) {
//	log.Println("开始写")
//	time.Sleep(time.Second)
//	log.Println("唤醒全部goroutine")
//	data.Add(5)
//	c.Broadcast()
//}
//
//func main() {
//	cond := sync.NewCond(&sync.Mutex{})
//	for i := 1; i <= 5; i++ {
//		go func(index int) {
//			read(index, cond)
//		}(i)
//	}
//	writeBroadcast(cond)
//	time.Sleep(time.Second * 3)
//}
