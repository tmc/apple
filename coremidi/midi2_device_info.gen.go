// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDI2DeviceInfo] class.
var (
	_MIDI2DeviceInfoClass     MIDI2DeviceInfoClass
	_MIDI2DeviceInfoClassOnce sync.Once
)

func getMIDI2DeviceInfoClass() MIDI2DeviceInfoClass {
	_MIDI2DeviceInfoClassOnce.Do(func() {
		_MIDI2DeviceInfoClass = MIDI2DeviceInfoClass{class: objc.GetClass("MIDI2DeviceInfo")}
	})
	return _MIDI2DeviceInfoClass
}

// GetMIDI2DeviceInfoClass returns the class object for MIDI2DeviceInfo.
func GetMIDI2DeviceInfoClass() MIDI2DeviceInfoClass {
	return getMIDI2DeviceInfoClass()
}

type MIDI2DeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDI2DeviceInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDI2DeviceInfoClass) Alloc() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MIDI2DeviceInfo.InitWithManufacturerIDFamilyModelNumberRevisionLevel]
//
// # Instance Properties
//
//   - [MIDI2DeviceInfo.Family]
//   - [MIDI2DeviceInfo.ManufacturerID]
//   - [MIDI2DeviceInfo.ModelNumber]
//   - [MIDI2DeviceInfo.RevisionLevel]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo
type MIDI2DeviceInfo struct {
	objectivec.Object
}

// MIDI2DeviceInfoFromID constructs a [MIDI2DeviceInfo] from an objc.ID.
func MIDI2DeviceInfoFromID(id objc.ID) MIDI2DeviceInfo {
	return MIDI2DeviceInfo{objectivec.Object{ID: id}}
}

// NOTE: MIDI2DeviceInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDI2DeviceInfo] class.
//
// # Initializers
//
//   - [IMIDI2DeviceInfo.InitWithManufacturerIDFamilyModelNumberRevisionLevel]
//
// # Instance Properties
//
//   - [IMIDI2DeviceInfo.Family]
//   - [IMIDI2DeviceInfo.ManufacturerID]
//   - [IMIDI2DeviceInfo.ModelNumber]
//   - [IMIDI2DeviceInfo.RevisionLevel]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo
type IMIDI2DeviceInfo interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithManufacturerIDFamilyModelNumberRevisionLevel(manufacturerID MIDI2DeviceManufacturer, family MIDIUInteger14, modelNumber MIDIUInteger14, revisionLevel MIDI2DeviceRevisionLevel) MIDI2DeviceInfo

	// Topic: Instance Properties

	Family() MIDIUInteger14
	ManufacturerID() MIDI2DeviceManufacturer
	ModelNumber() MIDIUInteger14
	RevisionLevel() MIDI2DeviceRevisionLevel
}

// Init initializes the instance.
func (m MIDI2DeviceInfo) Init() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDI2DeviceInfo) Autorelease() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDI2DeviceInfo creates a new MIDI2DeviceInfo instance.
func NewMIDI2DeviceInfo() MIDI2DeviceInfo {
	class := getMIDI2DeviceInfoClass()
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/init(manufacturerID:family:modelNumber:revisionLevel:)
func NewMIDI2DeviceInfoWithManufacturerIDFamilyModelNumberRevisionLevel(manufacturerID MIDI2DeviceManufacturer, family MIDIUInteger14, modelNumber MIDIUInteger14, revisionLevel MIDI2DeviceRevisionLevel) MIDI2DeviceInfo {
	instance := getMIDI2DeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithManufacturerID:family:modelNumber:revisionLevel:"), manufacturerID, family, modelNumber, revisionLevel)
	return MIDI2DeviceInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/init(manufacturerID:family:modelNumber:revisionLevel:)
func (m MIDI2DeviceInfo) InitWithManufacturerIDFamilyModelNumberRevisionLevel(manufacturerID MIDI2DeviceManufacturer, family MIDIUInteger14, modelNumber MIDIUInteger14, revisionLevel MIDI2DeviceRevisionLevel) MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m.ID, objc.Sel("initWithManufacturerID:family:modelNumber:revisionLevel:"), manufacturerID, family, modelNumber, revisionLevel)
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/family
func (m MIDI2DeviceInfo) Family() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m.ID, objc.Sel("family"))
	return MIDIUInteger14(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/manufacturerID
func (m MIDI2DeviceInfo) ManufacturerID() MIDI2DeviceManufacturer {
	rv := objc.Send[MIDI2DeviceManufacturer](m.ID, objc.Sel("manufacturerID"))
	return MIDI2DeviceManufacturer(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/modelNumber
func (m MIDI2DeviceInfo) ModelNumber() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m.ID, objc.Sel("modelNumber"))
	return MIDIUInteger14(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/revisionLevel
func (m MIDI2DeviceInfo) RevisionLevel() MIDI2DeviceRevisionLevel {
	rv := objc.Send[MIDI2DeviceRevisionLevel](m.ID, objc.Sel("revisionLevel"))
	return MIDI2DeviceRevisionLevel(rv)
}
