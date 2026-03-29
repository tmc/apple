// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSNibOutletConnector] class.
var (
	_NSNibOutletConnectorClass     NSNibOutletConnectorClass
	_NSNibOutletConnectorClassOnce sync.Once
)

func getNSNibOutletConnectorClass() NSNibOutletConnectorClass {
	_NSNibOutletConnectorClassOnce.Do(func() {
		_NSNibOutletConnectorClass = NSNibOutletConnectorClass{class: objc.GetClass("NSNibOutletConnector")}
	})
	return _NSNibOutletConnectorClass
}

// GetNSNibOutletConnectorClass returns the class object for NSNibOutletConnector.
func GetNSNibOutletConnectorClass() NSNibOutletConnectorClass {
	return getNSNibOutletConnectorClass()
}

type NSNibOutletConnectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSNibOutletConnectorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNibOutletConnectorClass) Alloc() NSNibOutletConnector {
	rv := objc.Send[NSNibOutletConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An outlet connection between Interface Builder objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector
type NSNibOutletConnector struct {
	NSNibConnector
}

// NSNibOutletConnectorFromID constructs a [NSNibOutletConnector] from an objc.ID.
//
// An outlet connection between Interface Builder objects.
func NSNibOutletConnectorFromID(id objc.ID) NSNibOutletConnector {
	return NSNibOutletConnector{NSNibConnector: NSNibConnectorFromID(id)}
}
// Ensure NSNibOutletConnector implements INSNibOutletConnector.
var _ INSNibOutletConnector = NSNibOutletConnector{}

// An interface definition for the [NSNibOutletConnector] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector
type INSNibOutletConnector interface {
	INSNibConnector
}

// Init initializes the instance.
func (n NSNibOutletConnector) Init() NSNibOutletConnector {
	rv := objc.Send[NSNibOutletConnector](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNibOutletConnector) Autorelease() NSNibOutletConnector {
	rv := objc.Send[NSNibOutletConnector](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNibOutletConnector creates a new NSNibOutletConnector instance.
func NewNSNibOutletConnector() NSNibOutletConnector {
	class := getNSNibOutletConnectorClass()
	rv := objc.Send[NSNibOutletConnector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

