package main

import (
	"fmt"
	"image"
	"os"
)

func main() {
	// File to perform actions on
	inputFile := "input.jpg"

	// Open input image
	file, err := openFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode image
	srcImage, _, _ := image.Decode(file)

	// Reset file read position
	_, _ = file.Seek(0, 0)

	// Get values from image exif
	imageData, err := parseExifValues(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print collected values
	fmt.Println("Key segments:", imageData.KeySegments)
	fmt.Println("Output width:", imageData.Width)
	fmt.Println("Output height: ", imageData.Height)

	newImage := unscramble(srcImage, imageData.KeySegments, imageData.Width, imageData.Height)

	// Save the final image
	outputFile := "output.jpg"
	err = saveFile(outputFile, newImage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
