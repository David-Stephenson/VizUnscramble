package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// openFile is a helper that opens a file. If it can't open the file, it tells us why.
func openFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// saveFile takes a picture and a name, then tries to save the picture with that name.
// If it can't make the file or save the picture, it tells us why.
func saveFile(name string, img image.Image) error {
	// Try to make a new file with the given name
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	// Make sure we close the file when we're done, even if something goes wrong
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Oops, trouble closing the file: %v\n", err)
		}
	}()

	// Try to save the picture to the file with the best quality
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		return err
	}

	return nil
}
