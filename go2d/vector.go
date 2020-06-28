package go2d

import (
	"math/rand"
	"time"
)

type Vector struct {
	X int
	Y int
}

func NewRandomVector(max Vector) Vector {
	return NewRandomVectorWithin(
        NewZeroRect(max.X, max.Y),
    )
}

func NewRandomVectorWithin(r Rect) Vector {
	rand.Seed(time.Now().UnixNano())
	return Vector{
		X: rand.Intn(r.X+r.Width-r.X) + r.X,
		Y: rand.Intn(r.Y+r.Height-r.Y) + r.Y,
	}
}

func NewZeroVector() Vector {
	return Vector{}
}

func (this *Vector) ConstrainTo(r Rect) {
	if this.X < r.X {
		this.X = r.X
	}

	if this.Y < r.Y {
		this.Y = r.Y
	}

	if this.X >= r.X + r.Width {
		this.X = r.X + r.Width - 1
	}

	if this.Y >= r.Y + r.Width {
		this.Y = r.Y + r.Width - 1
	}
}

func (this *Vector) Constrained(r Rect) Vector {
	resX := this.X
	resY := this.Y

	if this.X < r.X {
		resX = r.X
	}

	if this.Y < r.Y {
		resY = r.Y
	}

	if this.X >= r.X + r.Width {
		resX = r.X + r.Width - 1
	}

	if this.Y >= r.Y + r.Width {
		resY = r.Y + r.Width - 1
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

func (this *Vector) IsInsideOf(r Rect) bool {
    return r.Contains(*this)
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

func (this *Vector) Negative() Vector {
	return Vector{
		X: -this.X,
		Y: -this.Y,
	}
}
