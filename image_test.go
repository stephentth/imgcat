package main

import (
	"os"
	"testing"
)

func TestMake3DArray(t *testing.T) {
	array := make3DArray(3, 3, 3)
	if array[2][2][2] != 0 {
		t.Fatal()
	}
}

func TestCreateImage(t *testing.T) {
	image := NewImage(99, 100)
	if image.height != 100 || len(image.RGBA[99][99]) != 4 {
		t.Fatal()
	}
}

func TestLoadImage(t *testing.T) {
	file, err := os.Open("test/artifacts/drawer.jpg")
	if err != nil {
		t.Fatal("test artifact open err")
	}

	image, err := LoadImage(file)
	if err != nil {
		t.Fatal(err)
	}
	_ = image
}
