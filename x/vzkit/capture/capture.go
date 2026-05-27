package capture

import (
	"image"

	base "github.com/tmc/apple/x/capture"
)

// GoImageFromCGImage converts a CGImageRef to a Go image.Image.
var GoImageFromCGImage = base.GoImageFromCGImage

// GenerateDiff creates a diff image highlighting changes between two images.
func GenerateDiff(old, new image.Image) image.Image { return base.GenerateDiff(old, new) }

// ScaleImage resizes an image by scale using bilinear interpolation.
func ScaleImage(img image.Image, scale float64) image.Image { return base.ScaleImage(img, scale) }

// EncodeJPEG encodes img as JPEG with quality.
var EncodeJPEG = base.EncodeJPEG

// EncodePNG encodes img as PNG.
var EncodePNG = base.EncodePNG
