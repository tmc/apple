// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSPICTImageRep] class.
var (
	_NSPICTImageRepClass     NSPICTImageRepClass
	_NSPICTImageRepClassOnce sync.Once
)

func getNSPICTImageRepClass() NSPICTImageRepClass {
	_NSPICTImageRepClassOnce.Do(func() {
		_NSPICTImageRepClass = NSPICTImageRepClass{class: objc.GetClass("NSPICTImageRep")}
	})
	return _NSPICTImageRepClass
}

// GetNSPICTImageRepClass returns the class object for NSPICTImageRep.
func GetNSPICTImageRepClass() NSPICTImageRepClass {
	return getNSPICTImageRepClass()
}

type NSPICTImageRepClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPICTImageRepClass) Alloc() NSPICTImageRep {
	rv := objc.Send[NSPICTImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that renders an image from a PICT format data stream of version
// 1, version 2, and extended version 2.
//
// # Overview
//
// # Creating Representations of Images from PICT Data
//
//   - [NSPICTImageRep.InitWithData]: Returns a representation of an image from the specified data in the PICT file format.
//
// # Getting Data
//
//   - [NSPICTImageRep.BoundingBox]: The rectangle that bounds the image representation.
//   - [NSPICTImageRep.PICTRepresentation]: The image representation’s PICT data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep
type NSPICTImageRep struct {
	NSImageRep
}

// NSPICTImageRepFromID constructs a [NSPICTImageRep] from an objc.ID.
//
// An object that renders an image from a PICT format data stream of version
// 1, version 2, and extended version 2.
func NSPICTImageRepFromID(id objc.ID) NSPICTImageRep {
	return NSPICTImageRep{NSImageRep: NSImageRepFromID(id)}
}
// NOTE: NSPICTImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPICTImageRep] class.
//
// # Creating Representations of Images from PICT Data
//
//   - [INSPICTImageRep.InitWithData]: Returns a representation of an image from the specified data in the PICT file format.
//
// # Getting Data
//
//   - [INSPICTImageRep.BoundingBox]: The rectangle that bounds the image representation.
//   - [INSPICTImageRep.PICTRepresentation]: The image representation’s PICT data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep
type INSPICTImageRep interface {
	INSImageRep

	// Topic: Creating Representations of Images from PICT Data

	// Returns a representation of an image from the specified data in the PICT file format.
	InitWithData(pictData foundation.INSData) NSPICTImageRep

	// Topic: Getting Data

	// The rectangle that bounds the image representation.
	BoundingBox() corefoundation.CGRect
	// The image representation’s PICT data.
	PICTRepresentation() foundation.INSData
}

// Init initializes the instance.
func (p NSPICTImageRep) Init() NSPICTImageRep {
	rv := objc.Send[NSPICTImageRep](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPICTImageRep) Autorelease() NSPICTImageRep {
	rv := objc.Send[NSPICTImageRep](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPICTImageRep creates a new NSPICTImageRep instance.
func NewNSPICTImageRep() NSPICTImageRep {
	class := getNSPICTImageRepClass()
	rv := objc.Send[NSPICTImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewPICTImageRepWithCoder(coder foundation.INSCoder) NSPICTImageRep {
	instance := getNSPICTImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPICTImageRepFromID(rv)
}

// Returns a representation of an image from the specified data in the PICT
// file format.
//
// pictData: A data object containing the PICT data.
//
// # Return Value
// 
// An initialized [NSPICTImageRep], or `nil` if the object could not be
// initialized. Initialization may fail if the data does not conform to the
// PICT file format.
//
// # Discussion
// 
// If the PICT data is obtained directly from a PICT file or document, this
// method ignores most of the 512-byte header that occurs before the start of
// the actual picture data. It may retrieve some relevant meta information
// from the header.
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/init(data:)
func NewPICTImageRepWithData(pictData foundation.INSData) NSPICTImageRep {
	instance := getNSPICTImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), pictData)
	return NSPICTImageRepFromID(rv)
}

// Returns a representation of an image from the specified data in the PICT
// file format.
//
// pictData: A data object containing the PICT data.
//
// # Return Value
// 
// An initialized [NSPICTImageRep], or `nil` if the object could not be
// initialized. Initialization may fail if the data does not conform to the
// PICT file format.
//
// # Discussion
// 
// If the PICT data is obtained directly from a PICT file or document, this
// method ignores most of the 512-byte header that occurs before the start of
// the actual picture data. It may retrieve some relevant meta information
// from the header.
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/init(data:)
func (p NSPICTImageRep) InitWithData(pictData foundation.INSData) NSPICTImageRep {
	rv := objc.Send[NSPICTImageRep](p.ID, objc.Sel("initWithData:"), pictData)
	return rv
}

// The rectangle that bounds the image representation.
//
// # Discussion
// 
// The rectangle bounding the receiver. This rectangle is obtained from the
// the `picFrame` field in the picture header. See the Carbon QuickDraw
// Manager documentation for information on the picture header
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/boundingBox
func (p NSPICTImageRep) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("boundingBox"))
	return corefoundation.CGRect(rv)
}

// The image representation’s PICT data.
//
// # Discussion
// 
// The data does not include the 512-byte header, if it was present in the
// original data. If you want to write the data to a file, you must precede it
// with a 512-byte header (containing all zeros) if you want to conform to the
// PICT document format.
//
// See: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/pictRepresentation
func (p NSPICTImageRep) PICTRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("PICTRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}

