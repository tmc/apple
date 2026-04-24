// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWindowManagerSpace] class.
var (
	_SLSWindowManagerSpaceClass     SLSWindowManagerSpaceClass
	_SLSWindowManagerSpaceClassOnce sync.Once
)

func getSLSWindowManagerSpaceClass() SLSWindowManagerSpaceClass {
	_SLSWindowManagerSpaceClassOnce.Do(func() {
		_SLSWindowManagerSpaceClass = SLSWindowManagerSpaceClass{class: objc.GetClass("SLSWindowManagerSpace")}
	})
	return _SLSWindowManagerSpaceClass
}

// GetSLSWindowManagerSpaceClass returns the class object for SLSWindowManagerSpace.
func GetSLSWindowManagerSpaceClass() SLSWindowManagerSpaceClass {
	return getSLSWindowManagerSpaceClass()
}

type SLSWindowManagerSpaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWindowManagerSpaceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWindowManagerSpaceClass) Alloc() SLSWindowManagerSpace {
	rv := objc.Send[SLSWindowManagerSpace](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWindowManagerSpace._effectiveDisplayID]
//   - [SLSWindowManagerSpace.DisplayUUID]
//   - [SLSWindowManagerSpace.SetDisplayUUID]
//   - [SLSWindowManagerSpace.IsCurrentSpace]
//   - [SLSWindowManagerSpace.IsManagedSpace]
//   - [SLSWindowManagerSpace.Manager]
//   - [SLSWindowManagerSpace.SetManager]
//   - [SLSWindowManagerSpace.SpaceID]
//   - [SLSWindowManagerSpace.SetSpaceID]
//   - [SLSWindowManagerSpace.Type]
//   - [SLSWindowManagerSpace.SetType]
//   - [SLSWindowManagerSpace.WindowIDs]
//   - [SLSWindowManagerSpace.SetWindowIDs]
//   - [SLSWindowManagerSpace.CurrentSpace]
//   - [SLSWindowManagerSpace.ManagedSpace]
//   - [SLSWindowManagerSpace.SetManagedSpace]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace
type SLSWindowManagerSpace struct {
	objectivec.Object
}

// SLSWindowManagerSpaceFromID constructs a [SLSWindowManagerSpace] from an objc.ID.
func SLSWindowManagerSpaceFromID(id objc.ID) SLSWindowManagerSpace {
	return SLSWindowManagerSpace{objectivec.Object{ID: id}}
}

// Ensure SLSWindowManagerSpace implements ISLSWindowManagerSpace.
var _ ISLSWindowManagerSpace = SLSWindowManagerSpace{}

// An interface definition for the [SLSWindowManagerSpace] class.
//
// # Methods
//
//   - [ISLSWindowManagerSpace._effectiveDisplayID]
//   - [ISLSWindowManagerSpace.DisplayUUID]
//   - [ISLSWindowManagerSpace.SetDisplayUUID]
//   - [ISLSWindowManagerSpace.IsCurrentSpace]
//   - [ISLSWindowManagerSpace.IsManagedSpace]
//   - [ISLSWindowManagerSpace.Manager]
//   - [ISLSWindowManagerSpace.SetManager]
//   - [ISLSWindowManagerSpace.SpaceID]
//   - [ISLSWindowManagerSpace.SetSpaceID]
//   - [ISLSWindowManagerSpace.Type]
//   - [ISLSWindowManagerSpace.SetType]
//   - [ISLSWindowManagerSpace.WindowIDs]
//   - [ISLSWindowManagerSpace.SetWindowIDs]
//   - [ISLSWindowManagerSpace.CurrentSpace]
//   - [ISLSWindowManagerSpace.ManagedSpace]
//   - [ISLSWindowManagerSpace.SetManagedSpace]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace
type ISLSWindowManagerSpace interface {
	objectivec.IObject

	// Topic: Methods

	_effectiveDisplayID() objectivec.IObject
	DisplayUUID() string
	SetDisplayUUID(value string)
	IsCurrentSpace() bool
	IsManagedSpace() bool
	Manager() ISLSSpaceWindowManager
	SetManager(value ISLSSpaceWindowManager)
	SpaceID() uint64
	SetSpaceID(value uint64)
	Type() int
	SetType(value int)
	WindowIDs() foundation.INSSet
	SetWindowIDs(value foundation.INSSet)
	CurrentSpace() bool
	ManagedSpace() bool
	SetManagedSpace(value bool)
}

// Init initializes the instance.
func (s SLSWindowManagerSpace) Init() SLSWindowManagerSpace {
	rv := objc.Send[SLSWindowManagerSpace](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWindowManagerSpace) Autorelease() SLSWindowManagerSpace {
	rv := objc.Send[SLSWindowManagerSpace](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWindowManagerSpace creates a new SLSWindowManagerSpace instance.
func NewSLSWindowManagerSpace() SLSWindowManagerSpace {
	class := getSLSWindowManagerSpaceClass()
	rv := objc.Send[SLSWindowManagerSpace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/_effectiveDisplayID
func (s SLSWindowManagerSpace) _effectiveDisplayID() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_effectiveDisplayID"))
	return objectivec.Object{ID: rv}
}

// EffectiveDisplayID is an exported wrapper for the private method _effectiveDisplayID.
func (s SLSWindowManagerSpace) EffectiveDisplayID() objectivec.IObject {
	return s._effectiveDisplayID()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/isCurrentSpace
func (s SLSWindowManagerSpace) IsCurrentSpace() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isCurrentSpace"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/isManagedSpace
func (s SLSWindowManagerSpace) IsManagedSpace() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isManagedSpace"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/currentSpace
func (s SLSWindowManagerSpace) CurrentSpace() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("currentSpace"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/displayUUID
func (s SLSWindowManagerSpace) DisplayUUID() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayUUID"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSWindowManagerSpace) SetDisplayUUID(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayUUID:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/managedSpace
func (s SLSWindowManagerSpace) ManagedSpace() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("managedSpace"))
	return rv
}
func (s SLSWindowManagerSpace) SetManagedSpace(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setManagedSpace:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/manager
func (s SLSWindowManagerSpace) Manager() ISLSSpaceWindowManager {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("manager"))
	return SLSSpaceWindowManagerFromID(objc.ID(rv))
}
func (s SLSWindowManagerSpace) SetManager(value ISLSSpaceWindowManager) {
	objc.Send[struct{}](s.ID, objc.Sel("setManager:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/spaceID
func (s SLSWindowManagerSpace) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
func (s SLSWindowManagerSpace) SetSpaceID(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpaceID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/type
func (s SLSWindowManagerSpace) Type() int {
	rv := objc.Send[int](s.ID, objc.Sel("type"))
	return rv
}
func (s SLSWindowManagerSpace) SetType(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setType:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerSpace/windowIDs
func (s SLSWindowManagerSpace) WindowIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windowIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLSWindowManagerSpace) SetWindowIDs(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setWindowIDs:"), value)
}
