package style

import (
	"image"
	"image/color"
)

type circle struct {
	p image.Point
	r int
}

// NewCircle return a circle style
func NewCircle(img image.Image) MDAvatarBuildStyle {
	w := img.Bounds().Max.X - img.Bounds().Min.X
	h := img.Bounds().Max.Y - img.Bounds().Min.Y

	d := w
	if w > h {
		d = h
	}

	d /= 2
	return &circle{image.Point{X: d, Y: d}, d}
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
