package main

import (
	"log"
	"os"
	"runtime/trace"
)

// merge takes two sorted sub-arrays from slice and sorts them.
// The resulting array is put back in slice.
func merge(slice []int32, middle int) {
	sliceClone := make([]int32, len(slice))
	copy(sliceClone, slice)
	a := sliceClone[middle:]
	b := sliceClone[:middle]
	i := 0
	j := 0
	for k := 0; k < len(slice); k++ {
		if i >= len(a) {
			slice[k] = b[j]
			j++
		} else if j >= len(b) {
			slice[k] = a[i]
			i++
		} else if a[i] > b[j] {
			slice[k] = b[j]
			j++
		} else {
			slice[k] = a[i]
			i++
		}
	}
}

// Sequential merge sort.
func mergeSort(slice []int32) {
	if len(slice) > 1 {
		middle := len(slice) / 2
		mergeSort(slice[:middle])
		mergeSort(slice[middle:])
		merge(slice, middle)
	}
}

const threshold  = 16

// TODO: Parallel merge sort.
func parallelMergeSort(slice []int32, done chan bool) {
	if len(slice) > 1 {
		if len(slice) < threshold{
			mergeSort(slice)
		} else {
			middle := len(slice) / 2
			done1 := make(chan bool)
			done2 := make(chan bool)
			go parallelMergeSort(slice[:middle], done1)
			go parallelMergeSort(slice[middle:], done2)

			i := 0

			for i < 2 {
				select {
				case <-done1:
					i++
				case <-done2:
					i++
				}
			}
			merge(slice, middle)
		}
	}

	done <- true
}

// main starts tracing and in parallel sorts a small slice.
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	slice := make([]int32, 0, 100)
	for i := int32(100); i > 0; i-- {
		slice = append(slice, i)
	}

	done := make(chan bool)
	go parallelMergeSort(slice, done)
}
