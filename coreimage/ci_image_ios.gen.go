// Code generated from Apple documentation for CoreImage. DO NOT EDIT.
//go:build ios
// +build ios

package coreimage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Initializes an image object with the specified UIKit image object.
//
// image: An image containing the source data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:)
func (i CIImage) InitWithImage(image objectivec.IObject) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithImage:"), image)
	return rv
}

// Initializes an image object with the specified UIKit image object, using
// the specified options.
//
// image: An image containing the source data.
//
// options: A dictionary that contains options for creating an image object. You can
// supply such options as a pixel format and a color space. See `Image
// Dictionary Keys`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:options:)
func (i CIImage) InitWithImageOptions(image objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithImage:options:"), image, options)
	return rv
}
