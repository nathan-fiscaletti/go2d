package go2d

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
