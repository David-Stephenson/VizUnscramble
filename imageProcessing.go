package main

import (
	"image"
	"image/draw"
	"strconv"
)

// This function crops a part of the image and draws it on a new image
func cropAndDraw(img image.Image, newImage *image.RGBA, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationEndX, destinationEndY int) {
	// Define the rectangle to crop
	rect := image.Rect(cropStartX, cropStartY, cropEndX, cropEndY)
	// Crop the image
	croppedImage := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)
	// Draw the cropped image on the new image
	draw.Draw(newImage, image.Rect(destinationStartX, destinationStartY, destinationEndX, destinationEndY), croppedImage, image.Point{X: cropStartX, Y: cropStartY}, draw.Src)
}

// This function unscrambles an image using a key
func unscramble(img image.Image, key []string, width int, height int) *image.RGBA {
	// Define the size of each tile
	tileWidth := width / 10
	tileHeight := height / 15

	// Create a new image
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Crop top
	cropStartX := 0
	cropStartY := 0
	cropEndX := width
	cropEndY := tileHeight
	destinationStartX := 0
	destinationStartY := 0
	cropAndDraw(img, newImage, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationStartX+(cropEndX-cropStartX), destinationStartY+(cropEndY-cropStartY))

	// Crop left
	cropStartX = 0
	cropStartY = tileHeight + 10
	cropEndX = cropStartX + tileWidth
	cropEndY = cropStartY + height - (2 * tileHeight)
	destinationStartX = 0
	destinationStartY = tileHeight
	cropAndDraw(img, newImage, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationStartX+(cropEndX-cropStartX), destinationStartY+(cropEndY-cropStartY))

	// Crop right
	cropStartX = 9 * (tileWidth + 10)
	cropStartY = tileHeight + 10
	cropEndX = cropStartX + tileWidth + (width - 10*tileWidth)
	cropEndY = cropStartY + height - 2*tileHeight
	destinationStartX = 9 * tileWidth
	destinationStartY = tileHeight
	cropAndDraw(img, newImage, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationStartX+(cropEndX-cropStartX), destinationStartY+(cropEndY-cropStartY))

	// Crop bottom
	cropStartX = 0
	cropStartY = 14 * (tileHeight + 10)
	cropEndX = cropStartX + width
	cropEndY = cropStartY + img.Bounds().Dy() - 14*(tileHeight+10)
	destinationStartX = 0
	destinationStartY = 14 * tileHeight
	cropAndDraw(img, newImage, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationStartX+(cropEndX-cropStartX), destinationStartY+(cropEndY-cropStartY))

	// Loop through all segments in the key
	for i := 0; i < len(key); i++ {
		// Convert the current key segment from hexadecimal to an integer
		hexValue, _ := strconv.ParseInt(key[i], 16, 64)

		// Define the rectangle to crop
		// Define the destination rectangle on the new image
		cropStartX := ((i % 8) + 1) * (tileWidth + 10)
		cropStartY := ((i / 8) + 1) * (tileHeight + 10)
		cropEndX := cropStartX + tileWidth
		cropEndY := cropStartY + tileHeight

		destinationStartX := ((int(hexValue) % 8) + 1) * tileWidth
		destinationStartY := ((int(hexValue) / 8) + 1) * tileHeight
		destinationEndX := destinationStartX + tileWidth
		destinationEndY := destinationStartY + tileHeight

		// Crop the source cell out of the temporary image and paste it into the destination cell on the canvas
		cropAndDraw(img, newImage, cropStartX, cropStartY, cropEndX, cropEndY, destinationStartX, destinationStartY, destinationEndX, destinationEndY)
	}

	return newImage
}
