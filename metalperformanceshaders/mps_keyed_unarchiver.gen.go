// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSKeyedUnarchiver] class.
var (
	_MPSKeyedUnarchiverClass     MPSKeyedUnarchiverClass
	_MPSKeyedUnarchiverClassOnce sync.Once
)

func getMPSKeyedUnarchiverClass() MPSKeyedUnarchiverClass {
	_MPSKeyedUnarchiverClassOnce.Do(func() {
		_MPSKeyedUnarchiverClass = MPSKeyedUnarchiverClass{class: objc.GetClass("MPSKeyedUnarchiver")}
	})
	return _MPSKeyedUnarchiverClass
}

// GetMPSKeyedUnarchiverClass returns the class object for MPSKeyedUnarchiver.
func GetMPSKeyedUnarchiverClass() MPSKeyedUnarchiverClass {
	return getMPSKeyedUnarchiverClass()
}

type MPSKeyedUnarchiverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSKeyedUnarchiverClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSKeyedUnarchiverClass) Alloc() MPSKeyedUnarchiver {
	rv := objc.Send[MPSKeyedUnarchiver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A keyed archiver that supports Metal Performance Shaders kernel decoding.
//
// # Initializers
//
//   - [MPSKeyedUnarchiver.InitForReadingFromDataDeviceError]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver
type MPSKeyedUnarchiver struct {
	foundation.NSKeyedUnarchiver
}

// MPSKeyedUnarchiverFromID constructs a [MPSKeyedUnarchiver] from an objc.ID.
//
// A keyed archiver that supports Metal Performance Shaders kernel decoding.
func MPSKeyedUnarchiverFromID(id objc.ID) MPSKeyedUnarchiver {
	return MPSKeyedUnarchiver{NSKeyedUnarchiver: foundation.NSKeyedUnarchiverFromID(id)}
}

// NOTE: MPSKeyedUnarchiver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSKeyedUnarchiver] class.
//
// # Initializers
//
//   - [IMPSKeyedUnarchiver.InitForReadingFromDataDeviceError]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver
type IMPSKeyedUnarchiver interface {
	foundation.INSKeyedUnarchiver
	MPSDeviceProvider

	// Topic: Initializers

	InitForReadingFromDataDeviceError(data foundation.NSData, device metal.MTLDevice) (MPSKeyedUnarchiver, error)
}

// Init initializes the instance.
func (k MPSKeyedUnarchiver) Init() MPSKeyedUnarchiver {
	rv := objc.Send[MPSKeyedUnarchiver](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MPSKeyedUnarchiver) Autorelease() MPSKeyedUnarchiver {
	rv := objc.Send[MPSKeyedUnarchiver](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSKeyedUnarchiver creates a new MPSKeyedUnarchiver instance.
func NewMPSKeyedUnarchiver() MPSKeyedUnarchiver {
	class := getMPSKeyedUnarchiverClass()
	rv := objc.Send[MPSKeyedUnarchiver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/init(forReadingFrom:device:error:)
func NewKeyedUnarchiverForReadingFromDataDeviceError(data foundation.NSData, device metal.MTLDevice) (MPSKeyedUnarchiver, error) {
	var errorPtr objc.ID
	instance := getMPSKeyedUnarchiverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForReadingFromData:device:error:"), data, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSKeyedUnarchiver{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MPSKeyedUnarchiver{}, objc.ErrInitFailed
	}
	return MPSKeyedUnarchiverFromID(rv), nil
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/init(forReadingFrom:device:error:)
func (k MPSKeyedUnarchiver) InitForReadingFromDataDeviceError(data foundation.NSData, device metal.MTLDevice) (MPSKeyedUnarchiver, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("initForReadingFromData:device:error:"), data, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSKeyedUnarchiver{}, foundation.NSErrorFrom(errorPtr)
	}
	return MPSKeyedUnarchiverFromID(rv), nil

}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/mpsMTLDevice()
func (k MPSKeyedUnarchiver) MpsMTLDevice() metal.MTLDevice {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("mpsMTLDevice"))
	return metal.MTLDeviceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/unarchivedObject(of:from:device:)
func (_MPSKeyedUnarchiverClass MPSKeyedUnarchiverClass) UnarchivedObjectOfClassFromDataDeviceError(cls objectivec.Class, data foundation.NSData, device metal.MTLDevice) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MPSKeyedUnarchiverClass.class), objc.Sel("unarchivedObjectOfClass:fromData:device:error:"), cls, data, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/unarchivedObject(ofClasses:from:device:)
func (_MPSKeyedUnarchiverClass MPSKeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceError(classes foundation.INSSet, data foundation.NSData, device metal.MTLDevice) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MPSKeyedUnarchiverClass.class), objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), classes, data, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Protocol methods for MPSDeviceProvider
