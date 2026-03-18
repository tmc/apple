// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The optional methods that delegates of content storage objects implement to handle content processing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorageDelegate
type NSTextContentStorageDelegate interface {
	objectivec.IObject
	NSTextContentManagerDelegate
}

// NSTextContentStorageDelegateObject wraps an existing Objective-C object that conforms to the NSTextContentStorageDelegate protocol.
type NSTextContentStorageDelegateObject struct {
	objectivec.Object
}
func (o NSTextContentStorageDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextContentStorageDelegateObjectFromID constructs a [NSTextContentStorageDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextContentStorageDelegateObjectFromID(id objc.ID) NSTextContentStorageDelegateObject {
	return NSTextContentStorageDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a custom paragraph for a range that you provide from the object’s
// attributed string.
//
// textContentStorage: The object’s content manager.
//
// range: The [NSRange] that describes the extent of the string.
// //
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// # Return Value
// 
// A new [NSTextParagraph], or `nil`.
//
// # Discussion
// 
// When non-`nil`, `textContentStorage` uses the text paragraph instead of
// creating the standard [NSTextParagraph] with the attributed substring in
// range. The attributed string for a custom text paragraph must have a length
// of `range.Length()`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorageDelegate/textContentStorage(_:textParagraphWith:)

func (o NSTextContentStorageDelegateObject) TextContentStorageTextParagraphWithRange(textContentStorage INSTextContentStorage, range_ foundation.NSRange) INSTextParagraph {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textContentStorage:textParagraphWithRange:"), textContentStorage, range_)
	return NSTextParagraphFromID(rv)
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

func (o NSTextContentStorageDelegateObject) TextContentManagerTextElementAtLocation(textContentManager INSTextContentManager, location NSTextLocation) INSTextElement {
	
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

func (o NSTextContentStorageDelegateObject) TextContentManagerShouldEnumerateTextElementOptions(textContentManager INSTextContentManager, textElement INSTextElement, options NSTextContentManagerEnumerationOptions) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textContentManager:shouldEnumerateTextElement:options:"), textContentManager, textElement, options)
	return rv
	}

// NSTextContentStorageDelegateConfig holds optional typed callbacks for [NSTextContentStorageDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextcontentstoragedelegate
type NSTextContentStorageDelegateConfig struct {

	// Other Methods
	// TextContentStorageTextParagraphWithRange — Returns a custom paragraph for a range that you provide from the object’s attributed string.
	TextContentStorageTextParagraphWithRange func(textContentStorage NSTextContentStorage, range_ foundation.NSRange) NSTextParagraph
}

// NewNSTextContentStorageDelegate creates an Objective-C object implementing the [NSTextContentStorageDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextContentStorageDelegateObject] satisfies the [NSTextContentStorageDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextcontentstoragedelegate
func NewNSTextContentStorageDelegate(config NSTextContentStorageDelegateConfig) NSTextContentStorageDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextContentStorageDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TextContentStorageTextParagraphWithRange != nil {
		fn := config.TextContentStorageTextParagraphWithRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textContentStorage:textParagraphWithRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textContentStorageID objc.ID, range_ foundation.NSRange) objc.ID {
				textContentStorage := NSTextContentStorageFromID(textContentStorageID)
				return fn(textContentStorage, range_).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextContentStorageDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextContentStorageDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextContentStorageDelegateObjectFromID(instance)
}

