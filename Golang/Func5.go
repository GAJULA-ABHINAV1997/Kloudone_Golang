package main

import "fmt"

func main() {
	ID := []int{1, 2, 3, 4, 5}

	for i, value := range ID {
		fmt.Println(value, i)
	}

	fmt.Println(ID)
	value1 := ID[0:2]
	fmt.Println(value1)
	value2 := ID[3:5]
	fmt.Println(value2)

	copy(value1, value2)
	fmt.Println(value1)

	value3 := append(ID, 6, 7, 8, 9, 10)
	fmt.Println(value3)
}