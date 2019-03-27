package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// parseInput
// imgcat can receive 2 type of input (can receive in mix), filename of local disk image, or url of an image
// this function get all of the parammeters and return 2 array
func parseInput() ([]string, []string) {
	flag.Parse()
	var fileInput, urlInput []string

	if flag.NArg() == 0 {
		fmt.Println("Please provide inputs")
		os.Exit(1)
	}

	for _, item := range flag.Args() {
		if IsValidUrl(item) {
			urlInput = append(urlInput, item)
		} else if IsValidFile(item) {
			fileInput = append(fileInput, item)
		} else {
			fmt.Printf("%v is not a valid input (neigher valid file nor url)\n", item)
			os.Exit(1)
		}
	}
	return fileInput, urlInput
}

func main() {
	fileInput, urlInput := parseInput()
	var data io.Reader
	var err error

	for _, filename := range fileInput {
		data, err = os.Open(filename)
		if err != nil {
			fmt.Println("Invalid file input")
			os.Exit(1)
		}

		image, err := LoadImage(data)
		if err != nil {
			fmt.Printf("File %v can not be load.\n", filename)
			os.Exit(1)
		}
		image.Render()
	}

	for _, url := range urlInput {
		data, err = GetImageFromHTTP(url)
		if err != nil {
			fmt.Println("Network error")
			os.Exit(1)
		}

		image, err := LoadImage(data)
		if err != nil {
			fmt.Printf("URL %v can not be load.\n", url)
			os.Exit(1)
		}
		image.Render()
	}
}
