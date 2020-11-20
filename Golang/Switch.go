package main

import "fmt"

func main() {

	age := 87

	if age >= 18 {
		fmt.Println("You are Entered!!")
	} else {
		fmt.Println("You Can't be Entered!!")
	}

	mark := 111

	switch mark {
	case 40:
		fmt.Println("fail")
	case 50:
		fmt.Println("Pass")
	case 90:
		fmt.Println("Topper")
	default:
		fmt.Println("Invalid")

	}
}