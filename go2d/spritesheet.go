package go2d

import (
    "image"
    "os"
)

type SpriteSheet struct {
    Row    int
    Column int

    image  image.Image
}

func NewSpriteSheet(path string, column int, row int) (*SpriteSheet, error) {
    imgf, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    i, _, err := image.Decode(imgf)
    if err != nil {
        return nil, err
    }

    return &SpriteSheet {
        image: i,
        Column: column,
        Row: row,
    }, nil
}

func (this *SpriteSheet) GetSprite(location Vector) image.Image {
    return this.image.(interface {
        SubImage(r image.Rectangle) image.Image
    }).SubImage(
        image.Rect(
            int(location.X*float64(this.Row)),
            int(location.Y*float64(this.Column)),
            int((location.X*float64(this.Row))+float64(this.Row)),
            int((location.Y*float64(this.Column))+float64(this.Column))),
    )
}
