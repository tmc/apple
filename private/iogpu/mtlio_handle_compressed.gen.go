// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIOHandleCompressed] class.
var (
	_MTLIOHandleCompressedClass     MTLIOHandleCompressedClass
	_MTLIOHandleCompressedClassOnce sync.Once
)

func getMTLIOHandleCompressedClass() MTLIOHandleCompressedClass {
	_MTLIOHandleCompressedClassOnce.Do(func() {
		_MTLIOHandleCompressedClass = MTLIOHandleCompressedClass{class: objc.GetClass("_MTLIOHandleCompressed")}
	})
	return _MTLIOHandleCompressedClass
}

// GetMTLIOHandleCompressedClass returns the class object for _MTLIOHandleCompressed.
func GetMTLIOHandleCompressedClass() MTLIOHandleCompressedClass {
	return getMTLIOHandleCompressedClass()
}

type MTLIOHandleCompressedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLIOHandleCompressedClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIOHandleCompressedClass) Alloc() MTLIOHandleCompressed {
	rv := objc.SendIfResponds[MTLIOHandleCompressed](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other iogpu classes. [Full Topic]
type MTLIOHandleCompressed struct {
	objectivec.Object
}

// MTLIOHandleCompressedFromID constructs a [MTLIOHandleCompressed] from an objc.ID.
//
// A parent class referenced by other iogpu classes.
func MTLIOHandleCompressedFromID(id objc.ID) MTLIOHandleCompressed {
	return MTLIOHandleCompressed{objectivec.Object{ID: id}}
}

// Ensure MTLIOHandleCompressed implements IMTLIOHandleCompressed.
var _ IMTLIOHandleCompressed = MTLIOHandleCompressed{}

// An interface definition for the [MTLIOHandleCompressed] class.
type IMTLIOHandleCompressed interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MTLIOHandleCompressed) Init() MTLIOHandleCompressed {
	rv := objc.SendIfResponds[MTLIOHandleCompressed](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLIOHandleCompressed) Autorelease() MTLIOHandleCompressed {
	rv := objc.SendIfResponds[MTLIOHandleCompressed](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIOHandleCompressed creates a new MTLIOHandleCompressed instance.
func NewMTLIOHandleCompressed() MTLIOHandleCompressed {
	class := getMTLIOHandleCompressedClass()
	rv := objc.SendIfResponds[MTLIOHandleCompressed](objc.ID(class.class), objc.Sel("new"))
	return rv
}
