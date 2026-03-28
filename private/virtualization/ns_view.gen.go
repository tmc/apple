// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSView] class.
var (
	_NSViewClass     NSViewClass
	_NSViewClassOnce sync.Once
)

func getNSViewClass() NSViewClass {
	_NSViewClassOnce.Do(func() {
		_NSViewClass = NSViewClass{class: objc.GetClass("NSView")}
	})
	return _NSViewClass
}

// GetNSViewClass returns the class object for NSView.
func GetNSViewClass() NSViewClass {
	return getNSViewClass()
}

type NSViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSViewClass) Alloc() NSView {
	rv := objc.Send[NSView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other Virtualization classes. [Full Topic]
type NSView struct {
	objectivec.Object
}

// NSViewFromID constructs a [NSView] from an objc.ID.
//
// A parent class referenced by other Virtualization classes.
func NSViewFromID(id objc.ID) NSView {
	return NSView{objectivec.Object{ID: id}}
}
// Ensure NSView implements INSView.
var _ INSView = NSView{}

// An interface definition for the [NSView] class.
type INSView interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v NSView) Init() NSView {
	rv := objc.Send[NSView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSView) Autorelease() NSView {
	rv := objc.Send[NSView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSView creates a new NSView instance.
func NewNSView() NSView {
	class := getNSViewClass()
	rv := objc.Send[NSView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

