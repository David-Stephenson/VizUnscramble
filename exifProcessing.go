package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"strconv"
	"strings"
)

// ImageData struct to hold the parsed exif data
type ImageData struct {
	KeySegments []string
	Width       int
	Height      int
}

// parseExifValues function takes a reader and returns parsed exif data
func parseExifValues(rawData io.Reader) (*ImageData, error) {
	// Decode the exif data from the reader
	data, err := exif.Decode(rawData)
	if err != nil {
		// Return nil and the error if decoding fails
		return nil, fmt.Errorf("error decoding exif data: %v", err)
	}

	// Get the ImageUniqueID from the exif data
	imageUniqueID, err := data.Get(exif.ImageUniqueID)
	if err != nil {
		// Return nil and the error if ImageUniqueID is not found
		return nil, fmt.Errorf("ImageUniqueID not found: %v", err)
	}

	// Get the ImageWidth from the exif data
	imageWidth, err := data.Get(exif.ImageWidth)
	if err != nil {
		// Return nil and the error if ImageWidth is not found
		return nil, fmt.Errorf("ImageWidth not found: %v", err)
	}

	// Get the ImageLength from the exif data
	imageLength, err := data.Get(exif.ImageLength)
	if err != nil {
		// Return nil and the error if ImageLength is not found
		return nil, fmt.Errorf("ImageLength not found: %v", err)
	}

	// Remove the quotes from the ImageUniqueID and split it into segments
	key := strings.Trim(imageUniqueID.String(), "\"")
	keySegments := strings.Split(key, ":")

	// Convert the ImageWidth from string to int
	width, err := strconv.Atoi(imageWidth.String())
	if err != nil {
		// Return nil and the error if the conversion fails
		return nil, fmt.Errorf("error converting width to int: %v", err)
	}

	// Convert the ImageLength from string to int
	height, err := strconv.Atoi(imageLength.String())
	if err != nil {
		// Return nil and the error if the conversion fails
		return nil, fmt.Errorf("error converting height to int: %v", err)
	}

	// Return the parsed exif data
	return &ImageData{KeySegments: keySegments, Width: width, Height: height}, nil
}
