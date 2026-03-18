// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods implemented by the delegate of an [NSText](<doc://com.apple.appkit/documentation/AppKit/NSText>) object to edit text and change text formats.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate
type NSTextDelegate interface {
	objectivec.IObject
}

// NSTextDelegateObject wraps an existing Objective-C object that conforms to the NSTextDelegate protocol.
type NSTextDelegateObject struct {
	objectivec.Object
}
func (o NSTextDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextDelegateObjectFromID constructs a [NSTextDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextDelegateObjectFromID(id objc.ID) NSTextDelegateObject {
	return NSTextDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate that the text object has changed its characters or
// formatting attributes.
//
// # Discussion
// 
// The name of `aNotification` is [didChangeNotification].
//
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSText/didChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidChange(_:)

func (o NSTextDelegateObject) TextDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidChange:"), notification)
	}

// Invoked when a text object begins to change its text, this method requests
// permission for `aTextObject` to begin editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to make changes.
// If the delegate returns [false], the text object abandons the editing
// operation. This method is also invoked when the user drags and drops a file
// onto the text object.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldBeginEditing(_:)

func (o NSTextDelegateObject) TextShouldBeginEditing(textObject INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
	}

// Informs the delegate that the text object has begun editing (that the user
// has begun changing it).
//
// # Discussion
// 
// The name of `aNotification` is [didBeginEditingNotification].
//
// [didBeginEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didBeginEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidBeginEditing(_:)

func (o NSTextDelegateObject) TextDidBeginEditing(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidBeginEditing:"), notification)
	}

// Invoked from a text object’s implementation of [ResignFirstResponder],
// this method requests permission for `aTextObject` to end editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to finish editing
// and resign first responder status. If the delegate returns [false], the
// text object selects all of its text and remains the first responder.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldEndEditing(_:)

func (o NSTextDelegateObject) TextShouldEndEditing(textObject INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
	}

// Informs the delegate that the text object has finished editing (that it has
// resigned first responder status).
//
// # Discussion
// 
// The name of `aNotification` is [didEndEditingNotification].
//
// [didEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didEndEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidEndEditing(_:)

func (o NSTextDelegateObject) TextDidEndEditing(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidEndEditing:"), notification)
	}

// NSTextDelegateConfig holds optional typed callbacks for [NSTextDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextdelegate
type NSTextDelegateConfig struct {

	// Changing text formatting
	// TextDidChange — Informs the delegate that the text object has changed its characters or formatting attributes.
	TextDidChange func(notification foundation.NSNotification)

	// Editing text
	// TextShouldBeginEditing — Invoked when a text object begins to change its text, this method requests permission for `aTextObject` to begin editing.
	TextShouldBeginEditing func(textObject NSText) bool
	// TextDidBeginEditing — Informs the delegate that the text object has begun editing (that the user has begun changing it).
	TextDidBeginEditing func(notification foundation.NSNotification)
	// TextShouldEndEditing — Invoked from a text object’s implementation of [resignFirstResponder()](<doc://com.apple.appkit/documentation/AppKit/NSResponder/resignFirstResponder()>), this method requests permission for `aTextObject` to end editing.
	TextShouldEndEditing func(textObject NSText) bool
	// TextDidEndEditing — Informs the delegate that the text object has finished editing (that it has resigned first responder status).
	TextDidEndEditing func(notification foundation.NSNotification)
}

// NewNSTextDelegate creates an Objective-C object implementing the [NSTextDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextDelegateObject] satisfies the [NSTextDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextdelegate
func NewNSTextDelegate(config NSTextDelegateConfig) NSTextDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TextDidChange != nil {
		fn := config.TextDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.TextShouldBeginEditing != nil {
		fn := config.TextShouldBeginEditing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textShouldBeginEditing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textObjectID objc.ID) bool {
				textObject := NSTextFromID(textObjectID)
				return fn(textObject)
			},
		})
	}

	if config.TextDidBeginEditing != nil {
		fn := config.TextDidBeginEditing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textDidBeginEditing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.TextShouldEndEditing != nil {
		fn := config.TextShouldEndEditing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textShouldEndEditing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textObjectID objc.ID) bool {
				textObject := NSTextFromID(textObjectID)
				return fn(textObject)
			},
		})
	}

	if config.TextDidEndEditing != nil {
		fn := config.TextDidEndEditing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textDidEndEditing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextDelegateObjectFromID(instance)
}

