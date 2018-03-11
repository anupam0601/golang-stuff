package main

import (
	"fmt"
)

// Array of fixed size and type
var anupam = [3]int{1, 2, 3}

// undetermined array
var debnath = [...]int{4, 5, 6, 7, 8}

func main() {
	q := [3]string{"go", "java", "python"}
	fmt.Println("array- ", anupam)
	fmt.Println("array2- ", q)
	fmt.Println("undetermined-", debnath)
}
