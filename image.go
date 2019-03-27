package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

// Image modeling a color image (in RGB and RGBA)
type Image struct {
	height, width int
	RGBA          [][]Color
	RGB           [][]Color
}

// NewEmptyImage return an empty (all black) Image instance with height and width
func NewEmptyImage(height, width int) *Image {
	if height%2 == 1 {
		height++
	}
	image := &Image{
		height: height,
		width:  width,
	}
	image.RGBA = New2DColorArray(height, width)
	image.RGB = New2DColorArray(height, width)
	return image
}

// LoadImage get a Image from io.Reader
func LoadImage(file io.Reader) (*Image, error) {
	imgData, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	var r, g, b, a uint32
	bounds := imgData.Bounds()
	height := bounds.Max.Y
	width := bounds.Max.X
	image := NewEmptyImage(height, width)

	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x < width; x++ {
			r, g, b, a = imgData.At(x, y).RGBA()
			image.RGBA[y][x] = New16BitRGBAColor(r, g, b, a)
		}
	}
	return image, nil
}

// ConvertToRGB convert this Image to RGB color code
func (i *Image) ConvertToRGB() {
	for y := range i.RGBA {
		for x := range i.RGBA[y] {
			r, g, b, a := i.RGBA[y][x].R, i.RGBA[y][x].G, i.RGBA[y][x].B, i.RGBA[y][x].A
			_ = a

			i.RGB[y][x].R = r
			i.RGB[y][x].G = g
			i.RGB[y][x].B = b
		}
	}
}

// Convert8Bit convert this Image to 8bit color code
func (i *Image) Convert8Bit() {
	for y := range i.RGBA {
		for x := range i.RGBA[y] {
			i.RGBA[y][x].Convert8Bit()
			i.RGB[y][x].Convert8Bit()
		}
	}
}

// Render this Image to ANSI color code and output to stdout
func (image Image) Render() error {
	colSize, _, err := GetTerminalSize()
	if err != nil {
		return err
	}
	if image.width > colSize {
		newHeight := int(float64(colSize) / float64(image.width) * float64(image.height))
		image = NearestNeighborResampling(image, newHeight, colSize)
	}
	image.ConvertToRGB()
	image.Convert8Bit()
	rgbImage := image.RGB
	var upper, lower Color

	for y := 0; y < image.height; y += 2 {
		for x := 0; x < image.width; x++ {
			upper = New16BitRGBAColor(rgbImage[y][x].R, rgbImage[y][x].G, rgbImage[y][x].B, rgbImage[y][x].A)
			lower = New16BitRGBAColor(rgbImage[y+1][x].R, rgbImage[y+1][x].G, rgbImage[y+1][x].B, rgbImage[y+1][x].A)
			fmt.Printf("%v", getTrueColorEscapeString(upper, lower))
		}
		fmt.Print("\n")
	}
	return nil
}

// getTrueColorEscapeString convert 2 pixel color to ANSI sequence code
func getTrueColorEscapeString(upper, lower Color) string {
	return fmt.Sprintf("\033[38;2;%v;%v;%vm\033[48;2;%v;%v;%vm▀\033[0m",
		upper.R, upper.G, upper.B, lower.R, lower.G, lower.B)
}
