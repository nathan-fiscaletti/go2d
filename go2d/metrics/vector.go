package metrics

import (
	"math/rand"
	"time"
)

type Vector struct {
	X int
	Y int
}

func NewRandomVector(max Vector) Vector {
	return NewFullRandomVector(Vector{
		X: 0,
		Y: 0,
	}, max)
}

func NewFullRandomVector(min Vector, max Vector) Vector {
	rand.Seed(time.Now().UnixNano())
	return Vector{
		X: rand.Intn(max.X-min.X) + min.X,
		Y: rand.Intn(max.Y-min.Y) + min.Y,
	}
}

func NewZeroVector() Vector {
	return Vector{}
}

func (this *Vector) Constrain(min Vector, max Vector) {
	if this.X < min.X {
		this.X = min.X
	}

	if this.Y < min.Y {
		this.Y = min.Y
	}

	if this.X > max.X {
		this.X = max.X
	}

	if this.Y > max.Y {
		this.Y = max.Y
	}
}

func (this *Vector) Constrained(min Vector, max Vector) Vector {
	resX := this.X
	resY := this.Y

	if resX < min.X {
		resX = min.X
	}

	if resY < min.Y {
		resY = min.Y
	}

	if resX > max.X {
		resX = max.X
	}

	if resY > max.Y {
		resY = max.Y
	}

	return Vector{
		X: resX,
		Y: resY,
	}
}

func (this *Vector) PlusY(y int) Vector {
	return Vector{
		X: this.X,
		Y: this.Y + y,
	}
}

func (this *Vector) PlusX(x int) Vector {
	return Vector{
		X: this.X + x,
		Y: this.Y,
	}
}

func (this *Vector) Plus(v Vector) Vector {
	return Vector{
		X: this.X + v.X,
		Y: this.Y + v.Y,
	}
}

func (this *Vector) Times(v Vector) Vector {
	return Vector{
		X: this.X * v.X,
		Y: this.Y * v.Y,
	}
}

func (this *Vector) DividedBy(v Vector) Vector {
	return Vector{
		X: this.X / v.X,
		Y: this.Y / v.Y,
	}
}

func (this *Vector) IsLeftOf(v Vector) bool {
	return this.X < v.X
}

func (this *Vector) IsRightOf(v Vector) bool {
	return this.X > v.X
}

func (this *Vector) IsAbove(v Vector) bool {
	return this.Y < v.Y
}

func (this *Vector) IsBelow(v Vector) bool {
	return this.Y > v.Y
}

func (this *Vector) IsZero() bool {
	return this.Y == 0 && this.X == 0
}

func (this *Vector) IsInsideOf(r Rect) bool {
	return !(this.X < r.Position.X ||
		this.Y < r.Position.Y ||
		this.X > r.Position.X+r.Size.Width ||
		this.Y > r.Position.Y+r.Size.Height)
}

func (this *Vector) Negative() Vector {
	return Vector{
		X: -this.X,
		Y: -this.Y,
	}
}
