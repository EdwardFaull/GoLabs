package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ChrisGora/semaphore"
)

type buffer struct {
	b                 []int
	size, read, write int
}

func newBuffer(size int) buffer {
	return buffer{
		b:     make([]int, size),
		size:  size,
		read:  0,
		write: 0,
	}
}

func (buffer *buffer) get() int {
	x := buffer.b[buffer.read]
	fmt.Println("Get\t", x, "\t", buffer)
	buffer.read = (buffer.read + 1) % len(buffer.b)
	return x
}

func (buffer *buffer) put(x int) {
	buffer.b[buffer.write] = x
	//fmt.Println("Put\t", x, "\t", buffer)
	buffer.write = (buffer.write + 1) % len(buffer.b)
}

func producer(buffer *buffer, start, delta int,
	mutex *sync.Mutex, spaceAvailable *semaphore.Semaphore, workAvailable *semaphore.Semaphore) {
	x := start
	for {
		(*spaceAvailable).Wait()
		mutex.Lock()
		buffer.put(x)
		x = x + delta
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		(*workAvailable).Post()
		mutex.Unlock()
	}
}

func consumer(buffer *buffer, mutex *sync.Mutex, spaceAvailable *semaphore.Semaphore, workAvailable *semaphore.Semaphore) {
	for {
		(*workAvailable).Wait()
		mutex.Lock()
		_ = buffer.get()
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		(*spaceAvailable).Post()
		mutex.Unlock()
	}
}

func main() {

	mutex := sync.Mutex{}

	spaceAvailable := semaphore.Init(5, 5)
	workAvailable := semaphore.Init(5, 0)

	buffer := newBuffer(5)

	go producer(&buffer, 1, 1, &mutex, &spaceAvailable, &workAvailable)
	go producer(&buffer, 1000, -1, &mutex, &spaceAvailable, &workAvailable)

	consumer(&buffer, &mutex, &spaceAvailable, &workAvailable)
}
