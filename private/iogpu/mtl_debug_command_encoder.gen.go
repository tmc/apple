// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLDebugCommandEncoder] class.
var (
	_MTLDebugCommandEncoderClass     MTLDebugCommandEncoderClass
	_MTLDebugCommandEncoderClassOnce sync.Once
)

func getMTLDebugCommandEncoderClass() MTLDebugCommandEncoderClass {
	_MTLDebugCommandEncoderClassOnce.Do(func() {
		_MTLDebugCommandEncoderClass = MTLDebugCommandEncoderClass{class: objc.GetClass("_MTLDebugCommandEncoder")}
	})
	return _MTLDebugCommandEncoderClass
}

// GetMTLDebugCommandEncoderClass returns the class object for _MTLDebugCommandEncoder.
func GetMTLDebugCommandEncoderClass() MTLDebugCommandEncoderClass {
	return getMTLDebugCommandEncoderClass()
}

type MTLDebugCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLDebugCommandEncoderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLDebugCommandEncoderClass) Alloc() MTLDebugCommandEncoder {
	rv := objc.SendIfResponds[MTLDebugCommandEncoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other iogpu classes. [Full Topic]
type MTLDebugCommandEncoder struct {
	objectivec.Object
}

// MTLDebugCommandEncoderFromID constructs a [MTLDebugCommandEncoder] from an objc.ID.
//
// A parent class referenced by other iogpu classes.
func MTLDebugCommandEncoderFromID(id objc.ID) MTLDebugCommandEncoder {
	return MTLDebugCommandEncoder{objectivec.Object{ID: id}}
}

// Ensure MTLDebugCommandEncoder implements IMTLDebugCommandEncoder.
var _ IMTLDebugCommandEncoder = MTLDebugCommandEncoder{}

// An interface definition for the [MTLDebugCommandEncoder] class.
type IMTLDebugCommandEncoder interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MTLDebugCommandEncoder) Init() MTLDebugCommandEncoder {
	rv := objc.SendIfResponds[MTLDebugCommandEncoder](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLDebugCommandEncoder) Autorelease() MTLDebugCommandEncoder {
	rv := objc.SendIfResponds[MTLDebugCommandEncoder](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLDebugCommandEncoder creates a new MTLDebugCommandEncoder instance.
func NewMTLDebugCommandEncoder() MTLDebugCommandEncoder {
	class := getMTLDebugCommandEncoderClass()
	rv := objc.SendIfResponds[MTLDebugCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}
