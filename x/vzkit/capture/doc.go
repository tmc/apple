// Package capture provides image conversion and processing utilities
// for VM screenshot workflows.
//
// It converts CGImage data (BGRA pixel format) to Go [image.Image] values,
// generates visual diffs between successive frames, and encodes images
// to JPEG or PNG.
//
// Basic usage:
//
//	img, err := capture.GoImageFromCGImage(cgImage, 0)
//	data, err := capture.EncodeJPEG(img, 80)
package capture
