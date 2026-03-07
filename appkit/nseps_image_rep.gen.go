// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSEPSImageRep] class.
var (
	_NSEPSImageRepClass     NSEPSImageRepClass
	_NSEPSImageRepClassOnce sync.Once
)

func getNSEPSImageRepClass() NSEPSImageRepClass {
	_NSEPSImageRepClassOnce.Do(func() {
		_NSEPSImageRepClass = NSEPSImageRepClass{class: objc.GetClass("NSEPSImageRep")}
	})
	return _NSEPSImageRepClass
}

// GetNSEPSImageRepClass returns the class object for NSEPSImageRep.
func GetNSEPSImageRepClass() NSEPSImageRepClass {
	return getNSEPSImageRepClass()
}

type NSEPSImageRepClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSEPSImageRepClass) Alloc() NSEPSImageRep {
	rv := objc.Send[NSEPSImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that can render an image from encapsulated PostScript (EPS) code.
//
// # Getting Data
//
//   - [NSEPSImageRep.BoundingBox]: The rectangle that bounds the image representation.
//   - [NSEPSImageRep.EPSRepresentation]: The EPS representation of the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSEPSImageRep
type NSEPSImageRep struct {
	NSImageRep
}

// NSEPSImageRepFromID constructs a [NSEPSImageRep] from an objc.ID.
//
// An object that can render an image from encapsulated PostScript (EPS) code.
func NSEPSImageRepFromID(id objc.ID) NSEPSImageRep {
	return NSEPSImageRep{NSImageRep: NSImageRepFromID(id)}
}
// NOTE: NSEPSImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSEPSImageRep] class.
//
// # Getting Data
//
//   - [INSEPSImageRep.BoundingBox]: The rectangle that bounds the image representation.
//   - [INSEPSImageRep.EPSRepresentation]: The EPS representation of the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSEPSImageRep
type INSEPSImageRep interface {
	INSImageRep

	// Topic: Getting Data

	// The rectangle that bounds the image representation.
	BoundingBox() corefoundation.CGRect
	// The EPS representation of the image representation.
	EPSRepresentation() foundation.INSData

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (e NSEPSImageRep) Init() NSEPSImageRep {
	rv := objc.Send[NSEPSImageRep](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSEPSImageRep) Autorelease() NSEPSImageRep {
	rv := objc.Send[NSEPSImageRep](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSEPSImageRep creates a new NSEPSImageRep instance.
func NewNSEPSImageRep() NSEPSImageRep {
	class := getNSEPSImageRepClass()
	rv := objc.Send[NSEPSImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewEPSImageRepWithCoder(coder foundation.INSCoder) NSEPSImageRep {
	instance := getNSEPSImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSEPSImageRepFromID(rv)
}


// Returns a representation of an image initialized with the specified EPS
// data.
//
// epsData: The EPS data representing the desired image.
//
// # Return Value
// 
// The initialized [NSEPSImageRep] object or `nil` if the object could not be
// initialized
//
// # Discussion
// 
// The size of the receiver is set using the bounding box information
// specified in the EPS header comments.
//
// See: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/init(data:)
func NewEPSImageRepWithData(epsData foundation.INSData) NSEPSImageRep {
	instance := getNSEPSImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), epsData)
	return NSEPSImageRepFromID(rv)
}






func (e NSEPSImageRep) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The rectangle that bounds the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/boundingBox
func (e NSEPSImageRep) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e.ID, objc.Sel("boundingBox"))
	return corefoundation.CGRect(rv)
}



// The EPS representation of the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/epsRepresentation
func (e NSEPSImageRep) EPSRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("EPSRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}



























