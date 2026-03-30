// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [PDFActionRemoteGoTo] class.
var (
	_PDFActionRemoteGoToClass     PDFActionRemoteGoToClass
	_PDFActionRemoteGoToClassOnce sync.Once
)

func getPDFActionRemoteGoToClass() PDFActionRemoteGoToClass {
	_PDFActionRemoteGoToClassOnce.Do(func() {
		_PDFActionRemoteGoToClass = PDFActionRemoteGoToClass{class: objc.GetClass("PDFActionRemoteGoTo")}
	})
	return _PDFActionRemoteGoToClass
}

// GetPDFActionRemoteGoToClass returns the class object for PDFActionRemoteGoTo.
func GetPDFActionRemoteGoToClass() PDFActionRemoteGoToClass {
	return getPDFActionRemoteGoToClass()
}

type PDFActionRemoteGoToClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionRemoteGoToClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionRemoteGoToClass) Alloc() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFActionRemoteGoTo], a subclass of [PDFAction], defines methods for
// getting and setting the destination of a go-to action that targets another
// document.
//
// # Initializing the Remote Go-to Action
//
//   - [PDFActionRemoteGoTo.InitWithPageIndexAtPointFileURL]: Initializes the remote go-to action with the specified page index, point, and document URL.
//
// # Accessing the Page Index of the Referenced Document
//
//   - [PDFActionRemoteGoTo.PageIndex]: Returns the zero-based page index referenced by the remote go-to action.
//   - [PDFActionRemoteGoTo.SetPageIndex]
//
// # Accessing a Point on the Referenced Page
//
//   - [PDFActionRemoteGoTo.Point]: Sets the point, in page space, on the page referenced by the remote go-to action.
//   - [PDFActionRemoteGoTo.SetPoint]
//
// # Accessing the URL of the Referenced Document
//
//   - [PDFActionRemoteGoTo.URL]: Returns the URL of the document referenced by the remote go-to action.
//   - [PDFActionRemoteGoTo.SetURL]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo
type PDFActionRemoteGoTo struct {
	PDFAction
}

// PDFActionRemoteGoToFromID constructs a [PDFActionRemoteGoTo] from an objc.ID.
//
// [PDFActionRemoteGoTo], a subclass of [PDFAction], defines methods for
// getting and setting the destination of a go-to action that targets another
// document.
func PDFActionRemoteGoToFromID(id objc.ID) PDFActionRemoteGoTo {
	return PDFActionRemoteGoTo{PDFAction: PDFActionFromID(id)}
}

// NOTE: PDFActionRemoteGoTo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFActionRemoteGoTo] class.
//
// # Initializing the Remote Go-to Action
//
//   - [IPDFActionRemoteGoTo.InitWithPageIndexAtPointFileURL]: Initializes the remote go-to action with the specified page index, point, and document URL.
//
// # Accessing the Page Index of the Referenced Document
//
//   - [IPDFActionRemoteGoTo.PageIndex]: Returns the zero-based page index referenced by the remote go-to action.
//   - [IPDFActionRemoteGoTo.SetPageIndex]
//
// # Accessing a Point on the Referenced Page
//
//   - [IPDFActionRemoteGoTo.Point]: Sets the point, in page space, on the page referenced by the remote go-to action.
//   - [IPDFActionRemoteGoTo.SetPoint]
//
// # Accessing the URL of the Referenced Document
//
//   - [IPDFActionRemoteGoTo.URL]: Returns the URL of the document referenced by the remote go-to action.
//   - [IPDFActionRemoteGoTo.SetURL]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo
type IPDFActionRemoteGoTo interface {
	IPDFAction

	// Topic: Initializing the Remote Go-to Action

	// Initializes the remote go-to action with the specified page index, point, and document URL.
	InitWithPageIndexAtPointFileURL(pageIndex uint, point corefoundation.CGPoint, url foundation.INSURL) PDFActionRemoteGoTo

	// Topic: Accessing the Page Index of the Referenced Document

	// Returns the zero-based page index referenced by the remote go-to action.
	PageIndex() uint
	SetPageIndex(value uint)

	// Topic: Accessing a Point on the Referenced Page

	// Sets the point, in page space, on the page referenced by the remote go-to action.
	Point() corefoundation.CGPoint
	SetPoint(value corefoundation.CGPoint)

	// Topic: Accessing the URL of the Referenced Document

	// Returns the URL of the document referenced by the remote go-to action.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)
}

// Init initializes the instance.
func (p PDFActionRemoteGoTo) Init() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFActionRemoteGoTo) Autorelease() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionRemoteGoTo creates a new PDFActionRemoteGoTo instance.
func NewPDFActionRemoteGoTo() PDFActionRemoteGoTo {
	class := getPDFActionRemoteGoToClass()
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the remote go-to action with the specified page index, point,
// and document URL.
//
// pageIndex: The page index of the remote document.
//
// point: The point on the page in the remote document.
//
// url: The URL of the remote PDF document.
//
// # Return Value
//
// An initialized [PDFActionRemoteGoTo] instance, or [NULL] if the object
// could not be initialized.
//
// # Discussion
//
// The [PDFActionRemoteGoTo] object uses a zero-based page index, not a
// [PDFPage] object. This simplifies the handling of remote destinations for
// documents that may not be instantiated yet.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/init(pageIndex:at:fileURL:)
func NewPDFActionRemoteGoToWithPageIndexAtPointFileURL(pageIndex uint, point corefoundation.CGPoint, url foundation.INSURL) PDFActionRemoteGoTo {
	instance := getPDFActionRemoteGoToClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPageIndex:atPoint:fileURL:"), pageIndex, point, url)
	return PDFActionRemoteGoToFromID(rv)
}

// Initializes the remote go-to action with the specified page index, point,
// and document URL.
//
// pageIndex: The page index of the remote document.
//
// point: The point on the page in the remote document.
//
// url: The URL of the remote PDF document.
//
// # Return Value
//
// An initialized [PDFActionRemoteGoTo] instance, or [NULL] if the object
// could not be initialized.
//
// # Discussion
//
// The [PDFActionRemoteGoTo] object uses a zero-based page index, not a
// [PDFPage] object. This simplifies the handling of remote destinations for
// documents that may not be instantiated yet.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/init(pageIndex:at:fileURL:)
func (p PDFActionRemoteGoTo) InitWithPageIndexAtPointFileURL(pageIndex uint, point corefoundation.CGPoint, url foundation.INSURL) PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p.ID, objc.Sel("initWithPageIndex:atPoint:fileURL:"), pageIndex, point, url)
	return rv
}

// Returns the zero-based page index referenced by the remote go-to action.
//
// # Return Value
//
// The page index referenced by the remote go-to action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/pageIndex
func (p PDFActionRemoteGoTo) PageIndex() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("pageIndex"))
	return rv
}
func (p PDFActionRemoteGoTo) SetPageIndex(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setPageIndex:"), value)
}

// Sets the point, in page space, on the page referenced by the remote go-to
// action.
//
// # Discussion
//
// Page space is a 72-dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/point
func (p PDFActionRemoteGoTo) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}
func (p PDFActionRemoteGoTo) SetPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](p.ID, objc.Sel("setPoint:"), value)
}

// Returns the URL of the document referenced by the remote go-to action.
//
// # Return Value
//
// The URL of the remote document referenced by the action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/url
func (p PDFActionRemoteGoTo) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p PDFActionRemoteGoTo) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}
