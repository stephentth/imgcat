package main

type Color struct {
	R, G, B uint32
}

func NewColor(r, g, b uint32) Color {
	return Color{r, g, b}
}

func (c *Color) NormalizeValue() {
	c.R = uint32(float64(c.R) / 65535 * 255)
	c.G = uint32(float64(c.G) / 65535 * 255)
	c.B = uint32(float64(c.B) / 65535 * 255)
}