package go2d

// Dimensions is a simple struct that represents the width and height of an object.
type Dimensions struct {
	Width  float64
	Height float64
}

// NewSquareDimensions creates a new Dimensions struct with equal width and height.
func NewSquareDimensions(size float64) Dimensions {
	return Dimensions{
		Width:  size,
		Height: size,
	}
}

// NewZeroDimensions creates a new Dimensions struct with zero width and height.
func NewZeroDimensions() Dimensions {
	return NewSquareDimensions(0)
}

// Plus returns a new Dimensions struct that is the sum of this Dimensions struct
// and the given Dimensions struct.
func (this *Dimensions) Plus(other Dimensions) Dimensions {
	return Dimensions{
		Width:  this.Width + other.Width,
		Height: this.Height + other.Height,
	}
}

// PlusHeight returns a new Dimensions struct that is the sum of this Dimensions
// struct and the given height.
func (this *Dimensions) PlusHeight(height float64) Dimensions {
	return this.Plus(Dimensions{
		Width:  0,
		Height: height,
	})
}

// PlusWidth returns a new Dimensions struct that is the sum of this Dimensions
// struct and the given width.
func (this *Dimensions) PlusWidth(width float64) Dimensions {
	return this.Plus(Dimensions{
		Width:  width,
		Height: 0,
	})
}
