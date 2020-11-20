package main

import "fmt"

func main() {

	var name string = "Abhinav"
	const a float64 = 95.287
	pass := true
	x := 87

	fmt.Println(len(name))
	fmt.Println(name + "Gajula")

	fmt.Printf("%f \n", a)
	fmt.Printf("%.2f \n", a)
	fmt.Printf("%T \n", name)
	fmt.Printf("%t \n", name)
	fmt.Printf("%t \n", pass)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 7)
	fmt.Printf("%c \n", 6)
	fmt.Printf("%x \n", 87)
	fmt.Printf("%e \n", 7.28)
}