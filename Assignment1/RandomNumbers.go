package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(20)
	fmt.Printf("%d ", rand.Intn(100))
	fmt.Printf("%d ", rand.Intn(100))
	fmt.Printf("%d \n", rand.Intn(100))

	rand.Seed(20)
	fmt.Printf("%d ", rand.Intn(100))
	fmt.Printf("%d ", rand.Intn(100))
	fmt.Printf("%d \n", rand.Intn(100))

	fmt.Println()
}
