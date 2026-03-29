// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSHapticFeedbackManager] class.
var (
	_NSHapticFeedbackManagerClass     NSHapticFeedbackManagerClass
	_NSHapticFeedbackManagerClassOnce sync.Once
)

func getNSHapticFeedbackManagerClass() NSHapticFeedbackManagerClass {
	_NSHapticFeedbackManagerClassOnce.Do(func() {
		_NSHapticFeedbackManagerClass = NSHapticFeedbackManagerClass{class: objc.GetClass("NSHapticFeedbackManager")}
	})
	return _NSHapticFeedbackManagerClass
}

// GetNSHapticFeedbackManagerClass returns the class object for NSHapticFeedbackManager.
func GetNSHapticFeedbackManagerClass() NSHapticFeedbackManagerClass {
	return getNSHapticFeedbackManagerClass()
}

type NSHapticFeedbackManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSHapticFeedbackManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSHapticFeedbackManagerClass) Alloc() NSHapticFeedbackManager {
	rv := objc.Send[NSHapticFeedbackManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides access to the haptic feedback management attributes
// on a system with a Force Touch trackpad.
//
// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager
type NSHapticFeedbackManager struct {
	objectivec.Object
}

// NSHapticFeedbackManagerFromID constructs a [NSHapticFeedbackManager] from an objc.ID.
//
// An object that provides access to the haptic feedback management attributes
// on a system with a Force Touch trackpad.
func NSHapticFeedbackManagerFromID(id objc.ID) NSHapticFeedbackManager {
	return NSHapticFeedbackManager{objectivec.Object{ID: id}}
}
// NOTE: NSHapticFeedbackManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSHapticFeedbackManager] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager
type INSHapticFeedbackManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h NSHapticFeedbackManager) Init() NSHapticFeedbackManager {
	rv := objc.Send[NSHapticFeedbackManager](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h NSHapticFeedbackManager) Autorelease() NSHapticFeedbackManager {
	rv := objc.Send[NSHapticFeedbackManager](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSHapticFeedbackManager creates a new NSHapticFeedbackManager instance.
func NewNSHapticFeedbackManager() NSHapticFeedbackManager {
	class := getNSHapticFeedbackManagerClass()
	rv := objc.Send[NSHapticFeedbackManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests a haptic feedback performer object that is based on the current
// input device, accessibility settings, and user preferences.
//
// # Discussion
// 
// This method returns a haptic feedback performer object of type
// [NSHapticFeedbackPerformer] that is based on the current input device,
// accessibility settings, and user preferences. Because the current input
// device may change at any time, you should request the default performer
// whenever you need to provide haptic feedback to the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/defaultPerformer
func (_NSHapticFeedbackManagerClass NSHapticFeedbackManagerClass) DefaultPerformer() NSHapticFeedbackPerformer {
	rv := objc.Send[objc.ID](objc.ID(_NSHapticFeedbackManagerClass.class), objc.Sel("defaultPerformer"))
	return NSHapticFeedbackPerformerObjectFromID(rv)
}

