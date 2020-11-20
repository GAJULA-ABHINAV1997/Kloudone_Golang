package main

import "fmt"

func main() {
	rect := Rectangle{height: 55, width: 67}

	fmt.Println("Height = ", rect.height)
	fmt.Println("Width = ", rect.width)
	fmt.Println("Area of Rectagle  = ", rect.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}