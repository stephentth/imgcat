package main

// Color represent a RGB and RGBA color with specific number of decimal number of bit
type Color struct {
	R, G, B, A uint32
	bit        uint8
	isAlpha    bool
}

// New16BitRGBAColor return new Color instance store 16-bit RGB decimal
func New16BitRGBAColor(r, g, b, a uint32) Color {
	return Color{
		R:       r,
		G:       g,
		B:       b,
		A:       0,
		bit:     16,
		isAlpha: true,
	}
}

// NewEmptyColor retun a black color
func NewEmptyColor() Color {
	return Color{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}
}

// New2DColorArray return 2D array of Color
func New2DColorArray(height, width int) [][]Color {
	image := make([][]Color, height)

	for i := range image {
		// image[i] = make([]Color, width)
		image[i] = make([]Color, width)
		for j := range image[i] {
			image[i][j] = NewEmptyColor()
		}
	}
	return image
}

// Convert8Bit convert current Color into 8bit
func (c *Color) Convert8Bit() {
	if c.bit == 8 {
		return
	}
	c.bit = 8
	c.R = uint32(float64(c.R) / 65535 * 255)
	c.G = uint32(float64(c.G) / 65535 * 255)
	c.B = uint32(float64(c.B) / 65535 * 255)
}
