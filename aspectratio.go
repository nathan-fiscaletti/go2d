package go2d

// AspectRatioControlAxis is an enum that defines which axis of an aspect ratio
// should be in control when calculating new dimensions.
type AspectRatioControlAxis int

const (
	AspectRatioControlAxisWidth = iota
	AspectRatioControlAxisHeight
)

// AspectRatio is a simple aspect ratio implementation.
type AspectRatio struct {
	Dimensions
	ControlAxis AspectRatioControlAxis
}

// NewAspectRatio creates a new aspect ratio with the given width, height, and
// control axis.
func NewAspectRatio(w float64, h float64, c AspectRatioControlAxis) *AspectRatio {
	return &AspectRatio{
		Dimensions: Dimensions{
			Width:  w,
			Height: h,
		},
		ControlAxis: c,
	}
}

// NewDimensions returns a new Dimensions struct that is the result of applying
// this aspect ratio to the given value on the control axis.
func (this *AspectRatio) NewDimensions(v float64) Dimensions {
	if this.ControlAxis == AspectRatioControlAxisHeight {
		return Dimensions{
			Height: v,
			Width:  v / this.Height * this.Width,
		}
	} else {
		return Dimensions{
			Height: v / this.Width * this.Height,
			Width:  v,
		}
	}
}
