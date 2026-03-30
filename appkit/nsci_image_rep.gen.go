// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCIImageRep] class.
var (
	_NSCIImageRepClass     NSCIImageRepClass
	_NSCIImageRepClassOnce sync.Once
)

func getNSCIImageRepClass() NSCIImageRepClass {
	_NSCIImageRepClassOnce.Do(func() {
		_NSCIImageRepClass = NSCIImageRepClass{class: objc.GetClass("NSCIImageRep")}
	})
	return _NSCIImageRepClass
}

// GetNSCIImageRepClass returns the class object for NSCIImageRep.
func GetNSCIImageRepClass() NSCIImageRepClass {
	return getNSCIImageRepClass()
}

type NSCIImageRepClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCIImageRepClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCIImageRepClass) Alloc() NSCIImageRep {
	rv := objc.Send[NSCIImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that can render an image from a Core Image object.
//
// # Creating Representations of Core Image Objects
//
//   - [NSCIImageRep.InitWithCIImage]: Returns a representation of an image initialized to the specified Core Image instance.
//
// # Getting Data
//
//   - [NSCIImageRep.CIImage]: The Core Image instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep
type NSCIImageRep struct {
	NSImageRep
}

// NSCIImageRepFromID constructs a [NSCIImageRep] from an objc.ID.
//
// An object that can render an image from a Core Image object.
func NSCIImageRepFromID(id objc.ID) NSCIImageRep {
	return NSCIImageRep{NSImageRep: NSImageRepFromID(id)}
}

// NOTE: NSCIImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCIImageRep] class.
//
// # Creating Representations of Core Image Objects
//
//   - [INSCIImageRep.InitWithCIImage]: Returns a representation of an image initialized to the specified Core Image instance.
//
// # Getting Data
//
//   - [INSCIImageRep.CIImage]: The Core Image instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep
type INSCIImageRep interface {
	INSImageRep

	// Topic: Creating Representations of Core Image Objects

	// Returns a representation of an image initialized to the specified Core Image instance.
	InitWithCIImage(image coreimage.CIImage) NSCIImageRep

	// Topic: Getting Data

	// The Core Image instance.
	CIImage() coreimage.CIImage
}

// Init initializes the instance.
func (c NSCIImageRep) Init() NSCIImageRep {
	rv := objc.Send[NSCIImageRep](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCIImageRep) Autorelease() NSCIImageRep {
	rv := objc.Send[NSCIImageRep](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCIImageRep creates a new NSCIImageRep instance.
func NewNSCIImageRep() NSCIImageRep {
	class := getNSCIImageRepClass()
	rv := objc.Send[NSCIImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a representation of an image initialized to the specified Core
// Image instance.
//
// image: The [CIImage] instance.
//
// # Return Value
//
// An initialized [NSCIImageRep] object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep/init(ciImage:)
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
func NewCIImageRepWithCIImage(image coreimage.CIImage) NSCIImageRep {
	instance := getNSCIImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCIImage:"), image)
	return NSCIImageRepFromID(rv)
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewCIImageRepWithCoder(coder foundation.INSCoder) NSCIImageRep {
	instance := getNSCIImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCIImageRepFromID(rv)
}

// Returns a representation of an image initialized to the specified Core
// Image instance.
//
// image: The [CIImage] instance.
//
// # Return Value
//
// An initialized [NSCIImageRep] object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep/init(ciImage:)
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
func (c NSCIImageRep) InitWithCIImage(image coreimage.CIImage) NSCIImageRep {
	rv := objc.Send[NSCIImageRep](c.ID, objc.Sel("initWithCIImage:"), image)
	return rv
}

// Creates and returns a representation of an image initialized to the
// specified Core Image instance.
//
// image: The [CIImage] instance.
//
// # Return Value
//
// An initialized [NSCIImageRep] object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep/imageRepWithCIImage:
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
func (_NSCIImageRepClass NSCIImageRepClass) ImageRepWithCIImage(image coreimage.CIImage) NSCIImageRep {
	rv := objc.Send[objc.ID](objc.ID(_NSCIImageRepClass.class), objc.Sel("imageRepWithCIImage:"), image)
	return NSCIImageRepFromID(rv)
}

// The Core Image instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSCIImageRep/ciImage
func (c NSCIImageRep) CIImage() coreimage.CIImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CIImage"))
	return coreimage.CIImageFromID(objc.ID(rv))
}
