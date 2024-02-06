package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Myfloat float64

type Vertex struct {
	X, Y float64
}

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := Myfloat(math.Sqrt2)
	v := Vertex{4, -4}

	a = f // Myfloat implements Abser
	fmt.Println("Absolute value of Myfloat:", a.Abs())

	a = &v // *Vertex implements Abser
	fmt.Println("Absolute value of *Vertex:", a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v // This line will result in a compilation error

}
