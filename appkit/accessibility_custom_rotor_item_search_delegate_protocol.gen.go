// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A delegate for a custom rotor that finds the next item result after performing a search with the specified search parameters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotorItemSearchDelegate
type NSAccessibilityCustomRotorItemSearchDelegate interface {
	objectivec.IObject

	// Performs a search with the specified search parameters and returns the item result.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotorItemSearchDelegate/rotor(_:resultFor:)
	RotorResultForSearchParameters(rotor INSAccessibilityCustomRotor, searchParameters INSAccessibilityCustomRotorSearchParameters) INSAccessibilityCustomRotorItemResult
}

// NSAccessibilityCustomRotorItemSearchDelegateObject wraps an existing Objective-C object that conforms to the NSAccessibilityCustomRotorItemSearchDelegate protocol.
type NSAccessibilityCustomRotorItemSearchDelegateObject struct {
	objectivec.Object
}
func (o NSAccessibilityCustomRotorItemSearchDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityCustomRotorItemSearchDelegateObjectFromID constructs a [NSAccessibilityCustomRotorItemSearchDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityCustomRotorItemSearchDelegateObjectFromID(id objc.ID) NSAccessibilityCustomRotorItemSearchDelegateObject {
	return NSAccessibilityCustomRotorItemSearchDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Performs a search with the specified search parameters and returns the item
// result.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotorItemSearchDelegate/rotor(_:resultFor:)

func (o NSAccessibilityCustomRotorItemSearchDelegateObject) RotorResultForSearchParameters(rotor INSAccessibilityCustomRotor, searchParameters INSAccessibilityCustomRotorSearchParameters) INSAccessibilityCustomRotorItemResult {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("rotor:resultForSearchParameters:"), rotor, searchParameters)
	return NSAccessibilityCustomRotorItemResultFromID(rv)
	}

// NSAccessibilityCustomRotorItemSearchDelegateConfig holds optional typed callbacks for [NSAccessibilityCustomRotorItemSearchDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate
type NSAccessibilityCustomRotorItemSearchDelegateConfig struct {

	// Other Methods
	// RotorResultForSearchParameters — Performs a search with the specified search parameters and returns the item result.
	RotorResultForSearchParameters func(rotor NSAccessibilityCustomRotor, searchParameters NSAccessibilityCustomRotorSearchParameters) NSAccessibilityCustomRotorItemResult
}

// NewNSAccessibilityCustomRotorItemSearchDelegate creates an Objective-C object implementing the [NSAccessibilityCustomRotorItemSearchDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSAccessibilityCustomRotorItemSearchDelegateObject] satisfies the [NSAccessibilityCustomRotorItemSearchDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate
func NewNSAccessibilityCustomRotorItemSearchDelegate(config NSAccessibilityCustomRotorItemSearchDelegateConfig) NSAccessibilityCustomRotorItemSearchDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSAccessibilityCustomRotorItemSearchDelegate_%d", n)

	var methods []objc.MethodDef

	if config.RotorResultForSearchParameters != nil {
		fn := config.RotorResultForSearchParameters
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rotor:resultForSearchParameters:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rotorID objc.ID, searchParametersID objc.ID) objc.ID {
				rotor := NSAccessibilityCustomRotorFromID(rotorID)
				searchParameters := NSAccessibilityCustomRotorSearchParametersFromID(searchParametersID)
				return fn(rotor, searchParameters).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSAccessibilityCustomRotorItemSearchDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSAccessibilityCustomRotorItemSearchDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSAccessibilityCustomRotorItemSearchDelegateObjectFromID(instance)
}

