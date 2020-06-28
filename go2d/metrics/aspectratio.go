package metrics

type AspectRatioControlAxis int

const (
	AspectRatioControlAxisWidth = iota
	AspectRatioControlAxisHeight
)

type AspectRatio struct {
	Ratio       Dimensions
	ControlAxis AspectRatioControlAxis
}
