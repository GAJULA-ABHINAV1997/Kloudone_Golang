package main

import "fmt"

func main() {

	x := 78

	fmt.Println(x)
	fmt.Println(&x)

	value(&x)
	fmt.Println(x)
}

func value(x *int) {
	*x = 18
}