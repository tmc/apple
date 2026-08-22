// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDIUMPMutableEndpoint] class.
var (
	_MIDIUMPMutableEndpointClass     MIDIUMPMutableEndpointClass
	_MIDIUMPMutableEndpointClassOnce sync.Once
)

func getMIDIUMPMutableEndpointClass() MIDIUMPMutableEndpointClass {
	_MIDIUMPMutableEndpointClassOnce.Do(func() {
		_MIDIUMPMutableEndpointClass = MIDIUMPMutableEndpointClass{class: objc.GetClass("MIDIUMPMutableEndpoint")}
	})
	return _MIDIUMPMutableEndpointClass
}

// GetMIDIUMPMutableEndpointClass returns the class object for MIDIUMPMutableEndpoint.
func GetMIDIUMPMutableEndpointClass() MIDIUMPMutableEndpointClass {
	return getMIDIUMPMutableEndpointClass()
}

type MIDIUMPMutableEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPMutableEndpointClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPMutableEndpointClass) Alloc() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MIDIUMPMutableEndpoint.InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback]
//
// # Instance Properties
//
//   - [MIDIUMPMutableEndpoint.IsEnabled]
//   - [MIDIUMPMutableEndpoint.MutableFunctionBlocks]
//   - [MIDIUMPMutableEndpoint.SetMutableFunctionBlocks]
//
// # Instance Methods
//
//   - [MIDIUMPMutableEndpoint.RegisterFunctionBlocksMarkAsStaticError]
//   - [MIDIUMPMutableEndpoint.SetEnabledError]
//   - [MIDIUMPMutableEndpoint.SetNameError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint
type MIDIUMPMutableEndpoint struct {
	MIDIUMPEndpoint
}

// MIDIUMPMutableEndpointFromID constructs a [MIDIUMPMutableEndpoint] from an objc.ID.
func MIDIUMPMutableEndpointFromID(id objc.ID) MIDIUMPMutableEndpoint {
	return MIDIUMPMutableEndpoint{MIDIUMPEndpoint: MIDIUMPEndpointFromID(id)}
}

// NOTE: MIDIUMPMutableEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPMutableEndpoint] class.
//
// # Initializers
//
//   - [IMIDIUMPMutableEndpoint.InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback]
//
// # Instance Properties
//
//   - [IMIDIUMPMutableEndpoint.IsEnabled]
//   - [IMIDIUMPMutableEndpoint.MutableFunctionBlocks]
//   - [IMIDIUMPMutableEndpoint.SetMutableFunctionBlocks]
//
// # Instance Methods
//
//   - [IMIDIUMPMutableEndpoint.RegisterFunctionBlocksMarkAsStaticError]
//   - [IMIDIUMPMutableEndpoint.SetEnabledError]
//   - [IMIDIUMPMutableEndpoint.SetNameError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint
type IMIDIUMPMutableEndpoint interface {
	IMIDIUMPEndpoint

	// Topic: Initializers

	InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name string, deviceInfo IMIDI2DeviceInfo, productInstanceID string, MIDIProtocol MIDIProtocolID, destinationCallback MIDIEventListUnsafePointerHandler) MIDIUMPMutableEndpoint

	// Topic: Instance Properties

	IsEnabled() bool
	MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock
	SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock)

	// Topic: Instance Methods

	RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool) (bool, error)
	SetEnabledError(isEnabled bool) (bool, error)
	SetNameError(name string) (bool, error)
}

// Init initializes the instance.
func (m MIDIUMPMutableEndpoint) Init() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPMutableEndpoint) Autorelease() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPMutableEndpoint creates a new MIDIUMPMutableEndpoint instance.
func NewMIDIUMPMutableEndpoint() MIDIUMPMutableEndpoint {
	class := getMIDIUMPMutableEndpointClass()
	rv := objc.Send[MIDIUMPMutableEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/init(name:deviceInfo:productInstanceID:midiProtocol:destinationCallback:)
func NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name string, deviceInfo IMIDI2DeviceInfo, productInstanceID string, MIDIProtocol MIDIProtocolID, destinationCallback MIDIEventListUnsafePointerHandler) MIDIUMPMutableEndpoint {
	_block4, _ := NewMIDIEventListUnsafePointerBlock(destinationCallback)
	instance := getMIDIUMPMutableEndpointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:deviceInfo:productInstanceID:MIDIProtocol:destinationCallback:"), objc.String(name), deviceInfo, objc.String(productInstanceID), MIDIProtocol, _block4)
	return MIDIUMPMutableEndpointFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/init(name:deviceInfo:productInstanceID:midiProtocol:destinationCallback:)
func (m MIDIUMPMutableEndpoint) InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name string, deviceInfo IMIDI2DeviceInfo, productInstanceID string, MIDIProtocol MIDIProtocolID, destinationCallback MIDIEventListUnsafePointerHandler) MIDIUMPMutableEndpoint {
	_block4, _ := NewMIDIEventListUnsafePointerBlock(destinationCallback)
	rv := objc.Send[MIDIUMPMutableEndpoint](m.ID, objc.Sel("initWithName:deviceInfo:productInstanceID:MIDIProtocol:destinationCallback:"), objc.String(name), deviceInfo, objc.String(productInstanceID), MIDIProtocol, _block4)
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/registerFunctionBlocks(_:markAsStatic:)
func (m MIDIUMPMutableEndpoint) RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("registerFunctionBlocks:markAsStatic:error:"), objectivec.IObjectSliceToNSArray(functionBlocks), markAsStatic, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("registerFunctionBlocks:markAsStatic:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setEnabled(_:)
func (m MIDIUMPMutableEndpoint) SetEnabledError(isEnabled bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setEnabled:error:"), isEnabled, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEnabled:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setName(_:)
func (m MIDIUMPMutableEndpoint) SetNameError(name string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setName:error:"), objc.String(name), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/isEnabled
func (m MIDIUMPMutableEndpoint) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/mutableFunctionBlocks
func (m MIDIUMPMutableEndpoint) MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("mutableFunctionBlocks"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDIUMPMutableFunctionBlock {
		return MIDIUMPMutableFunctionBlockFromID(id)
	})
}
func (m MIDIUMPMutableEndpoint) SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock) {
	objc.Send[struct{}](m.ID, objc.Sel("setMutableFunctionBlocks:"), objectivec.IObjectSliceToNSArray(value))
}
