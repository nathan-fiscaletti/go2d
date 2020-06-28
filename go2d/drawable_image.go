package go2d

import(
    "image"

    "github.com/tfriedel6/canvas"
)

type ImageDrawable struct {
    Drawable

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
        Drawable: Drawable {
            Visible: true,
            Bounds: Rect {
                Dimensions: Dimensions {
                    Width: img.Bounds().Dx(),
                    Height: img.Bounds().Dy(),
                },
            },
        },
    }
}

func (this *ImageDrawable) GetImage() image.Image {
    return this.gImg
}

func (this *ImageDrawable) Render() {
    if this.Visible {
        this.canvas.DrawImage(
            this.cImg, 
            float64(this.Bounds.X), 
            float64(this.Bounds.Y), 
            float64(this.Bounds.Width), 
            float64(this.Bounds.Height),
        )
    }
}
