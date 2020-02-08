package main

import (
	"fmt"
	"math"
)

var PointRoot = Point{0, 0}

type Point struct {
	x, y int
}

var directions = map[byte]Point{
	'R': Point{0, 1},
	'L': Point{0, -1},
	'U': Point{1, 0},
	'D': Point{-1, 0},
}

func (p Point) add(p1 Point) Point {
	return Point{p.x + p1.x, p.y + p1.y}
}

func (p Point) apply(direction byte, times int) Wire {
	dir := directions[direction]
	res := make(Wire, times)
	curr := p

	for i := 0; i < times; i++ {
		curr = curr.add(dir)
		res[i] = curr
	}
	return res
}

func (p Point) hash() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

func (p1 Point) distance(p2 Point) int {
	return int(math.Abs(float64(p1.x+p2.x)) + math.Abs(float64(p1.y+p2.y)))
}

type Wire []Point

func (a Wire) Len() int           { return len(a) }
func (a Wire) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Wire) Less(i, j int) bool { return a[i].distance(PointRoot) < a[j].distance(PointRoot) }
