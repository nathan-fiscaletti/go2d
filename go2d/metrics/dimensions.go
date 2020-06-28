package metrics

type Dimensions struct {
	Width  int
	Height int
}

func NewSquareDimensions(size int) Dimensions {
	return Dimensions{
		Width:  size,
		Height: size,
	}
}

func NewZeroDimensions() Dimensions {
	return NewSquareDimensions(0)
}

func NewAspectDimensions(aspectRatio AspectRatio, v int) Dimensions {
	if aspectRatio.ControlAxis == AspectRatioControlAxisHeight {
		return Dimensions{
			Height: v,
			Width:  v / aspectRatio.Ratio.Height * aspectRatio.Ratio.Width,
		}
	} else {
		return Dimensions{
			Height: v / aspectRatio.Ratio.Width * aspectRatio.Ratio.Height,
			Width:  v,
		}
	}
}

func (this *Dimensions) Plus(other Dimensions) Dimensions {
	return Dimensions{
		Width:  this.Width + other.Width,
		Height: this.Height + other.Height,
	}
}

func (this *Dimensions) PlusHeight(height int) Dimensions {
	return this.Plus(Dimensions{
		Width:  0,
		Height: height,
	})
}

func (this *Dimensions) PlusWidth(width int) Dimensions {
	return this.Plus(Dimensions{
		Width:  width,
		Height: 0,
	})
}

func (this *Dimensions) Recalculate(aspectRatio AspectRatio) {
	if aspectRatio.ControlAxis == AspectRatioControlAxisHeight {
		this.Width = this.Height / aspectRatio.Ratio.Height * aspectRatio.Ratio.Width
	} else if aspectRatio.ControlAxis == AspectRatioControlAxisWidth {
		this.Height = this.Width / aspectRatio.Ratio.Width * aspectRatio.Ratio.Height
	}
}
