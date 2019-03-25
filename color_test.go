package main

import "testing"

func TestColorCreateNormalize(t *testing.T) {
	color := NewColor(65535, 65535, 65535) // white
	color.NormalizeValue()
	if color.R != 255 || color.G != 255 || color.B != 255 {
		t.Fatalf("color mismatch")
	}
}
