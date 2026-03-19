// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextInputContext] class.
var (
	_NSTextInputContextClass     NSTextInputContextClass
	_NSTextInputContextClassOnce sync.Once
)

func getNSTextInputContextClass() NSTextInputContextClass {
	_NSTextInputContextClassOnce.Do(func() {
		_NSTextInputContextClass = NSTextInputContextClass{class: objc.GetClass("NSTextInputContext")}
	})
	return _NSTextInputContextClass
}

// GetNSTextInputContextClass returns the class object for NSTextInputContext.
func GetNSTextInputContextClass() NSTextInputContextClass {
	return getNSTextInputContextClass()
}

type NSTextInputContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextInputContextClass) Alloc() NSTextInputContext {
	rv := objc.Send[NSTextInputContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the Cocoa text input system.
//
// # Overview
// 
// The text input system communicates primarily with the client of the
// activated input context via the [NSTextInputClient] protocol.
//
// # Creating an Input Context
//
//   - [NSTextInputContext.InitWithClient]: The designated initializer
//
// # Getting the Input Context and Client
//
//   - [NSTextInputContext.Client]: The owner of this input context. (read-only)
//
// # Configuring the Input Context
//
//   - [NSTextInputContext.AcceptsGlyphInfo]: A Boolean value that indicates whether the client handles [NSGlyphInfoAttributeName] or not.
//   - [NSTextInputContext.SetAcceptsGlyphInfo]
//   - [NSTextInputContext.AllowedInputSourceLocales]: The set of keyboard input source locales allowed when this input context is active.
//   - [NSTextInputContext.SetAllowedInputSourceLocales]
//
// # Activating the Input Context
//
//   - [NSTextInputContext.Activate]: Activates the receiver.
//   - [NSTextInputContext.Deactivate]: Deactivates the receiver.
//
// # Handling Input Sources
//
//   - [NSTextInputContext.HandleEvent]: Tells the Cocoa text input system to handle mouse or key events.
//   - [NSTextInputContext.DiscardMarkedText]: Tells the Cocoa text input system to discard the current conversion session.
//   - [NSTextInputContext.InvalidateCharacterCoordinates]: Notifies the Cocoa text input system that the position information previously queried via methods like `` needs to be updated.
//   - [NSTextInputContext.KeyboardInputSources]: The array of keyboard text input source identifier strings available to the receiver. (read-only)
//   - [NSTextInputContext.SelectedKeyboardInputSource]: The identifier string for the selected keyboard text input source.
//   - [NSTextInputContext.SetSelectedKeyboardInputSource]
//
// # Instance Methods
//
//   - [NSTextInputContext.TextInputClientDidEndScrollingOrZooming]
//   - [NSTextInputContext.TextInputClientWillStartScrollingOrZooming]
//   - [NSTextInputContext.TextInputClientDidScroll]
//   - [NSTextInputContext.TextInputClientDidUpdateSelection]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext
type NSTextInputContext struct {
	objectivec.Object
}

// NSTextInputContextFromID constructs a [NSTextInputContext] from an objc.ID.
//
// An object that represents the Cocoa text input system.
func NSTextInputContextFromID(id objc.ID) NSTextInputContext {
	return NSTextInputContext{objectivec.Object{ID: id}}
}
// NOTE: NSTextInputContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextInputContext] class.
//
// # Creating an Input Context
//
//   - [INSTextInputContext.InitWithClient]: The designated initializer
//
// # Getting the Input Context and Client
//
//   - [INSTextInputContext.Client]: The owner of this input context. (read-only)
//
// # Configuring the Input Context
//
//   - [INSTextInputContext.AcceptsGlyphInfo]: A Boolean value that indicates whether the client handles [NSGlyphInfoAttributeName] or not.
//   - [INSTextInputContext.SetAcceptsGlyphInfo]
//   - [INSTextInputContext.AllowedInputSourceLocales]: The set of keyboard input source locales allowed when this input context is active.
//   - [INSTextInputContext.SetAllowedInputSourceLocales]
//
// # Activating the Input Context
//
//   - [INSTextInputContext.Activate]: Activates the receiver.
//   - [INSTextInputContext.Deactivate]: Deactivates the receiver.
//
// # Handling Input Sources
//
//   - [INSTextInputContext.HandleEvent]: Tells the Cocoa text input system to handle mouse or key events.
//   - [INSTextInputContext.DiscardMarkedText]: Tells the Cocoa text input system to discard the current conversion session.
//   - [INSTextInputContext.InvalidateCharacterCoordinates]: Notifies the Cocoa text input system that the position information previously queried via methods like `` needs to be updated.
//   - [INSTextInputContext.KeyboardInputSources]: The array of keyboard text input source identifier strings available to the receiver. (read-only)
//   - [INSTextInputContext.SelectedKeyboardInputSource]: The identifier string for the selected keyboard text input source.
//   - [INSTextInputContext.SetSelectedKeyboardInputSource]
//
// # Instance Methods
//
//   - [INSTextInputContext.TextInputClientDidEndScrollingOrZooming]
//   - [INSTextInputContext.TextInputClientWillStartScrollingOrZooming]
//   - [INSTextInputContext.TextInputClientDidScroll]
//   - [INSTextInputContext.TextInputClientDidUpdateSelection]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext
type INSTextInputContext interface {
	objectivec.IObject

	// Topic: Creating an Input Context

	// The designated initializer
	InitWithClient(client NSTextInputClient) NSTextInputContext

	// Topic: Getting the Input Context and Client

	// The owner of this input context. (read-only)
	Client() NSTextInputClient

	// Topic: Configuring the Input Context

	// A Boolean value that indicates whether the client handles [NSGlyphInfoAttributeName] or not.
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	// The set of keyboard input source locales allowed when this input context is active.
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)

	// Topic: Activating the Input Context

	// Activates the receiver.
	Activate()
	// Deactivates the receiver.
	Deactivate()

	// Topic: Handling Input Sources

	// Tells the Cocoa text input system to handle mouse or key events.
	HandleEvent(event INSEvent) bool
	// Tells the Cocoa text input system to discard the current conversion session.
	DiscardMarkedText()
	// Notifies the Cocoa text input system that the position information previously queried via methods like `` needs to be updated.
	InvalidateCharacterCoordinates()
	// The array of keyboard text input source identifier strings available to the receiver. (read-only)
	KeyboardInputSources() []string
	// The identifier string for the selected keyboard text input source.
	SelectedKeyboardInputSource() NSTextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value NSTextInputSourceIdentifier)

	// Topic: Instance Methods

	TextInputClientDidEndScrollingOrZooming()
	TextInputClientWillStartScrollingOrZooming()
	TextInputClientDidScroll()
	TextInputClientDidUpdateSelection()
}

