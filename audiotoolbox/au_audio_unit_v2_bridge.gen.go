// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AUAudioUnitV2Bridge] class.
var (
	_AUAudioUnitV2BridgeClass     AUAudioUnitV2BridgeClass
	_AUAudioUnitV2BridgeClassOnce sync.Once
)

func getAUAudioUnitV2BridgeClass() AUAudioUnitV2BridgeClass {
	_AUAudioUnitV2BridgeClassOnce.Do(func() {
		_AUAudioUnitV2BridgeClass = AUAudioUnitV2BridgeClass{class: objc.GetClass("AUAudioUnitV2Bridge")}
	})
	return _AUAudioUnitV2BridgeClass
}

// GetAUAudioUnitV2BridgeClass returns the class object for AUAudioUnitV2Bridge.
func GetAUAudioUnitV2BridgeClass() AUAudioUnitV2BridgeClass {
	return getAUAudioUnitV2BridgeClass()
}

type AUAudioUnitV2BridgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUAudioUnitV2BridgeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUAudioUnitV2BridgeClass) Alloc() AUAudioUnitV2Bridge {
	rv := objc.Send[AUAudioUnitV2Bridge](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that wraps a version 2 audio unit as version 3 audio unit.
//
// # Overview
//
// A version 3 audio unit may subclass the [AUAudioUnitV2Bridge] class. If so,
// the audio unit’s component description should refer to a registered
// component with a version 2 implementation by using a factory function. The
// bridge will instantiate the version 2 audio unit via the factory function
// and communicate with it using version 2 audio unit APIs.
//
// Hosts should not access this class; it will be instantiated if needed when
// creating an audio unit.
//
// # Instance Properties
//
//   - [AUAudioUnitV2Bridge.AudioUnit]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitV2Bridge
type AUAudioUnitV2Bridge struct {
	AUAudioUnit
}

// AUAudioUnitV2BridgeFromID constructs a [AUAudioUnitV2Bridge] from an objc.ID.
//
// A class that wraps a version 2 audio unit as version 3 audio unit.
func AUAudioUnitV2BridgeFromID(id objc.ID) AUAudioUnitV2Bridge {
	return AUAudioUnitV2Bridge{AUAudioUnit: AUAudioUnitFromID(id)}
}

// NOTE: AUAudioUnitV2Bridge adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUAudioUnitV2Bridge] class.
//
// # Instance Properties
//
//   - [IAUAudioUnitV2Bridge.AudioUnit]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitV2Bridge
type IAUAudioUnitV2Bridge interface {
	IAUAudioUnit

	// Topic: Instance Properties

	AudioUnit() AudioUnit
}

// Init initializes the instance.
func (a AUAudioUnitV2Bridge) Init() AUAudioUnitV2Bridge {
	rv := objc.Send[AUAudioUnitV2Bridge](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AUAudioUnitV2Bridge) Autorelease() AUAudioUnitV2Bridge {
	rv := objc.Send[AUAudioUnitV2Bridge](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUAudioUnitV2Bridge creates a new AUAudioUnitV2Bridge instance.
func NewAUAudioUnitV2Bridge() AUAudioUnitV2Bridge {
	class := getAUAudioUnitV2BridgeClass()
	rv := objc.Send[AUAudioUnitV2Bridge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the convenience initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:)
func NewAudioUnitV2BridgeWithComponentDescriptionError(componentDescription AudioComponentDescription) (AUAudioUnitV2Bridge, error) {
	var errorPtr objc.ID
	instance := getAUAudioUnitV2BridgeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComponentDescription:error:"), componentDescription, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnitV2Bridge{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AUAudioUnitV2Bridge{}, objc.ErrInitFailed
	}
	return AUAudioUnitV2BridgeFromID(rv), nil
}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// options: Options for loading the unit in-process or out-of-process.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the designated initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:options:)
func NewAudioUnitV2BridgeWithComponentDescriptionOptionsError(componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions) (AUAudioUnitV2Bridge, error) {
	var errorPtr objc.ID
	instance := getAUAudioUnitV2BridgeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComponentDescription:options:error:"), componentDescription, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnitV2Bridge{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AUAudioUnitV2Bridge{}, objc.ErrInitFailed
	}
	return AUAudioUnitV2BridgeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitV2Bridge/audioUnit
func (a AUAudioUnitV2Bridge) AudioUnit() AudioUnit {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioUnit"))
	return AudioUnit(rv)
}
