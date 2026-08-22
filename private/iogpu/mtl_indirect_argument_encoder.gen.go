// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIndirectArgumentEncoder] class.
var (
	_MTLIndirectArgumentEncoderClass     MTLIndirectArgumentEncoderClass
	_MTLIndirectArgumentEncoderClassOnce sync.Once
)

func getMTLIndirectArgumentEncoderClass() MTLIndirectArgumentEncoderClass {
	_MTLIndirectArgumentEncoderClassOnce.Do(func() {
		_MTLIndirectArgumentEncoderClass = MTLIndirectArgumentEncoderClass{class: objc.GetClass("_MTLIndirectArgumentEncoder")}
	})
	return _MTLIndirectArgumentEncoderClass
}

// GetMTLIndirectArgumentEncoderClass returns the class object for _MTLIndirectArgumentEncoder.
func GetMTLIndirectArgumentEncoderClass() MTLIndirectArgumentEncoderClass {
	return getMTLIndirectArgumentEncoderClass()
}

type MTLIndirectArgumentEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLIndirectArgumentEncoderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIndirectArgumentEncoderClass) Alloc() MTLIndirectArgumentEncoder {
	rv := objc.SendIfResponds[MTLIndirectArgumentEncoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other iogpu classes. [Full Topic]
type MTLIndirectArgumentEncoder struct {
	objectivec.Object
}

// MTLIndirectArgumentEncoderFromID constructs a [MTLIndirectArgumentEncoder] from an objc.ID.
//
// A parent class referenced by other iogpu classes.
func MTLIndirectArgumentEncoderFromID(id objc.ID) MTLIndirectArgumentEncoder {
	return MTLIndirectArgumentEncoder{objectivec.Object{ID: id}}
}

// Ensure MTLIndirectArgumentEncoder implements IMTLIndirectArgumentEncoder.
var _ IMTLIndirectArgumentEncoder = MTLIndirectArgumentEncoder{}

// An interface definition for the [MTLIndirectArgumentEncoder] class.
type IMTLIndirectArgumentEncoder interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MTLIndirectArgumentEncoder) Init() MTLIndirectArgumentEncoder {
	rv := objc.SendIfResponds[MTLIndirectArgumentEncoder](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLIndirectArgumentEncoder) Autorelease() MTLIndirectArgumentEncoder {
	rv := objc.SendIfResponds[MTLIndirectArgumentEncoder](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIndirectArgumentEncoder creates a new MTLIndirectArgumentEncoder instance.
func NewMTLIndirectArgumentEncoder() MTLIndirectArgumentEncoder {
	class := getMTLIndirectArgumentEncoderClass()
	rv := objc.SendIfResponds[MTLIndirectArgumentEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}
