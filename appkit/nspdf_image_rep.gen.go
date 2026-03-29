// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSPDFImageRep] class.
var (
	_NSPDFImageRepClass     NSPDFImageRepClass
	_NSPDFImageRepClassOnce sync.Once
)

func getNSPDFImageRepClass() NSPDFImageRepClass {
	_NSPDFImageRepClassOnce.Do(func() {
		_NSPDFImageRepClass = NSPDFImageRepClass{class: objc.GetClass("NSPDFImageRep")}
	})
	return _NSPDFImageRepClass
}

// GetNSPDFImageRepClass returns the class object for NSPDFImageRep.
func GetNSPDFImageRepClass() NSPDFImageRepClass {
	return getNSPDFImageRepClass()
}

type NSPDFImageRepClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPDFImageRepClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPDFImageRepClass) Alloc() NSPDFImageRep {
	rv := objc.Send[NSPDFImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that can render an image from a PDF format data stream.
//
// # Creating Representations of Images from PDF Data
//
//   - [NSPDFImageRep.InitWithData]: Returns a representation of an image initialized with the specified PDF data.
//
// # Getting Data
//
//   - [NSPDFImageRep.Bounds]: The image representation’s bounding rectangle.
//   - [NSPDFImageRep.CurrentPage]: The page currently displayed by the image representation.
//   - [NSPDFImageRep.SetCurrentPage]
//   - [NSPDFImageRep.PageCount]: The number of pages in the image representation.
//   - [NSPDFImageRep.PDFRepresentation]: The PDF representation of the representation’s image.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep
type NSPDFImageRep struct {
	NSImageRep
}

// NSPDFImageRepFromID constructs a [NSPDFImageRep] from an objc.ID.
//
// An object that can render an image from a PDF format data stream.
func NSPDFImageRepFromID(id objc.ID) NSPDFImageRep {
	return NSPDFImageRep{NSImageRep: NSImageRepFromID(id)}
}
// NOTE: NSPDFImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPDFImageRep] class.
//
// # Creating Representations of Images from PDF Data
//
//   - [INSPDFImageRep.InitWithData]: Returns a representation of an image initialized with the specified PDF data.
//
// # Getting Data
//
//   - [INSPDFImageRep.Bounds]: The image representation’s bounding rectangle.
//   - [INSPDFImageRep.CurrentPage]: The page currently displayed by the image representation.
//   - [INSPDFImageRep.SetCurrentPage]
//   - [INSPDFImageRep.PageCount]: The number of pages in the image representation.
//   - [INSPDFImageRep.PDFRepresentation]: The PDF representation of the representation’s image.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep
type INSPDFImageRep interface {
	INSImageRep

	// Topic: Creating Representations of Images from PDF Data

	// Returns a representation of an image initialized with the specified PDF data.
	InitWithData(pdfData foundation.INSData) NSPDFImageRep

	// Topic: Getting Data

	// The image representation’s bounding rectangle.
	Bounds() corefoundation.CGRect
	// The page currently displayed by the image representation.
	CurrentPage() int
	SetCurrentPage(value int)
	// The number of pages in the image representation.
	PageCount() int
	// The PDF representation of the representation’s image.
	PDFRepresentation() foundation.INSData
}

// Init initializes the instance.
func (p NSPDFImageRep) Init() NSPDFImageRep {
	rv := objc.Send[NSPDFImageRep](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPDFImageRep) Autorelease() NSPDFImageRep {
	rv := objc.Send[NSPDFImageRep](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPDFImageRep creates a new NSPDFImageRep instance.
func NewNSPDFImageRep() NSPDFImageRep {
	class := getNSPDFImageRepClass()
	rv := objc.Send[NSPDFImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewPDFImageRepWithCoder(coder foundation.INSCoder) NSPDFImageRep {
	instance := getNSPDFImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPDFImageRepFromID(rv)
}

// Returns a representation of an image initialized with the specified PDF
// data.
//
// pdfData: A data object containing the PDF data for the image.
//
// # Return Value
// 
// An initialized [NSPDFImageRep] object, or `nil` if the object could not be
// initialized. Initialization may fail if the PDF data does not conform to
// the PDF file format.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/init(data:)
func NewPDFImageRepWithData(pdfData foundation.INSData) NSPDFImageRep {
	instance := getNSPDFImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), pdfData)
	return NSPDFImageRepFromID(rv)
}

// Returns a representation of an image initialized with the specified PDF
// data.
//
// pdfData: A data object containing the PDF data for the image.
//
// # Return Value
// 
// An initialized [NSPDFImageRep] object, or `nil` if the object could not be
// initialized. Initialization may fail if the PDF data does not conform to
// the PDF file format.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/init(data:)
func (p NSPDFImageRep) InitWithData(pdfData foundation.INSData) NSPDFImageRep {
	rv := objc.Send[NSPDFImageRep](p.ID, objc.Sel("initWithData:"), pdfData)
	return rv
}

// The image representation’s bounding rectangle.
//
// # Discussion
// 
// This value is equivalent to the crop box specified by the PDF data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/bounds
func (p NSPDFImageRep) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
// The page currently displayed by the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/currentPage
func (p NSPDFImageRep) CurrentPage() int {
	rv := objc.Send[int](p.ID, objc.Sel("currentPage"))
	return rv
}
func (p NSPDFImageRep) SetCurrentPage(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentPage:"), value)
}
// The number of pages in the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pageCount
func (p NSPDFImageRep) PageCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("pageCount"))
	return rv
}
// The PDF representation of the representation’s image.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pdfRepresentation
func (p NSPDFImageRep) PDFRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("PDFRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}

