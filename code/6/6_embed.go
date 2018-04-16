package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type TextPoint struct {
	Point
	Text string
}

func main() {
	var p = TextPoint{Point{1, 1}, "Point P"}
	var q = TextPoint{Point{5, 4}, "Point q"}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.Distance(q)                    // compile error: cannot use q (TextPoint) as Point
}
