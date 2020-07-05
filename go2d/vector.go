package go2d

import (
    "math"
    "math/rand"
    "time"
)

type Vector struct {
    X float64
    Y float64
}

func NewVector(x, y float64) Vector {
    return Vector{
        X: x,
        Y: y,
    }
}

func NewRandomVector(max Vector) Vector {
    return NewRandomVectorWithin(
        NewZeroRect(max.X, max.Y),
    )
}

func NewRandomVectorWithin(r Rect) Vector {
    rand.Seed(time.Now().UnixNano())
    return Vector{
        X: r.X + rand.Float64() * (r.X+r.Width - r.X),
        Y: r.Y + rand.Float64() * (r.Y+r.Width - r.Y),
    }
}

func NewZeroVector() Vector {
    return Vector{}
}

func DirectionUp() Vector {
    return Vector {
        Y: -1,
    }
}

func DirectionDown() Vector {
    return Vector {
        Y: 1,
    }
}

func DirectionLeft() Vector {
    return Vector {
        X: -1,
    }
}

func DirectionRight() Vector {
    return Vector {
        X: 1,
    }
}

func (this *Vector) DirectionTo(other Vector) Vector {
    angle := this.AngleTo(other)
    return Vector{
        X: math.Cos(angle),
        Y: math.Sin(angle),
    }
}

func (this *Vector) DistanceTo(other Vector) float64 {
    return math.Sqrt(math.Pow(this.X - other.X, 2) + math.Pow(this.Y - other.Y, 2))
}

func (this *Vector) AngleTo(other Vector) float64 {
    return math.Atan2(other.Y - this.Y, other.X - this.X)
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

func (this *Vector) Copy() Vector {
    return Vector{
        X: this.X,
        Y: this.Y,
    }
}

func (this Vector) Constrained(r Rect) Vector {
    copy := this.Copy()
    copy.ConstrainTo(r)
    return copy
}

func (this Vector) IsInsideOf(r Rect) bool {
    return r.Contains(this)
}

func (this Vector) IsLeftOf(v Vector) bool {
    return this.X < v.X
}

func (this Vector) IsRightOf(v Vector) bool {
    return this.X > v.X
}

func (this Vector) IsAbove(v Vector) bool {
    return this.Y < v.Y
}

func (this Vector) IsBelow(v Vector) bool {
    return this.Y > v.Y
}

func (this Vector) IsZero() bool {
    return this.Equals(NewZeroVector())
}

func (this Vector) Equals(other Vector) bool {
    return this.X == other.X && this.Y == other.Y
}

func (this Vector) Inverted() Vector {
    return Vector{
        X: -this.X,
        Y: -this.Y,
    }
}

func (this Vector) InvertedY() Vector {
    return Vector{
        X: this.X,
        Y: -this.Y,
    }
}

func (this Vector) InvertedX() Vector {
    return Vector{
        X: -this.X,
        Y: this.Y,
    }
}