// Init initializes the instance.
func (t NSTextInputContext) Init() NSTextInputContext {
	rv := objc.Send[NSTextInputContext](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextInputContext) Autorelease() NSTextInputContext {
	rv := objc.Send[NSTextInputContext](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextInputContext creates a new NSTextInputContext instance.
func NewNSTextInputContext() NSTextInputContext {
	class := getNSTextInputContextClass()
	rv := objc.Send[NSTextInputContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The designated initializer
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/init(client:)
func NewTextInputContextWithClient(client NSTextInputClient) NSTextInputContext {
	instance := getNSTextInputContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClient:"), client)
	return NSTextInputContextFromID(rv)
}

// The designated initializer
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/init(client:)
func (t NSTextInputContext) InitWithClient(client NSTextInputClient) NSTextInputContext {
	rv := objc.Send[NSTextInputContext](t.ID, objc.Sel("initWithClient:"), client)
	return rv
}
// Activates the receiver.
//
// # Discussion
// 
// You should not call this method directly; it is invoked by the system. It
// is provided as an override point for subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/activate()
func (t NSTextInputContext) Activate() {
	objc.Send[objc.ID](t.ID, objc.Sel("activate"))
}
// Deactivates the receiver.
//
// # Discussion
// 
// You should not call this method directly; it is invoked by the system. It
// is provided as an override point for subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/deactivate()
func (t NSTextInputContext) Deactivate() {
	objc.Send[objc.ID](t.ID, objc.Sel("deactivate"))
}
// Tells the Cocoa text input system to handle mouse or key events.
//
// event: The event to handle.
//
// # Return Value
// 
// [true] if the system consumed the event; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/handleEvent(_:)
func (t NSTextInputContext) HandleEvent(event INSEvent) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("handleEvent:"), event)
	return rv
}
// Tells the Cocoa text input system to discard the current conversion
// session.
//
// # Discussion
// 
// The client should clear its marked range when sending this message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/discardMarkedText()
func (t NSTextInputContext) DiscardMarkedText() {
	objc.Send[objc.ID](t.ID, objc.Sel("discardMarkedText"))
}
// Notifies the Cocoa text input system that the position information
// previously queried via methods like `` needs to be updated.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/invalidateCharacterCoordinates()
func (t NSTextInputContext) InvalidateCharacterCoordinates() {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateCharacterCoordinates"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientDidEndScrollingOrZooming()
func (t NSTextInputContext) TextInputClientDidEndScrollingOrZooming() {
	objc.Send[objc.ID](t.ID, objc.Sel("textInputClientDidEndScrollingOrZooming"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientWillStartScrollingOrZooming()
func (t NSTextInputContext) TextInputClientWillStartScrollingOrZooming() {
	objc.Send[objc.ID](t.ID, objc.Sel("textInputClientWillStartScrollingOrZooming"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientDidScroll()
func (t NSTextInputContext) TextInputClientDidScroll() {
	objc.Send[objc.ID](t.ID, objc.Sel("textInputClientDidScroll"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientDidUpdateSelection()
func (t NSTextInputContext) TextInputClientDidUpdateSelection() {
	objc.Send[objc.ID](t.ID, objc.Sel("textInputClientDidUpdateSelection"))
}

// Returns the display name for the given text input source identifier.
//
// inputSourceIdentifier: The text input source identifier.
//
// # Return Value
// 
// The localized display name for `inputSourceIdentifier`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/localizedName(forInputSource:)
func (_NSTextInputContextClass NSTextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier NSTextInputSourceIdentifier) string {
	rv := objc.Send[objc.ID](objc.ID(_NSTextInputContextClass.class), objc.Sel("localizedNameForInputSource:"), objc.String(string(inputSourceIdentifier)))
	return foundation.NSStringFromID(rv).String()
}

// The owner of this input context. (read-only)
//
// # Discussion
// 
// The client (owner) of the input context, typically an [NSView] instance,
// retains its [NSTextInputContext] instance. The [NSTextInputContext]
// instance doesn’t retain its client.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/client
func (t NSTextInputContext) Client() NSTextInputClient {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("client"))
	return NSTextInputClientObjectFromID(rv)
}
// A Boolean value that indicates whether the client handles
// [NSGlyphInfoAttributeName] or not.
//
// # Discussion
// 
// The default value is determined by examining the return value from sending
// a `validAttributesForMarkedText` message to the client at initialization.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/acceptsGlyphInfo
func (t NSTextInputContext) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}
func (t NSTextInputContext) SetAcceptsGlyphInfo(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}
// The set of keyboard input source locales allowed when this input context is
// active.
//
// # Discussion
// 
// [NSAllRomanInputSourcesLocaleIdentifier] can be specified as a valid
// locale.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/allowedInputSourceLocales
func (t NSTextInputContext) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("allowedInputSourceLocales"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTextInputContext) SetAllowedInputSourceLocales(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedInputSourceLocales:"), objectivec.StringSliceToNSArray(value))
}
// The array of keyboard text input source identifier strings available to the
// receiver. (read-only)
//
// # Discussion
// 
// The Text Input Source Services API identifies text input sources with text
// input source identifier strings (for example,
// `com.AppleXCUIElementTypeInputmethod().Kotoeri.Japanese`) supplied by the
// underlying text input sources framework. The ID corresponds to the
// `kTISPropertyInputSourceID` attribute.
// 
// For more information on the Text Input Source Services API, see `Text Input
// Source Services`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/keyboardInputSources
func (t NSTextInputContext) KeyboardInputSources() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("keyboardInputSources"))
	return objc.ConvertSliceToStrings(rv)
}
// The identifier string for the selected keyboard text input source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/selectedKeyboardInputSource
func (t NSTextInputContext) SelectedKeyboardInputSource() NSTextInputSourceIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedKeyboardInputSource"))
	return NSTextInputSourceIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTextInputContext) SetSelectedKeyboardInputSource(value NSTextInputSourceIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedKeyboardInputSource:"), objc.String(string(value)))
}

// Returns the current, activated, text input context object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/current
func (_NSTextInputContextClass NSTextInputContextClass) CurrentInputContext() NSTextInputContext {
	rv := objc.Send[objc.ID](objc.ID(_NSTextInputContextClass.class), objc.Sel("currentInputContext"))
	return NSTextInputContextFromID(objc.ID(rv))
}

