package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var inputFilename, inputURL string

func init() {
	inputFilenameFlag := flag.String("i", "", "Input file to display")
	inputURLFlag := flag.String("u", "", "Input URL to retrieve")
	flag.Parse()
	inputFilename = *inputFilenameFlag
	inputURL = *inputURLFlag

	if inputFilename == "" && inputURL == "" {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func getTrueColorEscapeString(upper, lower Color) string {
	upper.NormalizeValue()
	lower.NormalizeValue()
	return fmt.Sprintf("\033[38;2;%v;%v;%vm\033[48;2;%v;%v;%vm▀\033[0m",
		upper.R, upper.G, upper.B, lower.R, lower.G, lower.B)
}

func traverseImage(image Image) {
	image.ConvertToRGB()
	rgbImage := image.RGB
	var upper, lower Color

	for y := 0; y < image.height; y += 2 {
		for x := 0; x < image.width; x++ {
			upper = NewColor(rgbImage[y][x][0], rgbImage[y][x][1], rgbImage[y][x][2])
			lower = NewColor(rgbImage[y+1][x][0], rgbImage[y+1][x][1], rgbImage[y+1][x][2])
			fmt.Printf("%v", getTrueColorEscapeString(upper, lower))
		}
		fmt.Print("\n")
	}
}

func main() {
	var data io.Reader
	var err error

	if inputFilename != "" {
		data, err = os.Open(inputFilename)
		if err != nil {
			fmt.Println("Invalid file input")
			os.Exit(1)
		}
	} else if inputURL != "" {
		data, err = GetImageFromHTTP(inputURL)
		if err != nil {
			fmt.Println("Network error")
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}

	image, err := LoadImage(data)
	if err != nil {
		fmt.Println("Must provide valid input")
		os.Exit(1)
	}

	traverseImage(*image)
}
