package main

import (
	"fmt"
	"math"
)

// This shape interface and shape must be able to return its area and perimeter
// The below circle and rectangel fulfill the interface
type shape interface {
    area() float64
    perimeter() float64
}

// rectangle struct
type rect struct {
    width, height float64
}

// rectangle implements the shape interface (FULFILL)
func (r rect) area() float64 {
    return float64(r.width * r.height)
}

func (r rect) perimeter() float64 {
    return float64(2 * r.width + 2 * r.height)
}

// circle struct
type circle struct {
    radius float64
}

// circle implements the shape interface (FULFILL)
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func main() {

    small_rectangle := rect{ width: 10.99, height: 9.10 }
    fmt.Println("Small Rectangle: Area: ", small_rectangle.area())
    fmt.Println("Small Rectangle: Perimeter: ", small_rectangle.perimeter())

    small_circle := circle{ radius: 9.9 }
    fmt.Println("Small Circle: Area: ", small_circle.area())
    fmt.Println("Small Circle: Perimeter: ", small_circle.perimeter())
}
