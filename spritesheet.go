package go2d

import (
	"image"
	"os"
)

// SpriteSheet is a simple sprite sheet implementation.
type SpriteSheet struct {
	// RowSize is the size of each row in the sprite sheet.
	RowSize int
	// ColumnSize is the size of each column in the sprite sheet.
	ColumnSize int

	image image.Image
}

// NewSpriteSheet creates a new sprite sheet from the given image path with the
// given column and row size.
func NewSpriteSheet(path string, columnSize int, rowSize int) (*SpriteSheet, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return &SpriteSheet{
		image:      i,
		ColumnSize: columnSize,
		RowSize:    rowSize,
	}, nil
}

// GetSprite returns the sprite at the given location.
func (this *SpriteSheet) GetSprite(location Vector) image.Image {
	return this.image.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(
		image.Rect(
			int(location.X*float64(this.RowSize)),
			int(location.Y*float64(this.ColumnSize)),
			int((location.X*float64(this.RowSize))+float64(this.RowSize)),
			int((location.Y*float64(this.ColumnSize))+float64(this.ColumnSize))),
	)
}
