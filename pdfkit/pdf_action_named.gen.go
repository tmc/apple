// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PDFActionNamed] class.
var (
	_PDFActionNamedClass     PDFActionNamedClass
	_PDFActionNamedClassOnce sync.Once
)

func getPDFActionNamedClass() PDFActionNamedClass {
	_PDFActionNamedClassOnce.Do(func() {
		_PDFActionNamedClass = PDFActionNamedClass{class: objc.GetClass("PDFActionNamed")}
	})
	return _PDFActionNamedClass
}

// GetPDFActionNamedClass returns the class object for PDFActionNamed.
func GetPDFActionNamedClass() PDFActionNamedClass {
	return getPDFActionNamedClass()
}

type PDFActionNamedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionNamedClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionNamedClass) Alloc() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFActionNamed] defines methods used to work with actions in PDF
// documents, some of which are named in the Adobe PDF Specification.
//
// # Overview
//
// A [PDFActionNamed] object represents an action with a defined name, such as
// “Go back” or “Zoom in.”
//
// # Accessing the Name of the Action
//
//   - [PDFActionNamed.Name]: Returns the name of the named action.
//   - [PDFActionNamed.SetName]
//
// # Initializing the Action
//
//   - [PDFActionNamed.InitWithName]: Initializes the [PDFActionName] object with the specified named action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamed
type PDFActionNamed struct {
	PDFAction
}

// PDFActionNamedFromID constructs a [PDFActionNamed] from an objc.ID.
//
// [PDFActionNamed] defines methods used to work with actions in PDF
// documents, some of which are named in the Adobe PDF Specification.
func PDFActionNamedFromID(id objc.ID) PDFActionNamed {
	return PDFActionNamed{PDFAction: PDFActionFromID(id)}
}

// NOTE: PDFActionNamed adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFActionNamed] class.
//
// # Accessing the Name of the Action
//
//   - [IPDFActionNamed.Name]: Returns the name of the named action.
//   - [IPDFActionNamed.SetName]
//
// # Initializing the Action
//
//   - [IPDFActionNamed.InitWithName]: Initializes the [PDFActionName] object with the specified named action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamed
type IPDFActionNamed interface {
	IPDFAction

	// Topic: Accessing the Name of the Action

	// Returns the name of the named action.
	Name() PDFActionNamedName
	SetName(value PDFActionNamedName)

	// Topic: Initializing the Action

	// Initializes the [PDFActionName] object with the specified named action.
	InitWithName(name PDFActionNamedName) PDFActionNamed
}

// Init initializes the instance.
func (p PDFActionNamed) Init() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFActionNamed) Autorelease() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionNamed creates a new PDFActionNamed instance.
func NewPDFActionNamed() PDFActionNamed {
	class := getPDFActionNamedClass()
	rv := objc.Send[PDFActionNamed](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the [PDFActionName] object with the specified named action.
//
// name: The action name used to initialize the named action.
//
// # Return Value
//
// An initialized [PDFActionNamed] instance, or [NULL] if the object could not
// be initialized.
//
// # Discussion
//
// See [PDFActionNamed] for the names of named actions you can specify.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/init(name:)
func NewPDFActionNamedWithName(name PDFActionNamedName) PDFActionNamed {
	instance := getPDFActionNamedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), name)
	return PDFActionNamedFromID(rv)
}

// Initializes the [PDFActionName] object with the specified named action.
//
// name: The action name used to initialize the named action.
//
// # Return Value
//
// An initialized [PDFActionNamed] instance, or [NULL] if the object could not
// be initialized.
//
// # Discussion
//
// See [PDFActionNamed] for the names of named actions you can specify.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/init(name:)
func (p PDFActionNamed) InitWithName(name PDFActionNamedName) PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p.ID, objc.Sel("initWithName:"), name)
	return rv
}

// Returns the name of the named action.
//
// # Return Value
//
// The name of the named action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/name
func (p PDFActionNamed) Name() PDFActionNamedName {
	rv := objc.Send[PDFActionNamedName](p.ID, objc.Sel("name"))
	return PDFActionNamedName(rv)
}
func (p PDFActionNamed) SetName(value PDFActionNamedName) {
	objc.Send[struct{}](p.ID, objc.Sel("setName:"), value)
}
