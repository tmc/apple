package capture

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"golang.org/x/image/draw"
)

// GoImageFromCGImage converts a CGImageRef to a Go image.Image.
// CGImage uses BGRA pixel format; this function converts to RGBA.
// If cropTopPx > 0, that many pixels are cropped from the top
// (useful for removing window title bars from screenshots).
func GoImageFromCGImage(cgImage coregraphics.CGImageRef, cropTopPx int) (image.Image, error) {
	if cgImage == 0 {
		return nil, fmt.Errorf("nil cgimage")
	}

	width := int(coregraphics.CGImageGetWidth(cgImage))
	height := int(coregraphics.CGImageGetHeight(cgImage))
	if width == 0 || height == 0 {
		return nil, fmt.Errorf("invalid image dimensions: %dx%d", width, height)
	}

	dataProvider := coregraphics.CGImageGetDataProvider(cgImage)
	if dataProvider == 0 {
		return nil, fmt.Errorf("get data provider")
	}

	cfData := coregraphics.CGDataProviderCopyData(dataProvider)
	if cfData == 0 {
		return nil, fmt.Errorf("copy data from provider")
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cfData))

	dataLength := corefoundation.CFDataGetLength(cfData)
	dataPtr := corefoundation.CFDataGetBytePtr(cfData)
	if dataPtr == nil || dataLength == 0 {
		return nil, fmt.Errorf("get data bytes")
	}

	bytesPerRow := int(coregraphics.CGImageGetBytesPerRow(cgImage))
	contentHeight := height - cropTopPx
	if contentHeight <= 0 {
		return nil, fmt.Errorf("invalid cropped height: %d", contentHeight)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, width, contentHeight))
	srcData := unsafe.Slice((*byte)(dataPtr), dataLength)
	dstData := rgba.Pix

	for y := 0; y < contentHeight; y++ {
		srcRowStart := (y + cropTopPx) * bytesPerRow
		dstRowStart := y * rgba.Stride
		for x := 0; x < width; x++ {
			srcPixel := srcRowStart + x*4
			if srcPixel+3 < int(dataLength) {
				dstPixel := dstRowStart + x*4
				// BGRA to RGBA
				dstData[dstPixel] = srcData[srcPixel+2]   // R
				dstData[dstPixel+1] = srcData[srcPixel+1] // G
				dstData[dstPixel+2] = srcData[srcPixel]   // B
				dstData[dstPixel+3] = srcData[srcPixel+3] // A
			}
		}
	}

	return rgba, nil
}

// GenerateDiff creates a diff image highlighting changes between two images.
// Changed pixels are shown in full color; unchanged pixels are shown as
// dimmed grayscale. If the dimensions differ, the new image is returned
// with a red border.
func GenerateDiff(old, new image.Image) image.Image {
	oldBounds := old.Bounds()
	newBounds := new.Bounds()

	width := newBounds.Dx()
	height := newBounds.Dy()

	diff := image.NewRGBA(image.Rect(0, 0, width, height))

	if oldBounds.Dx() != width || oldBounds.Dy() != height {
		draw.Copy(diff, image.Point{}, new, newBounds, draw.Src, nil)
		red := color.RGBA{255, 0, 0, 255}
		for x := 0; x < width; x++ {
			diff.SetRGBA(x, 0, red)
			diff.SetRGBA(x, height-1, red)
		}
		for y := 0; y < height; y++ {
			diff.SetRGBA(0, y, red)
			diff.SetRGBA(width-1, y, red)
		}
		return diff
	}

	threshold := uint32(10 * 256)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldR, oldG, oldB, _ := old.At(x+oldBounds.Min.X, y+oldBounds.Min.Y).RGBA()
			newR, newG, newB, newA := new.At(x+newBounds.Min.X, y+newBounds.Min.Y).RGBA()

			dr := absDiff(oldR, newR)
			dg := absDiff(oldG, newG)
			db := absDiff(oldB, newB)

			if dr > threshold || dg > threshold || db > threshold {
				diff.SetRGBA(x, y, color.RGBA{
					R: uint8(newR >> 8),
					G: uint8(newG >> 8),
					B: uint8(newB >> 8),
					A: uint8(newA >> 8),
				})
			} else {
				gray := uint8((newR + newG + newB) / 3 >> 8)
				diff.SetRGBA(x, y, color.RGBA{gray / 3, gray / 3, gray / 3, 255})
			}
		}
	}

	return diff
}

// ScaleImage resizes an image by the given scale factor using bilinear interpolation.
func ScaleImage(img image.Image, scale float64) image.Image {
	bounds := img.Bounds()
	newWidth := int(float64(bounds.Dx()) * scale)
	newHeight := int(float64(bounds.Dy()) * scale)

	if newWidth < 1 {
		newWidth = 1
	}
	if newHeight < 1 {
		newHeight = 1
	}

	scaled := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(scaled, scaled.Bounds(), img, bounds, draw.Over, nil)
	return scaled
}

// EncodeJPEG encodes an image as JPEG with the given quality (1-100).
func EncodeJPEG(img image.Image, quality int) ([]byte, error) {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, fmt.Errorf("encode jpeg: %w", err)
	}
	return buf.Bytes(), nil
}

// EncodePNG encodes an image as PNG.
func EncodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}
	return buf.Bytes(), nil
}

func absDiff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}
