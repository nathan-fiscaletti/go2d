package go2d

import(
    "image"

    "./graphics"
    "./metrics"

    "github.com/tfriedel6/canvas"
)

type ImageDrawable struct {
    graphics.Drawable

    gImg image.Image

    canvas *canvas.Canvas
    cImg *canvas.Image
}

func NewImageDrawable(c *canvas.Canvas, img image.Image) *ImageDrawable {
    i,err := c.LoadImage(img)
    if err != nil {
        panic(err)
    }
    return &ImageDrawable {
        gImg: img,
        cImg: i,
        canvas: c,
        Drawable: graphics.Drawable {
            Bounds: metrics.Rect {
                Dimensions: metrics.Dimensions {
                    Width: img.Bounds().Dx(),
                    Height: img.Bounds().Dy(),
                },
            },
        },
    }
}

func (this *ImageDrawable) Render() {
    this.canvas.DrawImage(this.cImg, float64(this.Bounds.X), float64(this.Bounds.Y), float64(this.Bounds.Width), float64(this.Bounds.Height))
}
