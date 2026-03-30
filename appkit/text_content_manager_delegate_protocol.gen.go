// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The optional methods that delegates of content manager objects implement for customizing or validating text elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManagerDelegate
type NSTextContentManagerDelegate interface {
	objectivec.IObject
}

// NSTextContentManagerDelegateObject wraps an existing Objective-C object that conforms to the NSTextContentManagerDelegate protocol.
type NSTextContentManagerDelegateObject struct {
	objectivec.Object
}

func (o NSTextContentManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextContentManagerDelegateObjectFromID constructs a [NSTextContentManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextContentManagerDelegateObjectFromID(id objc.ID) NSTextContentManagerDelegateObject {
	return NSTextContentManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The method the framework calls to return the text element at a specific
// location.
//
// textContentManager: The content manager.
//
// location: The location of the element.
//
// # Return Value
//
// An [NSTextElement].
//
// # Discussion
//
// When non-`nil`, `textContentManager` uses the text element you specify
// instead of creating one based on its standard mapping logic.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManagerDelegate/textContentManager(_:textElementAt:)
func (o NSTextContentManagerDelegateObject) TextContentManagerTextElementAtLocation(textContentManager INSTextContentManager, location NSTextLocation) INSTextElement {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textContentManager:textElementAtLocation:"), textContentManager, location)
	return NSTextElementFromID(rv)
}

// Returns a Boolean value that indicates whether the framework should skip
// this text element in the enumeration.
//
// textContentManager: The content manager.
//
// textElement: The [NSTextElement] to evaluate.
//
// options: One of the available [NSTextElementProviderEnumerationOptions] options.
//
// # Return Value
//
// A Boolean value that informs the framework to skip this `textElement` in
// the enumeration. Returning `false` indicates `textElement` to be skipped;
// otherwise the element is included in the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManagerDelegate/textContentManager(_:shouldEnumerate:options:)
func (o NSTextContentManagerDelegateObject) TextContentManagerShouldEnumerateTextElementOptions(textContentManager INSTextContentManager, textElement INSTextElement, options NSTextContentManagerEnumerationOptions) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("textContentManager:shouldEnumerateTextElement:options:"), textContentManager, textElement, options)
	return rv
}

// NSTextContentManagerDelegateConfig holds optional typed callbacks for [NSTextContentManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextcontentmanagerdelegate
type NSTextContentManagerDelegateConfig struct {

	// Other Methods
	// TextContentManagerShouldEnumerateTextElementOptions — Returns a Boolean value that indicates whether the framework should skip this text element in the enumeration.
	TextContentManagerShouldEnumerateTextElementOptions func(textContentManager NSTextContentManager, textElement NSTextElement, options NSTextContentManagerEnumerationOptions) bool
}

// NewNSTextContentManagerDelegate creates an Objective-C object implementing the [NSTextContentManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextContentManagerDelegateObject] satisfies the [NSTextContentManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextcontentmanagerdelegate
func NewNSTextContentManagerDelegate(config NSTextContentManagerDelegateConfig) NSTextContentManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextContentManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TextContentManagerShouldEnumerateTextElementOptions != nil {
		fn := config.TextContentManagerShouldEnumerateTextElementOptions
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textContentManager:shouldEnumerateTextElement:options:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textContentManagerID objc.ID, textElementID objc.ID, options NSTextContentManagerEnumerationOptions) bool {
				textContentManager := NSTextContentManagerFromID(textContentManagerID)
				textElement := NSTextElementFromID(textElementID)
				return fn(textContentManager, textElement, options)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextContentManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextContentManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextContentManagerDelegateObjectFromID(instance)
}
