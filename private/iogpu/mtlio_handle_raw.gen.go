// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIOHandleRaw] class.
var (
	_MTLIOHandleRawClass     MTLIOHandleRawClass
	_MTLIOHandleRawClassOnce sync.Once
)

func getMTLIOHandleRawClass() MTLIOHandleRawClass {
	_MTLIOHandleRawClassOnce.Do(func() {
		_MTLIOHandleRawClass = MTLIOHandleRawClass{class: objc.GetClass("_MTLIOHandleRaw")}
	})
	return _MTLIOHandleRawClass
}

// GetMTLIOHandleRawClass returns the class object for _MTLIOHandleRaw.
func GetMTLIOHandleRawClass() MTLIOHandleRawClass {
	return getMTLIOHandleRawClass()
}

type MTLIOHandleRawClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLIOHandleRawClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIOHandleRawClass) Alloc() MTLIOHandleRaw {
	rv := objc.SendIfResponds[MTLIOHandleRaw](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other iogpu classes. [Full Topic]
type MTLIOHandleRaw struct {
	objectivec.Object
}

// MTLIOHandleRawFromID constructs a [MTLIOHandleRaw] from an objc.ID.
//
// A parent class referenced by other iogpu classes.
func MTLIOHandleRawFromID(id objc.ID) MTLIOHandleRaw {
	return MTLIOHandleRaw{objectivec.Object{ID: id}}
}

// Ensure MTLIOHandleRaw implements IMTLIOHandleRaw.
var _ IMTLIOHandleRaw = MTLIOHandleRaw{}

// An interface definition for the [MTLIOHandleRaw] class.
type IMTLIOHandleRaw interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MTLIOHandleRaw) Init() MTLIOHandleRaw {
	rv := objc.SendIfResponds[MTLIOHandleRaw](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLIOHandleRaw) Autorelease() MTLIOHandleRaw {
	rv := objc.SendIfResponds[MTLIOHandleRaw](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIOHandleRaw creates a new MTLIOHandleRaw instance.
func NewMTLIOHandleRaw() MTLIOHandleRaw {
	class := getMTLIOHandleRawClass()
	rv := objc.SendIfResponds[MTLIOHandleRaw](objc.ID(class.class), objc.Sel("new"))
	return rv
}
