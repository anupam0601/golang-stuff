package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 5)
	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------------------")
	people := append(mySlice, 1, 2, 3)
	fmt.Println(people)
	fmt.Println(people[:1])
	fmt.Println(people[1:])

	// for i := 0; i < 20; i++ {
	// 	mySlice = append(mySlice, i)
	// 	fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	// }
}
