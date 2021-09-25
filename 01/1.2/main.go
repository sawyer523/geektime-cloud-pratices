package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("receiving: ", v)
	}
}
