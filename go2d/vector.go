package go2d

import (
    "math/rand"
    "time"
)

type Vector struct {
    X int
    Y int
}

func NewVector(x, y int) Vector {
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
        X: rand.Intn(r.X+r.Width-r.X) + r.X,
        Y: rand.Intn(r.Y+r.Height-r.Y) + r.Y,
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