package main

import (
	"fmt"
)

var PointRoot = WirePoint{Point{0, 0}, 0}

type Point struct {
	x, y int
}

type WirePoint struct {
	Point
	ordinal int
}

var directions = map[byte]Point{
	'R': Point{0, 1},
	'L': Point{0, -1},
	'U': Point{1, 0},
	'D': Point{-1, 0},
}

func (p WirePoint) add(p1 Point) WirePoint {
	return WirePoint{Point{p.x + p1.x, p.y + p1.y}, p.ordinal + 1}
}

func (p WirePoint) apply(direction byte, times int) Wire {
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

type Wire []WirePoint

func (a Wire) Len() int           { return len(a) }
func (a Wire) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Wire) Less(i, j int) bool { return a[i].ordinal < a[j].ordinal }
