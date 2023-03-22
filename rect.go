package go2d

// RectSide is an enumeration of the sides of a rectangle.
type RectSide int

const (
	RectSideLeft = iota
	RectSideRight
	RectSideTop
	RectSideBottom
)

// Rect is a simple rectangle implementation.
type Rect struct {
	Vector
	Dimensions
}

// NewZeroRect creates a new rectangle with the given width and height at the
// origin.
func NewZeroRect(w, h float64) Rect {
	return NewRect(0, 0, w, h)
}

// NewRect creates a new rectangle with the given width and height at the given
// position.
func NewRect(x, y, w, h float64) Rect {
	return Rect{
		Vector: Vector{
			X: x,
			Y: y,
		},
		Dimensions: Dimensions{
			Width:  w,
			Height: h,
		},
	}
}

// Equals returns true if the given rectangle is equal to this rectangle.
func (this Rect) Equals(other Rect) bool {
	return this.X == other.X && this.Y == other.Y &&
		this.Width == other.Width && this.Height == other.Height
}

// IntersectsWith returns true if this rectangle intersects with the given
// rectangle.
func (this Rect) IntersectsWith(other Rect) bool {
	return this.X < other.X+other.Width &&
		this.X+this.Width > other.X &&
		this.Y < other.Y+other.Height &&
		this.Y+this.Height > other.Y
}

// Contains returns true if the given vector is contained within this rectangle.
func (this Rect) Contains(v Vector) bool {
	return !(v.X < this.X ||
		v.Y < this.Y ||
		v.X > this.X+this.Width ||
		v.Y > this.Y+this.Height)
}

// Center returns the center of this rectangle.
func (this Rect) Center() Vector {
	return Vector{
		X: this.X + this.Width/2,
		Y: this.Y + this.Height/2,
	}
}

// Constrain constrains this rectangle to the given rectangle. If this rectangle
// is outside of the given rectangle, it will be moved to the closest point on
// the rectangle. The return value is a slice of the sides that were constrained.
func (this *Rect) Constrain(r Rect) []RectSide {
	res := []RectSide{}

	if this.X < r.X {
		this.X = r.X
		res = append(res, RectSideLeft)
	}

	if this.X+this.Width > r.X+r.Width {
		this.X = r.X + r.Width - this.Width
		res = append(res, RectSideRight)
	}

	if this.Y < r.Y {
		this.Y = r.Y
		res = append(res, RectSideTop)
	}

	if this.Y+this.Height > r.Y+r.Height {
		this.Y = r.Y + r.Height - this.Height
		res = append(res, RectSideBottom)
	}

	return res
}
