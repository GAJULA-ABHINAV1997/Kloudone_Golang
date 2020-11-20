
  
package main

import "fmt"

func main() {

	x, y := 47, 67
	fmt.Println(add(x, y))

	fmt.Println(fact(x))

	fmt.Println(addnumber(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
}

func add(n1, n2 int) int {
	return (n1 + n2)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func addnumber(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}