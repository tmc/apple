// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataMachineReadableCodeObject] class.
var (
	_AVMetadataMachineReadableCodeObjectClass     AVMetadataMachineReadableCodeObjectClass
	_AVMetadataMachineReadableCodeObjectClassOnce sync.Once
)

func getAVMetadataMachineReadableCodeObjectClass() AVMetadataMachineReadableCodeObjectClass {
	_AVMetadataMachineReadableCodeObjectClassOnce.Do(func() {
		_AVMetadataMachineReadableCodeObjectClass = AVMetadataMachineReadableCodeObjectClass{class: objc.GetClass("AVMetadataMachineReadableCodeObject")}
	})
	return _AVMetadataMachineReadableCodeObjectClass
}

// GetAVMetadataMachineReadableCodeObjectClass returns the class object for AVMetadataMachineReadableCodeObject.
func GetAVMetadataMachineReadableCodeObjectClass() AVMetadataMachineReadableCodeObjectClass {
	return getAVMetadataMachineReadableCodeObjectClass()
}

type AVMetadataMachineReadableCodeObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataMachineReadableCodeObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataMachineReadableCodeObjectClass) Alloc() AVMetadataMachineReadableCodeObject {
	rv := objc.Send[AVMetadataMachineReadableCodeObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Barcode information detected by a metadata capture output.
//
// # Overview
//
// The [AVMetadataMachineReadableCodeObject] class is a concrete subclass of
// [AVMetadataObject] defining the features of a detected one-dimensional or
// two-dimensional barcode.
//
// An [AVMetadataMachineReadableCodeObject] instance represents a single
// detected machine readable code in an image. It’s an immutable object
// describing the features and payload of a barcode.
//
// On supported platforms, the [AVCaptureMetadataOutput] class outputs arrays
// of detected machine readable code objects.
//
// # Getting machine-readable code values
//
//   - [AVMetadataMachineReadableCodeObject.Descriptor]: A barcode description for use in Core Image.
//   - [AVMetadataMachineReadableCodeObject.StringValue]: Returns the error-corrected data decoded into a human-readable string.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject
type AVMetadataMachineReadableCodeObject struct {
	AVMetadataObject
}

// AVMetadataMachineReadableCodeObjectFromID constructs a [AVMetadataMachineReadableCodeObject] from an objc.ID.
//
// Barcode information detected by a metadata capture output.
func AVMetadataMachineReadableCodeObjectFromID(id objc.ID) AVMetadataMachineReadableCodeObject {
	return AVMetadataMachineReadableCodeObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}

// NOTE: AVMetadataMachineReadableCodeObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataMachineReadableCodeObject] class.
//
// # Getting machine-readable code values
//
//   - [IAVMetadataMachineReadableCodeObject.Descriptor]: A barcode description for use in Core Image.
//   - [IAVMetadataMachineReadableCodeObject.StringValue]: Returns the error-corrected data decoded into a human-readable string.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject
type IAVMetadataMachineReadableCodeObject interface {
	IAVMetadataObject

	// Topic: Getting machine-readable code values

	// A barcode description for use in Core Image.
	Descriptor() coreimage.CIBarcodeDescriptor
	// Returns the error-corrected data decoded into a human-readable string.
	StringValue() string

	// The points defining the (x, 	y) locations of the corners.
	Corners() foundation.INSDictionary
}

// Init initializes the instance.
func (m AVMetadataMachineReadableCodeObject) Init() AVMetadataMachineReadableCodeObject {
	rv := objc.Send[AVMetadataMachineReadableCodeObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataMachineReadableCodeObject) Autorelease() AVMetadataMachineReadableCodeObject {
	rv := objc.Send[AVMetadataMachineReadableCodeObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataMachineReadableCodeObject creates a new AVMetadataMachineReadableCodeObject instance.
func NewAVMetadataMachineReadableCodeObject() AVMetadataMachineReadableCodeObject {
	class := getAVMetadataMachineReadableCodeObjectClass()
	rv := objc.Send[AVMetadataMachineReadableCodeObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A barcode description for use in Core Image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject/descriptor
func (m AVMetadataMachineReadableCodeObject) Descriptor() coreimage.CIBarcodeDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("descriptor"))
	return coreimage.CIBarcodeDescriptorFromID(objc.ID(rv))
}

// Returns the error-corrected data decoded into a human-readable string.
//
// # Discussion
//
// The value of this property is an [NSString] created by decoding the binary
// payload according to the format of the machine-readable code, or `nil` if a
// string representation cannot be created.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject/stringValue
func (m AVMetadataMachineReadableCodeObject) StringValue() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}

// The points defining the (x, y) locations of the corners.
//
// # Discussion
//
// The value of this property is an array of [CFDictionary] objects, each of
// which has been created from a [CGPoint] struct using the
// [dictionaryRepresentation] function, representing the coordinates of the
// corners of the object with respect to the image in which it resides.
//
// If the metadata originates from video, the points may be expressed as
// scalar values from `0` to `1`.
//
// The points in the corners differ from the bounds rectangle in that bounds
// is axis aligned to orientation of the captured image, and the values of the
// corners reside within the bounds rectangle.
//
// The points are arranged in counterclockwise order (clockwise if the code or
// image is mirrored), starting with the top left of the code in its canonical
// orientation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject/corners-8f6bv
//
// [dictionaryRepresentation]: https://developer.apple.com/documentation/CoreFoundation/CGPoint/dictionaryRepresentation
func (m AVMetadataMachineReadableCodeObject) Corners() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("corners"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
