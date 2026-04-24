// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLVirtualDisplay] class.
var (
	_SLVirtualDisplayClass     SLVirtualDisplayClass
	_SLVirtualDisplayClassOnce sync.Once
)

func getSLVirtualDisplayClass() SLVirtualDisplayClass {
	_SLVirtualDisplayClassOnce.Do(func() {
		_SLVirtualDisplayClass = SLVirtualDisplayClass{class: objc.GetClass("SLVirtualDisplay")}
	})
	return _SLVirtualDisplayClass
}

// GetSLVirtualDisplayClass returns the class object for SLVirtualDisplay.
func GetSLVirtualDisplayClass() SLVirtualDisplayClass {
	return getSLVirtualDisplayClass()
}

type SLVirtualDisplayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLVirtualDisplayClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLVirtualDisplayClass) Alloc() SLVirtualDisplay {
	rv := objc.Send[SLVirtualDisplay](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLVirtualDisplay.ApplySettingsError]
//   - [SLVirtualDisplay.Delegate]
//   - [SLVirtualDisplay.SetDelegate]
//   - [SLVirtualDisplay.Destroy]
//   - [SLVirtualDisplay.DisplayID]
//   - [SLVirtualDisplay.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay
type SLVirtualDisplay struct {
	objectivec.Object
}

// SLVirtualDisplayFromID constructs a [SLVirtualDisplay] from an objc.ID.
func SLVirtualDisplayFromID(id objc.ID) SLVirtualDisplay {
	return SLVirtualDisplay{objectivec.Object{ID: id}}
}

// Ensure SLVirtualDisplay implements ISLVirtualDisplay.
var _ ISLVirtualDisplay = SLVirtualDisplay{}

// An interface definition for the [SLVirtualDisplay] class.
//
// # Methods
//
//   - [ISLVirtualDisplay.ApplySettingsError]
//   - [ISLVirtualDisplay.Delegate]
//   - [ISLVirtualDisplay.SetDelegate]
//   - [ISLVirtualDisplay.Destroy]
//   - [ISLVirtualDisplay.DisplayID]
//   - [ISLVirtualDisplay.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay
type ISLVirtualDisplay interface {
	objectivec.IObject

	// Topic: Methods

	ApplySettingsError(settings objectivec.IObject) (bool, error)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	Destroy()
	DisplayID() uint32
	InitWithConfigurationError(configuration objectivec.IObject) (SLVirtualDisplay, error)
}

// Init initializes the instance.
func (s SLVirtualDisplay) Init() SLVirtualDisplay {
	rv := objc.Send[SLVirtualDisplay](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplay) Autorelease() SLVirtualDisplay {
	rv := objc.Send[SLVirtualDisplay](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplay creates a new SLVirtualDisplay instance.
func NewSLVirtualDisplay() SLVirtualDisplay {
	class := getSLVirtualDisplayClass()
	rv := objc.Send[SLVirtualDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/initWithConfiguration:error:
func NewSLVirtualDisplayWithConfigurationError(configuration objectivec.IObject) (SLVirtualDisplay, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplay{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayFromID(rv), nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/applySettings:error:
func (s SLVirtualDisplay) ApplySettingsError(settings objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("applySettings:error:"), settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("applySettings:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/destroy
func (s SLVirtualDisplay) Destroy() {
	objc.Send[objc.ID](s.ID, objc.Sel("destroy"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/initWithConfiguration:error:
func (s SLVirtualDisplay) InitWithConfigurationError(configuration objectivec.IObject) (SLVirtualDisplay, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplay{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayFromID(rv), nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/capabilities
func (_SLVirtualDisplayClass SLVirtualDisplayClass) Capabilities() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayClass.class), objc.Sel("capabilities"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/delegate
func (s SLVirtualDisplay) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (s SLVirtualDisplay) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplay/displayID
func (s SLVirtualDisplay) DisplayID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("displayID"))
	return rv
}
