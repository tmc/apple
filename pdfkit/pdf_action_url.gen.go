// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [PDFActionURL] class.
var (
	_PDFActionURLClass     PDFActionURLClass
	_PDFActionURLClassOnce sync.Once
)

func getPDFActionURLClass() PDFActionURLClass {
	_PDFActionURLClassOnce.Do(func() {
		_PDFActionURLClass = PDFActionURLClass{class: objc.GetClass("PDFActionURL")}
	})
	return _PDFActionURLClass
}

// GetPDFActionURLClass returns the class object for PDFActionURL.
func GetPDFActionURLClass() PDFActionURLClass {
	return getPDFActionURLClass()
}

type PDFActionURLClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionURLClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionURLClass) Alloc() PDFActionURL {
	rv := objc.Send[PDFActionURL](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFActionURL], a subclass of [PDFAction], defines methods for getting and
// setting the URL associated with a URL action.
//
// # Initializing a URL Action
//
//   - [PDFActionURL.InitWithURL]: Initializes a URL action with the specified URL.
//
// # Accessing and Changing the URL
//
//   - [PDFActionURL.URL]: Returns the URL associated with the URL action.
//   - [PDFActionURL.SetURL]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionURL
type PDFActionURL struct {
	PDFAction
}

// PDFActionURLFromID constructs a [PDFActionURL] from an objc.ID.
//
// [PDFActionURL], a subclass of [PDFAction], defines methods for getting and
// setting the URL associated with a URL action.
func PDFActionURLFromID(id objc.ID) PDFActionURL {
	return PDFActionURL{PDFAction: PDFActionFromID(id)}
}

// NOTE: PDFActionURL adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFActionURL] class.
//
// # Initializing a URL Action
//
//   - [IPDFActionURL.InitWithURL]: Initializes a URL action with the specified URL.
//
// # Accessing and Changing the URL
//
//   - [IPDFActionURL.URL]: Returns the URL associated with the URL action.
//   - [IPDFActionURL.SetURL]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionURL
type IPDFActionURL interface {
	IPDFAction

	// Topic: Initializing a URL Action

	// Initializes a URL action with the specified URL.
	InitWithURL(url foundation.INSURL) PDFActionURL

	// Topic: Accessing and Changing the URL

	// Returns the URL associated with the URL action.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)
}

// Init initializes the instance.
func (p PDFActionURL) Init() PDFActionURL {
	rv := objc.Send[PDFActionURL](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFActionURL) Autorelease() PDFActionURL {
	rv := objc.Send[PDFActionURL](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionURL creates a new PDFActionURL instance.
func NewPDFActionURL() PDFActionURL {
	class := getPDFActionURLClass()
	rv := objc.Send[PDFActionURL](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a URL action with the specified URL.
//
// url: The URL to set the action to.
//
// # Return Value
//
// An initialized [PDFActionURL] instance, or [NULL] if the object could not
// be initialized.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionURL/init(url:)
func NewPDFActionURLWithURL(url foundation.INSURL) PDFActionURL {
	instance := getPDFActionURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return PDFActionURLFromID(rv)
}

// Initializes a URL action with the specified URL.
//
// url: The URL to set the action to.
//
// # Return Value
//
// An initialized [PDFActionURL] instance, or [NULL] if the object could not
// be initialized.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionURL/init(url:)
func (p PDFActionURL) InitWithURL(url foundation.INSURL) PDFActionURL {
	rv := objc.Send[PDFActionURL](p.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Returns the URL associated with the URL action.
//
// # Return Value
//
// The URL associated with the action, or [NULL] if no URL is specified.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionURL/url
func (p PDFActionURL) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p PDFActionURL) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}
