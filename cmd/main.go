package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kashifkhan0771/cli-shapes/shapes"
)

type Shape int

const (
	shapeSize = 5

	Triangle Shape = iota
	RightTriangle
	Diamond
	ChillGuy
	Butterfly
	RedFlag
	Circle
)

// ShapeStrings maps the enum to its string representation.
var ShapeStrings = map[Shape]string{
	Triangle:      "triangle",
	RightTriangle: "right-triangle",
	Diamond:       "diamond",
	ChillGuy:      "chill-guy",
	Butterfly:     "butterfly",
	RedFlag:       "red-flag",
	Circle:        "circle",
}

// StringShapes maps the string to its shape representation.
var StringShapes = map[string]Shape{
	"triangle":       Triangle,
	"right-triangle": RightTriangle,
	"diamond":        Diamond,
	"chill-guy":      ChillGuy,
	"butterfly":      Butterfly,
	"red-flag":       RedFlag,
	"circle":         Circle,
}

func main() {
	shapeFlag := flag.String("shape", "", "Shape to print")
	shapeSizeFlag := flag.Int("size", shapeSize, "Size of the shape")
	helpFlag := flag.Bool("help", false, "List available shapes")

	flag.Parse()

	if *helpFlag {
		printAvailableShapes()
	}

	shapeName := *shapeFlag
	shapeSize := *shapeSizeFlag

	if shapeName == "" {
		fmt.Print("Enter shape name: ")
		_, err := fmt.Scanln(&shapeName)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
	}

	if !shapeExist(shapeName) {
		fmt.Println("Invalid Shape name. Shape not found. Run with --help to see available shapes.")
		os.Exit(1)
	}

	switch StringShapes[shapeName] {
	case Triangle:
		shapes.EquilateralTriangle(shapeSize)
	case RightTriangle:
		shapes.RightTriangle(shapeSize)
	case Diamond:
		shapes.Diamond(shapeSize)
	case ChillGuy:
		shapes.ChillGuy()
	case Butterfly:
		shapes.Butterfly()
	case RedFlag:
		shapes.RedFlag()
	case Circle:
		shapes.Circle(shapeSize)
	}
}

func printAvailableShapes() {
	fmt.Println("Available shapes:")
	for _, shape := range ShapeStrings {
		fmt.Println("-", shape)
	}

	os.Exit(0)
}

func shapeExist(shapeName string) bool {
	if _, exist := StringShapes[shapeName]; exist {
		return true
	}

	return false
}
