// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIDevice] class.
var (
	_MIDICIDeviceClass     MIDICIDeviceClass
	_MIDICIDeviceClassOnce sync.Once
)

func getMIDICIDeviceClass() MIDICIDeviceClass {
	_MIDICIDeviceClassOnce.Do(func() {
		_MIDICIDeviceClass = MIDICIDeviceClass{class: objc.GetClass("MIDICIDevice")}
	})
	return _MIDICIDeviceClass
}

// GetMIDICIDeviceClass returns the class object for MIDICIDevice.
func GetMIDICIDeviceClass() MIDICIDeviceClass {
	return getMIDICIDeviceClass()
}

type MIDICIDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIDeviceClass) Alloc() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDICIDevice.DeviceInfo]
//   - [MIDICIDevice.DeviceType]
//   - [MIDICIDevice.MaxPropertyExchangeRequests]
//   - [MIDICIDevice.MaxSysExSize]
//   - [MIDICIDevice.MUID]
//   - [MIDICIDevice.Profiles]
//   - [MIDICIDevice.SupportsProcessInquiry]
//   - [MIDICIDevice.SupportsProfileConfiguration]
//   - [MIDICIDevice.SupportsPropertyExchange]
//   - [MIDICIDevice.SupportsProtocolNegotiation]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice
type MIDICIDevice struct {
	objectivec.Object
}

// MIDICIDeviceFromID constructs a [MIDICIDevice] from an objc.ID.
func MIDICIDeviceFromID(id objc.ID) MIDICIDevice {
	return MIDICIDevice{objectivec.Object{ID: id}}
}

// NOTE: MIDICIDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIDevice] class.
//
// # Instance Properties
//
//   - [IMIDICIDevice.DeviceInfo]
//   - [IMIDICIDevice.DeviceType]
//   - [IMIDICIDevice.MaxPropertyExchangeRequests]
//   - [IMIDICIDevice.MaxSysExSize]
//   - [IMIDICIDevice.MUID]
//   - [IMIDICIDevice.Profiles]
//   - [IMIDICIDevice.SupportsProcessInquiry]
//   - [IMIDICIDevice.SupportsProfileConfiguration]
//   - [IMIDICIDevice.SupportsPropertyExchange]
//   - [IMIDICIDevice.SupportsProtocolNegotiation]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice
type IMIDICIDevice interface {
	objectivec.IObject

	// Topic: Instance Properties

	DeviceInfo() IMIDI2DeviceInfo
	DeviceType() MIDICIDeviceType
	MaxPropertyExchangeRequests() uint
	MaxSysExSize() uint
	MUID() MIDICIMUID
	Profiles() []MIDIUMPCIProfile
	SupportsProcessInquiry() bool
	SupportsProfileConfiguration() bool
	SupportsPropertyExchange() bool
	SupportsProtocolNegotiation() bool
}

// Init initializes the instance.
func (m MIDICIDevice) Init() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIDevice) Autorelease() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDevice creates a new MIDICIDevice instance.
func NewMIDICIDevice() MIDICIDevice {
	class := getMIDICIDeviceClass()
	rv := objc.Send[MIDICIDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/deviceInfo
func (m MIDICIDevice) DeviceInfo() IMIDI2DeviceInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deviceInfo"))
	return MIDI2DeviceInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/deviceType
func (m MIDICIDevice) DeviceType() MIDICIDeviceType {
	rv := objc.Send[MIDICIDeviceType](m.ID, objc.Sel("deviceType"))
	return MIDICIDeviceType(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/maxPropertyExchangeRequests
func (m MIDICIDevice) MaxPropertyExchangeRequests() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxPropertyExchangeRequests"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/maxSysExSize
func (m MIDICIDevice) MaxSysExSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxSysExSize"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/muid
func (m MIDICIDevice) MUID() MIDICIMUID {
	rv := objc.Send[MIDICIMUID](m.ID, objc.Sel("MUID"))
	return MIDICIMUID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/profiles
func (m MIDICIDevice) Profiles() []MIDIUMPCIProfile {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("profiles"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDIUMPCIProfile {
		return MIDIUMPCIProfileFromID(id)
	})
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProcessInquiry
func (m MIDICIDevice) SupportsProcessInquiry() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProcessInquiry"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProfileConfiguration
func (m MIDICIDevice) SupportsProfileConfiguration() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProfileConfiguration"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsPropertyExchange
func (m MIDICIDevice) SupportsPropertyExchange() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsPropertyExchange"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProtocolNegotiation
func (m MIDICIDevice) SupportsProtocolNegotiation() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProtocolNegotiation"))
	return rv
}
