package main

import "testing"

func TestColorCreateNormalize(t *testing.T) {
	color := New16BitRGBAColor(65535, 65535, 65535, 65535) // white
	color.Convert8Bit()
	if color.R != 255 || color.G != 255 || color.B != 255 {
		t.Fatalf("color mismatch")
	}
}
