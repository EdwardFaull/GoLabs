package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 0; i < 5; i++{
		go helloRoutine(i)
	}
	time.Sleep(1 * time.Second)
}

func helloRoutine(x int){
	fmt.Println("Hello from goroutine ", x)
}