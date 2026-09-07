// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTimer] class.
var (
	_NSTimerClass     NSTimerClass
	_NSTimerClassOnce sync.Once
)

func getNSTimerClass() NSTimerClass {
	_NSTimerClassOnce.Do(func() {
		_NSTimerClass = NSTimerClass{class: objc.GetClass("NSTimer")}
	})
	return _NSTimerClass
}

// GetNSTimerClass returns the class object for NSTimer.
func GetNSTimerClass() NSTimerClass {
	return getNSTimerClass()
}

type NSTimerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTimerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTimerClass) Alloc() NSTimer {
	rv := objc.SendIfResponds[NSTimer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other speechobjects classes. [Full Topic]
type NSTimer struct {
	objectivec.Object
}

// NSTimerFromID constructs a [NSTimer] from an objc.ID.
//
// A parent class referenced by other speechobjects classes.
func NSTimerFromID(id objc.ID) NSTimer {
	return NSTimer{objectivec.Object{ID: id}}
}

// Ensure NSTimer implements INSTimer.
var _ INSTimer = NSTimer{}

// An interface definition for the [NSTimer] class.
type INSTimer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n NSTimer) Init() NSTimer {
	rv := objc.SendIfResponds[NSTimer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSTimer) Autorelease() NSTimer {
	rv := objc.SendIfResponds[NSTimer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTimer creates a new NSTimer instance.
func NewNSTimer() NSTimer {
	class := getNSTimerClass()
	rv := objc.SendIfResponds[NSTimer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
