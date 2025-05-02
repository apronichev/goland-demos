package main

import "fmt"

type Point struct {
	X, Y float64
}

func CreatePoint(x, y float64) *Point {
	var point *Point
	point.X = x
	point.Y = y
	return point
}

func main() {
	p := CreatePoint(10, 20)
	fmt.Printf("point = (%f, %f)", p.X, p.Y)
}
