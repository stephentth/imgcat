package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"io"
)

type Image struct {
	height, width int
	RGBA          [][][]uint32
	RGB           [][][]uint32
}

func NewImage(height, width int) *Image {
	image := &Image{
		height: height,
		width:  width,
	}
	if height%2 == 1 {
		image.height++
	}
	image.RGBA = make3DArray(height, width, 4)
	if height%2 == 1 {
		image.RGBA[image.height-1] = make([][]uint32, width)
		for i := range image.RGBA[image.height-1] {
			image.RGBA[image.height-1][i] = make([]uint32, 4)
		}
	}
	return image
}

func (i *Image) ConvertToRGB() {
	i.RGB = make3DArray(i.height, i.width, 3)
	for y := range i.RGBA {
		for x := range i.RGBA[y] {
			r, g, b, _ := i.RGBA[y][x][0], i.RGBA[y][x][1], i.RGBA[y][x][2], i.RGBA[y][x][3]

			i.RGB[y][x][0] = r
			i.RGB[y][x][1] = g
			i.RGB[y][x][2] = b
		}
	}
}

func make3DArray(n, m, d int) (matrix [][][]uint32) {
	matrix = make([][][]uint32, n)

	for i := range matrix {
		matrix[i] = make([][]uint32, m)
		for j := range matrix[i] {
			matrix[i][j] = make([]uint32, d)
		}
	}
	return
}

func LoadImage(file io.Reader) (*Image, error) {
	imgData, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	var r, g, b, a uint32
	bounds := imgData.Bounds()
	height := bounds.Max.Y
	width := bounds.Max.X
	image := NewImage(height, width)

	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x < width; x++ {
			r, g, b, a = imgData.At(x, y).RGBA()
			image.RGBA[y][x] = []uint32{r, g, b, a}
		}
	}
	return image, nil
}
