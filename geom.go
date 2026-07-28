package control

import (
	"math"

	"github.com/ecwid/control/protocol/dom"
)

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Quad []Point

func (p Point) Equal(a Point) bool {
	return p.X == a.X && p.Y == a.Y
}

func convertQuads(dq []dom.Quad) []Quad {
	var p = make([]Quad, len(dq))
	for n, q := range dq {
		p[n] = Quad{
			Point{q[0], q[1]},
			Point{q[2], q[3]},
			Point{q[4], q[5]},
			Point{q[6], q[7]},
		}
	}
	return p
}

// Middle calc middle of quad
func (q Quad) Middle() Point {
	x := 0.0
	y := 0.0
	for i := 0; i < 4; i++ {
		x += q[i].X
		y += q[i].Y
	}
	return Point{X: x / 4, Y: y / 4}
}

func (q Quad) Area() float64 {
	var area float64
	var x1, x2, y1, y2 float64
	var vertices = len(q)
	for i := 0; i < vertices; i++ {
		x1 = q[i].X
		y1 = q[i].Y
		x2 = q[(i+1)%vertices].X
		y2 = q[(i+1)%vertices].Y
		area += (x1*y2 - x2*y1) / 2
	}
	return math.Abs(area)
}
