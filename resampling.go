package main

// NearestNeighborResampling using Nearest Neighbor algorithm to scale the image
func NearestNeighborResampling(inputImage Image, height, width int) Image {
	outputImage := NewEmptyImage(height, width)
	inputHeight := inputImage.height
	inputWidth := inputImage.width

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			yNearest := int(float64(i) / float64(height) * float64(inputHeight))
			xNearest := int(float64(j) / float64(width) * float64(inputWidth))
			outputImage.RGBA[i][j] = inputImage.RGBA[yNearest][xNearest]
		}
	}
	return *outputImage
}
