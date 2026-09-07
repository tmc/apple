// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSThread] class.
var (
	_NSThreadClass     NSThreadClass
	_NSThreadClassOnce sync.Once
)

func getNSThreadClass() NSThreadClass {
	_NSThreadClassOnce.Do(func() {
		_NSThreadClass = NSThreadClass{class: objc.GetClass("NSThread")}
	})
	return _NSThreadClass
}

// GetNSThreadClass returns the class object for NSThread.
func GetNSThreadClass() NSThreadClass {
	return getNSThreadClass()
}

type NSThreadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSThreadClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSThreadClass) Alloc() NSThread {
	rv := objc.SendIfResponds[NSThread](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other texttospeech classes. [Full Topic]
type NSThread struct {
	objectivec.Object
}

// NSThreadFromID constructs a [NSThread] from an objc.ID.
//
// A parent class referenced by other texttospeech classes.
func NSThreadFromID(id objc.ID) NSThread {
	return NSThread{objectivec.Object{ID: id}}
}

// Ensure NSThread implements INSThread.
var _ INSThread = NSThread{}

// An interface definition for the [NSThread] class.
type INSThread interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n NSThread) Init() NSThread {
	rv := objc.SendIfResponds[NSThread](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSThread) Autorelease() NSThread {
	rv := objc.SendIfResponds[NSThread](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSThread creates a new NSThread instance.
func NewNSThread() NSThread {
	class := getNSThreadClass()
	rv := objc.SendIfResponds[NSThread](objc.ID(class.class), objc.Sel("new"))
	return rv
}
