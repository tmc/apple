// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSNibControlConnector] class.
var (
	_NSNibControlConnectorClass     NSNibControlConnectorClass
	_NSNibControlConnectorClassOnce sync.Once
)

func getNSNibControlConnectorClass() NSNibControlConnectorClass {
	_NSNibControlConnectorClassOnce.Do(func() {
		_NSNibControlConnectorClass = NSNibControlConnectorClass{class: objc.GetClass("NSNibControlConnector")}
	})
	return _NSNibControlConnectorClass
}

// GetNSNibControlConnectorClass returns the class object for NSNibControlConnector.
func GetNSNibControlConnectorClass() NSNibControlConnectorClass {
	return getNSNibControlConnectorClass()
}

type NSNibControlConnectorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNibControlConnectorClass) Alloc() NSNibControlConnector {
	rv := objc.Send[NSNibControlConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A control connection between two Interface Builder objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSNibControlConnector
type NSNibControlConnector struct {
	NSNibConnector
}

// NSNibControlConnectorFromID constructs a [NSNibControlConnector] from an objc.ID.
//
// A control connection between two Interface Builder objects.
func NSNibControlConnectorFromID(id objc.ID) NSNibControlConnector {
	return NSNibControlConnector{NSNibConnector: NSNibConnectorFromID(id)}
}
// Ensure NSNibControlConnector implements INSNibControlConnector.
var _ INSNibControlConnector = NSNibControlConnector{}





// An interface definition for the [NSNibControlConnector] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSNibControlConnector
type INSNibControlConnector interface {
	INSNibConnector
}





// Init initializes the instance.
func (n NSNibControlConnector) Init() NSNibControlConnector {
	rv := objc.Send[NSNibControlConnector](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNibControlConnector) Autorelease() NSNibControlConnector {
	rv := objc.Send[NSNibControlConnector](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNibControlConnector creates a new NSNibControlConnector instance.
func NewNSNibControlConnector() NSNibControlConnector {
	class := getNSNibControlConnectorClass()
	rv := objc.Send[NSNibControlConnector](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































