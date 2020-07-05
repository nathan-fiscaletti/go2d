package go2d

type AspectRatioControlAxis int

const (
    AspectRatioControlAxisWidth = iota
    AspectRatioControlAxisHeight
)

type AspectRatio struct {
    Dimensions
    ControlAxis AspectRatioControlAxis
}

func NewAspectRatio(w float64, h float64, c AspectRatioControlAxis) *AspectRatio {
    return &AspectRatio{
        Dimensions: Dimensions {
            Width: w,
            Height: h,
        },
        ControlAxis: c,
    }
}

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