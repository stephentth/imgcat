package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"
)

// TotalNumberPokemon number of pokemon in PokemonDB.net
const TotalNumberPokemon = 809

// PokemonAPIEndpoint online db to query pokemon name from ID
const PokemonAPIEndpoint = "https://pokeapi.co/api/v2/pokemon/"

// PokemonImageSchema url format to load pokemon image from name
const PokemonImageSchema = "https://img.pokemondb.net/artwork/%v.jpg"

var inputFilename, inputURL string
var inputRandom bool

type color struct {
	R, G, B uint32
}

func newColor(r, g, b uint32) color {
	return color{r, g, b}
}

func (c *color) normalizeValue() {
	c.R = uint32(float64(c.R) / 65535 * 255)
	c.G = uint32(float64(c.G) / 65535 * 255)
	c.B = uint32(float64(c.B) / 65535 * 255)
}

func init() {
	inputFilenameFlag := flag.String("i", "", "Input file to display")
	inputURLFlag := flag.String("u", "", "Input URL to retrieve")
	inputRandomFlag := flag.Bool("r", false, "Choose random pokemon")
	flag.Parse()
	inputFilename = *inputFilenameFlag
	inputURL = *inputURLFlag
	inputRandom = *inputRandomFlag

	if inputFilename == "" && inputURL == "" && inputRandom == false {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func getTrueColorEscapeString(upper, lower color) string {
	upper.normalizeValue()
	lower.normalizeValue()
	return fmt.Sprintf("\033[38;2;%v;%v;%vm\033[48;2;%v;%v;%vm▀\033[0m",
		upper.R, upper.G, upper.B, lower.R, lower.G, lower.B)
}

func traverseImage(image image.Image) {
	var r, g, b uint32
	bounds := image.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ = image.At(x, y).RGBA()
			upper := newColor(r, g, b)
			r, g, b, _ = image.At(x, y+1).RGBA()
			lower := newColor(r, g, b)
			fmt.Printf("%v", getTrueColorEscapeString(upper, lower))
		}
		fmt.Print("\n")
	}
}

func main() {
	var data io.Reader
	var err error

	if inputRandom == true {
		pokemonName, err := GetRandomPokemon(0)
		if err != nil {
			fmt.Println("Network error")
			os.Exit(1)
		}
		fmt.Println("Got pokemon", pokemonName)
		data, err = GetImageFromHTTP(fmt.Sprintf(PokemonImageSchema, pokemonName))
		if err != nil {
			fmt.Println("Network error")
			os.Exit(1)
		}
	} else if inputFilename != "" {
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

	imgData, _, err := image.Decode(data)
	if err != nil {
		fmt.Println("Must provide valid input")
		os.Exit(1)
	}

	traverseImage(imgData)
}
