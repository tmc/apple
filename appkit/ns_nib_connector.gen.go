// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSNibConnector] class.
var (
	_NSNibConnectorClass     NSNibConnectorClass
	_NSNibConnectorClassOnce sync.Once
)

func getNSNibConnectorClass() NSNibConnectorClass {
	_NSNibConnectorClassOnce.Do(func() {
		_NSNibConnectorClass = NSNibConnectorClass{class: objc.GetClass("NSNibConnector")}
	})
	return _NSNibConnectorClass
}

// GetNSNibConnectorClass returns the class object for NSNibConnector.
func GetNSNibConnectorClass() NSNibConnectorClass {
	return getNSNibConnectorClass()
}

type NSNibConnectorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNibConnectorClass) Alloc() NSNibConnector {
	rv := objc.Send[NSNibConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A connection between two nibs.
//
// # Instance Properties
//
//   - [NSNibConnector.Destination]
//   - [NSNibConnector.SetDestination]
//   - [NSNibConnector.Label]
//   - [NSNibConnector.SetLabel]
//   - [NSNibConnector.Source]
//   - [NSNibConnector.SetSource]
//
// # Instance Methods
//
//   - [NSNibConnector.EstablishConnection]
//   - [NSNibConnector.ReplaceObjectWithObject]
//
// See: https://developer.apple.com/documentation/AppKit/NSNibConnector
type NSNibConnector struct {
	objectivec.Object
}

// NSNibConnectorFromID constructs a [NSNibConnector] from an objc.ID.
//
// A connection between two nibs.
func NSNibConnectorFromID(id objc.ID) NSNibConnector {
	return NSNibConnector{objectivec.Object{ID: id}}
}
// NOTE: NSNibConnector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSNibConnector] class.
//
// # Instance Properties
//
//   - [INSNibConnector.Destination]
//   - [INSNibConnector.SetDestination]
//   - [INSNibConnector.Label]
//   - [INSNibConnector.SetLabel]
//   - [INSNibConnector.Source]
//   - [INSNibConnector.SetSource]
//
// # Instance Methods
//
//   - [INSNibConnector.EstablishConnection]
//   - [INSNibConnector.ReplaceObjectWithObject]
//
// See: https://developer.apple.com/documentation/AppKit/NSNibConnector
type INSNibConnector interface {
	objectivec.IObject

	// Topic: Instance Properties

	Destination() objectivec.IObject
	SetDestination(value objectivec.IObject)
	Label() string
	SetLabel(value string)
	Source() objectivec.IObject
	SetSource(value objectivec.IObject)

	// Topic: Instance Methods

	EstablishConnection()
	ReplaceObjectWithObject(oldObject objectivec.IObject, newObject objectivec.IObject)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n NSNibConnector) Init() NSNibConnector {
	rv := objc.Send[NSNibConnector](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNibConnector) Autorelease() NSNibConnector {
	rv := objc.Send[NSNibConnector](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNibConnector creates a new NSNibConnector instance.
func NewNSNibConnector() NSNibConnector {
	class := getNSNibConnectorClass()
	rv := objc.Send[NSNibConnector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSNibConnector/establishConnection
func (n NSNibConnector) EstablishConnection() {
	objc.Send[objc.ID](n.ID, objc.Sel("establishConnection"))
}
//
// See: https://developer.apple.com/documentation/AppKit/NSNibConnector/replaceObject:withObject:
func (n NSNibConnector) ReplaceObjectWithObject(oldObject objectivec.IObject, newObject objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("replaceObject:withObject:"), oldObject, newObject)
}
func (n NSNibConnector) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AppKit/NSNibConnector/destination
func (n NSNibConnector) Destination() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("destination"))
	return objectivec.Object{ID: rv}
}
func (n NSNibConnector) SetDestination(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setDestination:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSNibConnector/label
func (n NSNibConnector) Label() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (n NSNibConnector) SetLabel(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setLabel:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSNibConnector/source
func (n NSNibConnector) Source() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("source"))
	return objectivec.Object{ID: rv}
}
func (n NSNibConnector) SetSource(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setSource:"), value)
}

