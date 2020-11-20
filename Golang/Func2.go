package main

import "fmt"

func main() {

	defer run1()
	run2()

	fmt.Println(div(65, 37))
	thepanic()
}

func run1()  { fmt.Println("I was First") }
func run2() { fmt.Println("I was Second") }

func div(n1, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	return (n1 / n2)
}

func thepanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Enjoyed")
}
