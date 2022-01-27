package main

import "fmt"

func main()  {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("%+v \n", mySlice)

	fullSilce := myArray[:]
	fmt.Printf("%+v \n", fullSilce)

	appendSilce := append(fullSilce, 6)
	fmt.Printf("%+v \n", appendSilce)

	deleteSilce := deleteItem(fullSilce, 0)
	fmt.Printf("%+v \n", deleteSilce)
}

func deleteItem(since []int, index int) []int {
	return append(since[:index], since[index+1:]...)
}
