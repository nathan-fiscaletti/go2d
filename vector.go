package go2d

import (
	"math"
	"math/rand"
	"time"
)

// TICK_DURATION is intended to be used as the duration for a VelocityVector
// that should move the given amount in one tick.
const TICK_DURATION = time.Duration(0)

// VelocityVector is a vector that has a duration. The duration is the amount
// of time it should take for the vector to move from one point to another. If
// the duration is 0, then the vector will move the given amount in one tick.
type VelocityVector struct {
	Vector
	Duration time.Duration
}

// NewVelocityVector creates a new VelocityVector with the given x and y
// components and duration. The duration is the amount of time it should take
// for the vector to move from one point to another. If the duration is 0, then
// the vector will move the given amount in one tick.
func NewVelocityVector(x, y float64, d time.Duration) VelocityVector {
	return VelocityVector{
		Vector: Vector{
			X: x,
			Y: y,
		},
		Duration: d,
	}
}

// GetNextMovement returns the next movement that should be applied to the
// vector. This is used to calculate the next movement of the vector based on
// the duration. If the duration is 0, then the vector will move the given
// amount in one tick.
func (this VelocityVector) GetNextMovement() Vector {
	if this.Duration == TICK_DURATION {
		return this.Vector
	}

	framesPerNanoSecond := 60 * time.Second
	desiredXMovementPerNanoSecond := this.Vector.X / float64(this.Duration)
	desiredYMovementPerNanoSecond := this.Vector.Y / float64(this.Duration)

	return Vector{
		X: float64(framesPerNanoSecond) / desiredXMovementPerNanoSecond,
		Y: float64(framesPerNanoSecond) / desiredYMovementPerNanoSecond,
	}
}

// Vector is a 2D vector with an X and Y component.
type Vector struct {
	X float64
	Y float64
}

// NewVector creates a new Vector with the given x and y components.
func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

// NewRandomVector creates a new Vector with random x and y components within
// the given maximum.
func NewRandomVector(max Vector) Vector {
	return NewRandomVectorWithin(
		NewZeroRect(max.X, max.Y),
	)
}

// NewRandomVectorWithin creates a new Vector with random x and y components
// within the given rectangle.
func NewRandomVectorWithin(r Rect) Vector {
	rand.Seed(time.Now().UnixNano())
	return Vector{
		X: r.X + rand.Float64()*(r.X+r.Width-r.X),
		Y: r.Y + rand.Float64()*(r.Y+r.Width-r.Y),
	}
}

// NewZeroVector creates a new Vector with x and y components of 0.
func NewZeroVector() Vector {
	return Vector{}
}

// DirectionUp returns a Vector with a y component of -1.
func DirectionUp() Vector {
	return Vector{
		Y: -1,
	}
}

// DirectionDown returns a Vector with a y component of 1.
func DirectionDown() Vector {
	return Vector{
		Y: 1,
	}
}

// DirectionLeft returns a Vector with an x component of -1.
func DirectionLeft() Vector {
	return Vector{
		X: -1,
	}
}

// DirectionRight returns a Vector with an x component of 1.
func DirectionRight() Vector {
	return Vector{
		X: 1,
	}
}

// DirectionTo returns a Vector that points from this vector to the given
// vector.
func (this *Vector) DirectionTo(other Vector) Vector {
	angle := this.AngleTo(other)
	return Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

// DistanceTo returns the distance between this vector and the given vector.
func (this *Vector) DistanceTo(other Vector) float64 {
	return math.Sqrt(math.Pow(this.X-other.X, 2) + math.Pow(this.Y-other.Y, 2))
}

// AngleTo returns the angle between this vector and the given vector.
func (this *Vector) AngleTo(other Vector) float64 {
	return math.Atan2(other.Y-this.Y, other.X-this.X)
}

// ConstrainTo constrains this vector to the given rectangle. If the vector is
// outside of the rectangle, it will be moved to the closest point on the
// rectangle.
func (this *Vector) ConstrainTo(r Rect) {
	if this.X < r.X {
		this.X = r.X
	}

	if this.Y < r.Y {
		this.Y = r.Y
	}

	if this.X >= r.X+r.Width {
		this.X = r.X + r.Width - 1
	}

	if this.Y >= r.Y+r.Width {
		this.Y = r.Y + r.Width - 1
	}
}

// Copy returns a copy of this vector.
func (this *Vector) Copy() Vector {
	return Vector{
		X: this.X,
		Y: this.Y,
	}
}

// Constrained returns a copy of this vector constrained to the given
// rectangle. If the vector is outside of the rectangle, it will be moved to
// the closest point on the rectangle.
func (this Vector) Constrained(r Rect) Vector {
	copy := this.Copy()
	copy.ConstrainTo(r)
	return copy
}

// IsInsideOf returns true if this vector is inside of the given rectangle.
func (this Vector) IsInsideOf(r Rect) bool {
	return r.Contains(this)
}

// IsLeftOf returns true if this vector is to the left of the given vector.
func (this Vector) IsLeftOf(v Vector) bool {
	return this.X < v.X
}

// IsRightOf returns true if this vector is to the right of the given vector.
func (this Vector) IsRightOf(v Vector) bool {
	return this.X > v.X
}

// IsAbove returns true if this vector is above the given vector.
func (this Vector) IsAbove(v Vector) bool {
	return this.Y < v.Y
}

// IsBelow returns true if this vector is below the given vector.
func (this Vector) IsBelow(v Vector) bool {
	return this.Y > v.Y
}

// IsZero returns true if this vector is equal to the zero vector.
func (this Vector) IsZero() bool {
	return this.Equals(NewZeroVector())
}

// Equals returns true if this vector is equal to the given vector.
func (this Vector) Equals(other Vector) bool {
	return this.X == other.X && this.Y == other.Y
}

// Inverted returns a copy of this vector with its x and y components inverted.
func (this Vector) Inverted() Vector {
	return Vector{
		X: -this.X,
		Y: -this.Y,
	}
}

// InvertedY returns a copy of this vector with its y component inverted.
func (this Vector) InvertedY() Vector {
	return Vector{
		X: this.X,
		Y: -this.Y,
	}
}

// InvertedX returns a copy of this vector with its x component inverted.
func (this Vector) InvertedX() Vector {
	return Vector{
		X: -this.X,
		Y: this.Y,
	}
}
