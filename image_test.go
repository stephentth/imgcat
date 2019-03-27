package main

import (
	"os"
	"testing"
)

func TestNewEmptyImage(t *testing.T) {
	image := NewEmptyImage(3, 3)
	if image.height != 4 || image.width != 3 {
		t.Fatal()
	}
	if len(image.RGBA) != 4 || len(image.RGBA[0]) != 3 {
		t.Fatal()
	}
	if len(image.RGB) != 4 || len(image.RGB[0]) != 3 {
		t.Fatal()
	}
}

func TestSmokeLoadImage(t *testing.T) {
	files := []string{"test/artifacts/4.jpg", "test/artifacts/9.jpg"}
	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			t.Fatal("test artifact open err")
		}

		image, err := LoadImage(file)
		if err != nil {
			t.Fatal(err)
		}
		image.Render()
	}
}
