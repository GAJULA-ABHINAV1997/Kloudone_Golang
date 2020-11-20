package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for _, value := range 10 {
		fmt.Println(value)
	}
}