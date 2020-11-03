package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	sumChannel := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		//wg.Add(1)
		go func() {
			sumChannel <- sum + 1
			//wg.Done()
		}()
		sum = <-sumChannel
	}

	wg.Wait()
	fmt.Println(sum)
}
