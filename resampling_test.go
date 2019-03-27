package main

import (
	"testing"
)

func TestNearestNeighborResampling(t *testing.T) {
	image := NewEmptyImage(10, 10)
	output := NearestNeighborResampling(*image, 2, 2)
	output.Render()
	output = NearestNeighborResampling(*image, 3, 3)
	output.Render()
}
