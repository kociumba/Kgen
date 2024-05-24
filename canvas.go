package main

import (
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
)

type Canvas struct {
	img *image.RGBA
}

// NewCanvas creates a new Canvas with the specified width and height
func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		img: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// ColorPixel colors a specific pixel using the provided seed
func (c *Canvas) ColorPixel(seed int64) {
	if c == nil || c.img == nil {
		return
	}

	width, height := c.img.Bounds().Dx(), c.img.Bounds().Dy()

	// Define buffer for image data
	buf := make([]uint8, len(c.img.Pix))
	copy(buf, c.img.Pix)

	getPixelSeed := func(x, y int) int64 {
		return seed * int64(x+y*c.img.Stride)
	}

	// Populate the buffer with pixel colors
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			setColorInBuffer(buf, x, y, width, getPixelSeed(x, y))
		}
	}

	// Increase the number of iterations to elongate patterns
	iterations := 1000

	for iter := 0; iter < iterations; iter++ {
		direction, count := rand.Intn(2), rand.Intn(width*1000) // Increase count for longer patterns

		// Iterate over each pixel in the image buffer
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				// Decrement count and switch direction if count reaches 0
				count--
				if count == 0 {
					count, direction = rand.Intn(width*2), rand.Intn(2)
				}

				// If the direction is right-to-left, reverse the order of pixels in the row in the buffer
				if direction == 1 {
					start, end := y*c.img.Stride+width*4, (y+1)*c.img.Stride
					if start < end && start < len(buf) && end <= len(buf) {
						reversePixelsInBuffer(buf, start, end)
					}
				}
			}
		}
	}

	// Copy the modified buffer back to the image
	copy(c.img.Pix, buf)
}

// setColorInBuffer sets the color of the pixel at (x, y) in the buffer based on the provided seed
func setColorInBuffer(buf []uint8, x, y, width int, seed int64) {
	// Calculate index of the pixel in the buffer
	index := (y * width * 4) + (x * 4)

	// Use seed-based randomness for each color channel
	r := uint8(math.Sin(float64(seed)) * 255)
	g := uint8(math.Cos(float64(seed)) * 255)
	b := uint8(math.Tan(float64(seed)) * 255)

	// Set the color of the pixel in the buffer
	buf[index] = r
	buf[index+1] = g
	buf[index+2] = b
	buf[index+3] = 255 // Alpha channel
}

// reversePixelsInBuffer reverses the order of pixels in the specified range of the buffer
func reversePixelsInBuffer(buf []uint8, start, end int) {
	for i, j := start, end-4; i < j; i, j = i+4, j-4 {
		// Swap each RGBA pixel
		buf[i], buf[j] = buf[j], buf[i]
		buf[i+1], buf[j+1] = buf[j+1], buf[i+1]
		buf[i+2], buf[j+2] = buf[j+2], buf[i+2]
		buf[i+3], buf[j+3] = buf[j+3], buf[i+3]
	}
}

// Save saves the image to disk as a PNG
func (c *Canvas) Save(filename string) error {
	file, err := os.Create(filename + ".png")
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, c.img)
}
