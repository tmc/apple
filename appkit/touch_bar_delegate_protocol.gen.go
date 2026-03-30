// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that allows you to provide the items for a bar dynamically.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarDelegate
type NSTouchBarDelegate interface {
	objectivec.IObject
}

// NSTouchBarDelegateObject wraps an existing Objective-C object that conforms to the NSTouchBarDelegate protocol.
type NSTouchBarDelegateObject struct {
	objectivec.Object
}

func (o NSTouchBarDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTouchBarDelegateObjectFromID constructs a [NSTouchBarDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTouchBarDelegateObjectFromID(id objc.ID) NSTouchBarDelegateObject {
	return NSTouchBarDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate object for the bar item for the specified bar and item
// identifier.
//
// touchBar: The bar that’s requesting the bar item.
//
// identifier: The item identifier associated with the item being requested.
//
// # Return Value
//
// A fully initialized bar item for the specified bar and identifier.
//
// # Discussion
//
// When the system needs to populate a bar’s items array, the system calls
// this delegate method to retrieve an item if that item can’t be found in
// the bar’s private array or in the bar’s [TemplateItems] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarDelegate/touchBar(_:makeItemForIdentifier:)
func (o NSTouchBarDelegateObject) TouchBarMakeItemForIdentifier(touchBar INSTouchBar, identifier NSTouchBarItemIdentifier) INSTouchBarItem {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("touchBar:makeItemForIdentifier:"), touchBar, objc.String(string(identifier)))
	return NSTouchBarItemFromID(rv)
}

// NSTouchBarDelegateConfig holds optional typed callbacks for [NSTouchBarDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstouchbardelegate
type NSTouchBarDelegateConfig struct {

	// Providing bar items
	// TouchBarMakeItemForIdentifier — Asks the delegate object for the bar item for the specified bar and item identifier.
	TouchBarMakeItemForIdentifier func(touchBar NSTouchBar, identifier NSTouchBarItemIdentifier) NSTouchBarItem
}

// NewNSTouchBarDelegate creates an Objective-C object implementing the [NSTouchBarDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTouchBarDelegateObject] satisfies the [NSTouchBarDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstouchbardelegate
func NewNSTouchBarDelegate(config NSTouchBarDelegateConfig) NSTouchBarDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTouchBarDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TouchBarMakeItemForIdentifier != nil {
		fn := config.TouchBarMakeItemForIdentifier
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("touchBar:makeItemForIdentifier:"),
			Fn: func(self objc.ID, _cmd objc.SEL, touchBarID objc.ID, identifierID objc.ID) objc.ID {
				touchBar := NSTouchBarFromID(touchBarID)
				identifier := NSTouchBarItemIdentifier(objc.GoString(objc.Send[*byte](identifierID, objc.Sel("UTF8String"))))
				return fn(touchBar, identifier).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTouchBarDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTouchBarDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTouchBarDelegateObjectFromID(instance)
}
