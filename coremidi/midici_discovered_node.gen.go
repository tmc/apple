// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIDiscoveredNode] class.
var (
	_MIDICIDiscoveredNodeClass     MIDICIDiscoveredNodeClass
	_MIDICIDiscoveredNodeClassOnce sync.Once
)

func getMIDICIDiscoveredNodeClass() MIDICIDiscoveredNodeClass {
	_MIDICIDiscoveredNodeClassOnce.Do(func() {
		_MIDICIDiscoveredNodeClass = MIDICIDiscoveredNodeClass{class: objc.GetClass("MIDICIDiscoveredNode")}
	})
	return _MIDICIDiscoveredNodeClass
}

// GetMIDICIDiscoveredNodeClass returns the class object for MIDICIDiscoveredNode.
func GetMIDICIDiscoveredNodeClass() MIDICIDiscoveredNodeClass {
	return getMIDICIDiscoveredNodeClass()
}

type MIDICIDiscoveredNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIDiscoveredNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIDiscoveredNodeClass) Alloc() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A discovered MIDI-CI node that represents a MIDI source and destination
// that respond to capability inquiries.
//
// # Inspecting a Node
//
//   - [MIDICIDiscoveredNode.Destination]: The node’s MIDI destination.
//   - [MIDICIDiscoveredNode.DeviceInfo]: The available MIDI-CI device information.
//   - [MIDICIDiscoveredNode.SupportsProfiles]: A Boolean value that indicates whether this node supports MIDI-CI profiles.
//   - [MIDICIDiscoveredNode.SupportsProperties]: A Boolean value that indicates whether this node supports MIDI-CI properties.
//   - [MIDICIDiscoveredNode.MaximumSysExSize]: The maximum size of a System Exclusive (SysEx) message this node supports.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode
type MIDICIDiscoveredNode struct {
	objectivec.Object
}

// MIDICIDiscoveredNodeFromID constructs a [MIDICIDiscoveredNode] from an objc.ID.
//
// A discovered MIDI-CI node that represents a MIDI source and destination
// that respond to capability inquiries.
func MIDICIDiscoveredNodeFromID(id objc.ID) MIDICIDiscoveredNode {
	return MIDICIDiscoveredNode{objectivec.Object{ID: id}}
}

// NOTE: MIDICIDiscoveredNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIDiscoveredNode] class.
//
// # Inspecting a Node
//
//   - [IMIDICIDiscoveredNode.Destination]: The node’s MIDI destination.
//   - [IMIDICIDiscoveredNode.DeviceInfo]: The available MIDI-CI device information.
//   - [IMIDICIDiscoveredNode.SupportsProfiles]: A Boolean value that indicates whether this node supports MIDI-CI profiles.
//   - [IMIDICIDiscoveredNode.SupportsProperties]: A Boolean value that indicates whether this node supports MIDI-CI properties.
//   - [IMIDICIDiscoveredNode.MaximumSysExSize]: The maximum size of a System Exclusive (SysEx) message this node supports.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode
type IMIDICIDiscoveredNode interface {
	objectivec.IObject

	// Topic: Inspecting a Node

	// The node’s MIDI destination.
	Destination() MIDIEntityRef
	// The available MIDI-CI device information.
	DeviceInfo() IMIDICIDeviceInfo
	// A Boolean value that indicates whether this node supports MIDI-CI profiles.
	SupportsProfiles() bool
	// A Boolean value that indicates whether this node supports MIDI-CI properties.
	SupportsProperties() bool
	// The maximum size of a System Exclusive (SysEx) message this node supports.
	MaximumSysExSize() foundation.NSNumber

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MIDICIDiscoveredNode) Init() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIDiscoveredNode) Autorelease() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveredNode creates a new MIDICIDiscoveredNode instance.
func NewMIDICIDiscoveredNode() MIDICIDiscoveredNode {
	class := getMIDICIDiscoveredNodeClass()
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MIDICIDiscoveredNode) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The node’s MIDI destination.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/destination
func (m MIDICIDiscoveredNode) Destination() MIDIEntityRef {
	rv := objc.Send[MIDIEntityRef](m.ID, objc.Sel("destination"))
	return MIDIEntityRef(rv)
}

// The available MIDI-CI device information.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/deviceInfo
func (m MIDICIDiscoveredNode) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deviceInfo"))
	return MIDICIDeviceInfoFromID(objc.ID(rv))
}

// A Boolean value that indicates whether this node supports MIDI-CI profiles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProfiles
func (m MIDICIDiscoveredNode) SupportsProfiles() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProfiles"))
	return rv
}

// A Boolean value that indicates whether this node supports MIDI-CI
// properties.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProperties
func (m MIDICIDiscoveredNode) SupportsProperties() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProperties"))
	return rv
}

// The maximum size of a System Exclusive (SysEx) message this node supports.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/maximumSysExSize
func (m MIDICIDiscoveredNode) MaximumSysExSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("maximumSysExSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
