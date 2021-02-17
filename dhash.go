/*
* Inspired by Dr. Neal Krawetz
* http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html
 */

package dhash

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

// reduceSizeAndColor reducing size of image to 9x8 pixels
// and also scales for each pixel RGB values between 0-255 equally.
func reduceSizeAndColor(path string) *image.NRGBA {
	input, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	input = imaging.Grayscale(input)
	output := imaging.Resize(input, 9, 8, imaging.Lanczos)
	return output
}

// createByteArray parses the Pix attribute of *image.NRGBA and
// creates a new byte array with Red value while ditching Green, Blue and Alpha values.
func createByteArray(image *image.NRGBA) []byte {
	byteArray := make([]byte, 0)
	for y := 0; y < 8; y++ {
		for x := 0; x < 9; x++ {
			r, _, _, _ := image.At(x, y).RGBA()
			byteArray = append(byteArray, byte(r))
		}
	}
	return byteArray
}

// horizontalGradientHash calculates the second half of the hash value by
// comparing the pixels horizontally.
// For example; If 0x0 pixels brighter than 1x0 it returns 1, else 0.
func horizontalGradientHash(imageArray []uint8) []uint8 {
	var previousValue uint8
	hashArray := make([]uint8, 0)
	for y := 0; y < 8; y++ {
		for x := 0; x < 9; x++ {
			currentValue := imageArray[(y*8)+(x)]
			if x > 0 {
				if previousValue < currentValue {
					hashArray = append(hashArray, 1)
				} else {
					hashArray = append(hashArray, 0)
				}
			}
			previousValue = currentValue
		}
	}
	return hashArray
}

// verticalGradientHash calculates the first half of the hash value by
// comparing the pixels vertically.
// For example; If 0x0 pixels brighter than 0x1 it returns 1, else 0.
func verticalGradientHash(imageArray []uint8) []uint8 {
	var previousValue uint8
	hashArray := make([]uint8, 0)
	for x := 0; x < 8; x++ {
		for y := 0; y < 9; y++ {
			currentValue := imageArray[(y*8)+(x)]
			if y > 0 {
				if previousValue < currentValue {
					hashArray = append(hashArray, 1)
				} else {
					hashArray = append(hashArray, 0)
				}
			}
			previousValue = currentValue
		}
	}
	return hashArray
}

// CalculateHash calculates the dhash value for given image.
func CalculateHash(path string) []uint8 {
	imageArray := createByteArray(reduceSizeAndColor(path))
	verticalHash := verticalGradientHash(imageArray)
	horizontalHash := horizontalGradientHash(imageArray)
	return append(verticalHash, horizontalHash...)
}

// CalculateHammingDistance calculates the hamming distance between two 128-bit hash values.
func CalculateHammingDistance(firstHash []uint8, secondHash []uint8) float64 {
	similarity := 0

	for i := 0; i < 128; i++ {
		if firstHash[i] == secondHash[i] {
			similarity++
		}
	}

	return float64(similarity) / 128
}
