package graphics

import (
	"image"

	"../metrics"
)

type SpriteSheet struct {
	Image  *image.Image
	Row    int
	Column int
}

func (this *SpriteSheet) getSprite(location metrics.Vector) *image.Image {
	return this.Image.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(
		image.Rect(
			location.X*this.Row,
			location.Y*this.Column,
			(location.X*this.Row)+this.Row,
			(location.Y*this.Column)+this.Column),
	)
}
