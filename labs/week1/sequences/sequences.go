package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	newSlice := make([]int, len(slice) * 2)
	newSlice = append(slice, slice...)
	slice = newSlice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice{
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int{
	for i, num := range array{
		array[i] = f(num)
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)
	//intsArray := [5]int{1, 2, 3, 4, 5}
	//intsArray = mapArray(addOne, intsArray)
	//fmt.Print(intsArray)

	//newSlice := intsSlice[1:3]
	//mapSlice(square, newSlice)
	//fmt.Println(intsSlice)
	//fmt.Println(newSlice)

	double(intsSlice)
	fmt.Println(intsSlice)
}
