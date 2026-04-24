// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyEventSequenceTrackerProviderDefault] class.
var (
	_CPXKeyEventSequenceTrackerProviderDefaultClass     CPXKeyEventSequenceTrackerProviderDefaultClass
	_CPXKeyEventSequenceTrackerProviderDefaultClassOnce sync.Once
)

func getCPXKeyEventSequenceTrackerProviderDefaultClass() CPXKeyEventSequenceTrackerProviderDefaultClass {
	_CPXKeyEventSequenceTrackerProviderDefaultClassOnce.Do(func() {
		_CPXKeyEventSequenceTrackerProviderDefaultClass = CPXKeyEventSequenceTrackerProviderDefaultClass{class: objc.GetClass("CPXKeyEventSequenceTrackerProviderDefault")}
	})
	return _CPXKeyEventSequenceTrackerProviderDefaultClass
}

// GetCPXKeyEventSequenceTrackerProviderDefaultClass returns the class object for CPXKeyEventSequenceTrackerProviderDefault.
func GetCPXKeyEventSequenceTrackerProviderDefaultClass() CPXKeyEventSequenceTrackerProviderDefaultClass {
	return getCPXKeyEventSequenceTrackerProviderDefaultClass()
}

type CPXKeyEventSequenceTrackerProviderDefaultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyEventSequenceTrackerProviderDefaultClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyEventSequenceTrackerProviderDefaultClass) Alloc() CPXKeyEventSequenceTrackerProviderDefault {
	rv := objc.Send[CPXKeyEventSequenceTrackerProviderDefault](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXKeyEventSequenceTrackerProviderDefault.CurrentRegionID]
//   - [CPXKeyEventSequenceTrackerProviderDefault.EventLimit]
//   - [CPXKeyEventSequenceTrackerProviderDefault.MainDisplayHeight]
//   - [CPXKeyEventSequenceTrackerProviderDefault.StructuralRegionForID]
//   - [CPXKeyEventSequenceTrackerProviderDefault.WindowByID]
//   - [CPXKeyEventSequenceTrackerProviderDefault.WindowHeightForWindow]
//   - [CPXKeyEventSequenceTrackerProviderDefault.DebugDescription]
//   - [CPXKeyEventSequenceTrackerProviderDefault.Description]
//   - [CPXKeyEventSequenceTrackerProviderDefault.Hash]
//   - [CPXKeyEventSequenceTrackerProviderDefault.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault
type CPXKeyEventSequenceTrackerProviderDefault struct {
	objectivec.Object
}

// CPXKeyEventSequenceTrackerProviderDefaultFromID constructs a [CPXKeyEventSequenceTrackerProviderDefault] from an objc.ID.
func CPXKeyEventSequenceTrackerProviderDefaultFromID(id objc.ID) CPXKeyEventSequenceTrackerProviderDefault {
	return CPXKeyEventSequenceTrackerProviderDefault{objectivec.Object{ID: id}}
}

// Ensure CPXKeyEventSequenceTrackerProviderDefault implements ICPXKeyEventSequenceTrackerProviderDefault.
var _ ICPXKeyEventSequenceTrackerProviderDefault = CPXKeyEventSequenceTrackerProviderDefault{}

// An interface definition for the [CPXKeyEventSequenceTrackerProviderDefault] class.
//
// # Methods
//
//   - [ICPXKeyEventSequenceTrackerProviderDefault.CurrentRegionID]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.EventLimit]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.MainDisplayHeight]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.StructuralRegionForID]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.WindowByID]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.WindowHeightForWindow]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.DebugDescription]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.Description]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.Hash]
//   - [ICPXKeyEventSequenceTrackerProviderDefault.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault
type ICPXKeyEventSequenceTrackerProviderDefault interface {
	objectivec.IObject

	// Topic: Methods

	CurrentRegionID() uint64
	EventLimit() uint64
	MainDisplayHeight() uint16
	StructuralRegionForID(id uint64) WSStructuralRegionRef
	WindowByID(id uint32) unsafe.Pointer
	WindowHeightForWindow(window unsafe.Pointer) uint16
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXKeyEventSequenceTrackerProviderDefault) Init() CPXKeyEventSequenceTrackerProviderDefault {
	rv := objc.Send[CPXKeyEventSequenceTrackerProviderDefault](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyEventSequenceTrackerProviderDefault) Autorelease() CPXKeyEventSequenceTrackerProviderDefault {
	rv := objc.Send[CPXKeyEventSequenceTrackerProviderDefault](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyEventSequenceTrackerProviderDefault creates a new CPXKeyEventSequenceTrackerProviderDefault instance.
func NewCPXKeyEventSequenceTrackerProviderDefault() CPXKeyEventSequenceTrackerProviderDefault {
	class := getCPXKeyEventSequenceTrackerProviderDefaultClass()
	rv := objc.Send[CPXKeyEventSequenceTrackerProviderDefault](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/eventLimit
func (c CPXKeyEventSequenceTrackerProviderDefault) EventLimit() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("eventLimit"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/structuralRegionForID:
func (c CPXKeyEventSequenceTrackerProviderDefault) StructuralRegionForID(id uint64) WSStructuralRegionRef {
	rv := objc.Send[WSStructuralRegionRef](c.ID, objc.Sel("structuralRegionForID:"), id)
	return WSStructuralRegionRef(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/windowByID:
func (c CPXKeyEventSequenceTrackerProviderDefault) WindowByID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("windowByID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/windowHeightForWindow:
func (c CPXKeyEventSequenceTrackerProviderDefault) WindowHeightForWindow(window unsafe.Pointer) uint16 {
	rv := objc.Send[uint16](c.ID, objc.Sel("windowHeightForWindow:"), window)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/currentRegionID
func (c CPXKeyEventSequenceTrackerProviderDefault) CurrentRegionID() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("currentRegionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/debugDescription
func (c CPXKeyEventSequenceTrackerProviderDefault) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/description
func (c CPXKeyEventSequenceTrackerProviderDefault) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/hash
func (c CPXKeyEventSequenceTrackerProviderDefault) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/mainDisplayHeight
func (c CPXKeyEventSequenceTrackerProviderDefault) MainDisplayHeight() uint16 {
	rv := objc.Send[uint16](c.ID, objc.Sel("mainDisplayHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTrackerProviderDefault/superclass
func (c CPXKeyEventSequenceTrackerProviderDefault) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
