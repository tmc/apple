// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDIUMPEndpoint] class.
var (
	_MIDIUMPEndpointClass     MIDIUMPEndpointClass
	_MIDIUMPEndpointClassOnce sync.Once
)

func getMIDIUMPEndpointClass() MIDIUMPEndpointClass {
	_MIDIUMPEndpointClassOnce.Do(func() {
		_MIDIUMPEndpointClass = MIDIUMPEndpointClass{class: objc.GetClass("MIDIUMPEndpoint")}
	})
	return _MIDIUMPEndpointClass
}

// GetMIDIUMPEndpointClass returns the class object for MIDIUMPEndpoint.
func GetMIDIUMPEndpointClass() MIDIUMPEndpointClass {
	return getMIDIUMPEndpointClass()
}

type MIDIUMPEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPEndpointClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPEndpointClass) Alloc() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDIUMPEndpoint.DeviceInfo]
//   - [MIDIUMPEndpoint.EndpointType]
//   - [MIDIUMPEndpoint.FunctionBlocks]
//   - [MIDIUMPEndpoint.SetFunctionBlocks]
//   - [MIDIUMPEndpoint.HasJRTSReceiveCapability]
//   - [MIDIUMPEndpoint.HasJRTSTransmitCapability]
//   - [MIDIUMPEndpoint.HasStaticFunctionBlocks]
//   - [MIDIUMPEndpoint.MIDIDestination]
//   - [MIDIUMPEndpoint.MIDIProtocol]
//   - [MIDIUMPEndpoint.MIDISource]
//   - [MIDIUMPEndpoint.Name]
//   - [MIDIUMPEndpoint.ProductInstanceID]
//   - [MIDIUMPEndpoint.SupportedMIDIProtocols]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint
type MIDIUMPEndpoint struct {
	objectivec.Object
}

// MIDIUMPEndpointFromID constructs a [MIDIUMPEndpoint] from an objc.ID.
func MIDIUMPEndpointFromID(id objc.ID) MIDIUMPEndpoint {
	return MIDIUMPEndpoint{objectivec.Object{ID: id}}
}

// NOTE: MIDIUMPEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPEndpoint] class.
//
// # Instance Properties
//
//   - [IMIDIUMPEndpoint.DeviceInfo]
//   - [IMIDIUMPEndpoint.EndpointType]
//   - [IMIDIUMPEndpoint.FunctionBlocks]
//   - [IMIDIUMPEndpoint.SetFunctionBlocks]
//   - [IMIDIUMPEndpoint.HasJRTSReceiveCapability]
//   - [IMIDIUMPEndpoint.HasJRTSTransmitCapability]
//   - [IMIDIUMPEndpoint.HasStaticFunctionBlocks]
//   - [IMIDIUMPEndpoint.MIDIDestination]
//   - [IMIDIUMPEndpoint.MIDIProtocol]
//   - [IMIDIUMPEndpoint.MIDISource]
//   - [IMIDIUMPEndpoint.Name]
//   - [IMIDIUMPEndpoint.ProductInstanceID]
//   - [IMIDIUMPEndpoint.SupportedMIDIProtocols]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint
type IMIDIUMPEndpoint interface {
	objectivec.IObject

	// Topic: Instance Properties

	DeviceInfo() IMIDI2DeviceInfo
	EndpointType() MIDIUMPCIObjectBackingType
	FunctionBlocks() []MIDIUMPFunctionBlock
	SetFunctionBlocks(value []MIDIUMPFunctionBlock)
	HasJRTSReceiveCapability() bool
	HasJRTSTransmitCapability() bool
	HasStaticFunctionBlocks() bool
	MIDIDestination() MIDIEndpointRef
	MIDIProtocol() MIDIProtocolID
	MIDISource() MIDIEndpointRef
	Name() string
	ProductInstanceID() string
	SupportedMIDIProtocols() MIDIUMPProtocolOptions
}

// Init initializes the instance.
func (m MIDIUMPEndpoint) Init() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPEndpoint) Autorelease() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpoint creates a new MIDIUMPEndpoint instance.
func NewMIDIUMPEndpoint() MIDIUMPEndpoint {
	class := getMIDIUMPEndpointClass()
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/deviceInfo
func (m MIDIUMPEndpoint) DeviceInfo() IMIDI2DeviceInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deviceInfo"))
	return MIDI2DeviceInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/endpointType
func (m MIDIUMPEndpoint) EndpointType() MIDIUMPCIObjectBackingType {
	rv := objc.Send[MIDIUMPCIObjectBackingType](m.ID, objc.Sel("endpointType"))
	return MIDIUMPCIObjectBackingType(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/functionBlocks
func (m MIDIUMPEndpoint) FunctionBlocks() []MIDIUMPFunctionBlock {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("functionBlocks"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDIUMPFunctionBlock {
		return MIDIUMPFunctionBlockFromID(id)
	})
}
func (m MIDIUMPEndpoint) SetFunctionBlocks(value []MIDIUMPFunctionBlock) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionBlocks:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSReceiveCapability
func (m MIDIUMPEndpoint) HasJRTSReceiveCapability() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasJRTSReceiveCapability"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSTransmitCapability
func (m MIDIUMPEndpoint) HasJRTSTransmitCapability() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasJRTSTransmitCapability"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasStaticFunctionBlocks
func (m MIDIUMPEndpoint) HasStaticFunctionBlocks() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasStaticFunctionBlocks"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiDestination
func (m MIDIUMPEndpoint) MIDIDestination() MIDIEndpointRef {
	rv := objc.Send[MIDIEndpointRef](m.ID, objc.Sel("MIDIDestination"))
	return MIDIEndpointRef(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiProtocol
func (m MIDIUMPEndpoint) MIDIProtocol() MIDIProtocolID {
	rv := objc.Send[MIDIProtocolID](m.ID, objc.Sel("MIDIProtocol"))
	return MIDIProtocolID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiSource
func (m MIDIUMPEndpoint) MIDISource() MIDIEndpointRef {
	rv := objc.Send[MIDIEndpointRef](m.ID, objc.Sel("MIDISource"))
	return MIDIEndpointRef(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/name
func (m MIDIUMPEndpoint) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/productInstanceID
func (m MIDIUMPEndpoint) ProductInstanceID() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("productInstanceID"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/supportedMIDIProtocols
func (m MIDIUMPEndpoint) SupportedMIDIProtocols() MIDIUMPProtocolOptions {
	rv := objc.Send[MIDIUMPProtocolOptions](m.ID, objc.Sel("supportedMIDIProtocols"))
	return MIDIUMPProtocolOptions(rv)
}
