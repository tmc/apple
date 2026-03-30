// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PDFActionGoTo] class.
var (
	_PDFActionGoToClass     PDFActionGoToClass
	_PDFActionGoToClassOnce sync.Once
)

func getPDFActionGoToClass() PDFActionGoToClass {
	_PDFActionGoToClassOnce.Do(func() {
		_PDFActionGoToClass = PDFActionGoToClass{class: objc.GetClass("PDFActionGoTo")}
	})
	return _PDFActionGoToClass
}

// GetPDFActionGoToClass returns the class object for PDFActionGoTo.
func GetPDFActionGoToClass() PDFActionGoToClass {
	return getPDFActionGoToClass()
}

type PDFActionGoToClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionGoToClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionGoToClass) Alloc() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFActionGoTo], a subclass of [PDFAction], defines methods for getting and
// setting the destination of a go-to action.
//
// # Overview
//
// A [PDFActionGoTo] object represents the action of going to a specific
// location within the PDF document.
//
// # Accessing the Destination
//
//   - [PDFActionGoTo.Destination]: Returns the destination associated with the action.
//   - [PDFActionGoTo.SetDestination]
//
// # Initializing the Action
//
//   - [PDFActionGoTo.InitWithDestination]: Initializes the go-to action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo
type PDFActionGoTo struct {
	PDFAction
}

// PDFActionGoToFromID constructs a [PDFActionGoTo] from an objc.ID.
//
// [PDFActionGoTo], a subclass of [PDFAction], defines methods for getting and
// setting the destination of a go-to action.
func PDFActionGoToFromID(id objc.ID) PDFActionGoTo {
	return PDFActionGoTo{PDFAction: PDFActionFromID(id)}
}

// NOTE: PDFActionGoTo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFActionGoTo] class.
//
// # Accessing the Destination
//
//   - [IPDFActionGoTo.Destination]: Returns the destination associated with the action.
//   - [IPDFActionGoTo.SetDestination]
//
// # Initializing the Action
//
//   - [IPDFActionGoTo.InitWithDestination]: Initializes the go-to action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo
type IPDFActionGoTo interface {
	IPDFAction

	// Topic: Accessing the Destination

	// Returns the destination associated with the action.
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)

	// Topic: Initializing the Action

	// Initializes the go-to action.
	InitWithDestination(destination IPDFDestination) PDFActionGoTo
}

// Init initializes the instance.
func (p PDFActionGoTo) Init() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFActionGoTo) Autorelease() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionGoTo creates a new PDFActionGoTo instance.
func NewPDFActionGoTo() PDFActionGoTo {
	class := getPDFActionGoToClass()
	rv := objc.Send[PDFActionGoTo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the go-to action.
//
// destination: The destination with which to initialize the go-to action.
//
// # Return Value
//
// An initialized [PDFActionGoTo] instance, or [NULL] if the object could not
// be initialized.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/init(destination:)
func NewPDFActionGoToWithDestination(destination IPDFDestination) PDFActionGoTo {
	instance := getPDFActionGoToClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestination:"), destination)
	return PDFActionGoToFromID(rv)
}

// Initializes the go-to action.
//
// destination: The destination with which to initialize the go-to action.
//
// # Return Value
//
// An initialized [PDFActionGoTo] instance, or [NULL] if the object could not
// be initialized.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/init(destination:)
func (p PDFActionGoTo) InitWithDestination(destination IPDFDestination) PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p.ID, objc.Sel("initWithDestination:"), destination)
	return rv
}

// Returns the destination associated with the action.
//
// # Return Value
//
// The destination specified by the go-to action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/destination
func (p PDFActionGoTo) Destination() IPDFDestination {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destination"))
	return PDFDestinationFromID(objc.ID(rv))
}
func (p PDFActionGoTo) SetDestination(value IPDFDestination) {
	objc.Send[struct{}](p.ID, objc.Sel("setDestination:"), value)
}
