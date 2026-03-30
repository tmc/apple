// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSSound](<doc://com.apple.appkit/documentation/AppKit/NSSound>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSSoundDelegate
type NSSoundDelegate interface {
	objectivec.IObject
}

// NSSoundDelegateObject wraps an existing Objective-C object that conforms to the NSSoundDelegate protocol.
type NSSoundDelegateObject struct {
	objectivec.Object
}

func (o NSSoundDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSoundDelegateObjectFromID constructs a [NSSoundDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSoundDelegateObjectFromID(id objc.ID) NSSoundDelegateObject {
	return NSSoundDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// This delegate method is called when an [NSSound] instance has completed
// playback of its sound data.
//
// sound: The [NSSound] that has completed playback of its sound data.
//
// flag: true when playback was successful; false otherwise.
//
// See: https://developer.apple.com/documentation/AppKit/NSSoundDelegate/sound(_:didFinishPlaying:)
func (o NSSoundDelegateObject) SoundDidFinishPlaying(sound INSSound, flag bool) {
	objc.Send[struct{}](o.ID, objc.Sel("sound:didFinishPlaying:"), sound, flag)
}

// NSSoundDelegateConfig holds optional typed callbacks for [NSSoundDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssounddelegate
type NSSoundDelegateConfig struct {

	// Playing Sounds
	// SoundDidFinishPlaying — This delegate method is called when an [NSSound] instance has completed playback of its sound data.
	SoundDidFinishPlaying func(sound NSSound, flag bool)
}

// NewNSSoundDelegate creates an Objective-C object implementing the [NSSoundDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSoundDelegateObject] satisfies the [NSSoundDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssounddelegate
func NewNSSoundDelegate(config NSSoundDelegateConfig) NSSoundDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSoundDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SoundDidFinishPlaying != nil {
		fn := config.SoundDidFinishPlaying
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sound:didFinishPlaying:"),
			Fn: func(self objc.ID, _cmd objc.SEL, soundID objc.ID, flag bool) {
				sound := NSSoundFromID(soundID)
				fn(sound, flag)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSoundDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSoundDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSoundDelegateObjectFromID(instance)
}
