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

func NewAspectRatio(w int, h int, c AspectRatioControlAxis) *AspectRatio {
    return &AspectRatio{
        Dimensions: Dimensions {
            Width: w,
            Height: h,
        },
        ControlAxis: c,
    }
}

func (this *AspectRatio) NewDimensions(v int) Dimensions {
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