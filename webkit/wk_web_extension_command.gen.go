// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionCommand] class.
var (
	_WKWebExtensionCommandClass     WKWebExtensionCommandClass
	_WKWebExtensionCommandClassOnce sync.Once
)

func getWKWebExtensionCommandClass() WKWebExtensionCommandClass {
	_WKWebExtensionCommandClassOnce.Do(func() {
		_WKWebExtensionCommandClass = WKWebExtensionCommandClass{class: objc.GetClass("WKWebExtensionCommand")}
	})
	return _WKWebExtensionCommandClass
}

// GetWKWebExtensionCommandClass returns the class object for WKWebExtensionCommand.
func GetWKWebExtensionCommandClass() WKWebExtensionCommandClass {
	return getWKWebExtensionCommandClass()
}

type WKWebExtensionCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionCommandClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionCommandClass) Alloc() WKWebExtensionCommand {
	rv := objc.Send[WKWebExtensionCommand](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the properties for an individual web extension
// command.
//
// # Overview
//
// Provides access to command properties such as a unique identifier, a
// descriptive title, and shortcut keys. Commands can be used by a web
// extension to perform specific actions within a web extension context, such
// toggling features, or interacting with web content. These commands enhance
// the functionality of the extension by allowing users to invoke actions
// quickly.
//
// # Instance Properties
//
//   - [WKWebExtensionCommand.ActivationKey]: The primary key used to trigger the command, distinct from any modifier flags.
//   - [WKWebExtensionCommand.SetActivationKey]
//   - [WKWebExtensionCommand.Identifier]: A unique identifier for the command.
//   - [WKWebExtensionCommand.MenuItem]: A menu item representation of the web extension command for use in menus.
//   - [WKWebExtensionCommand.ModifierFlags]: The modifier flags used with the activation key to trigger the command.
//   - [WKWebExtensionCommand.SetModifierFlags]
//   - [WKWebExtensionCommand.Title]: A descriptive title for the command to help discoverability.
//   - [WKWebExtensionCommand.WebExtensionContext]: The web extension context associated with the command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command
type WKWebExtensionCommand struct {
	objectivec.Object
}

// WKWebExtensionCommandFromID constructs a [WKWebExtensionCommand] from an objc.ID.
//
// An object that encapsulates the properties for an individual web extension
// command.
func WKWebExtensionCommandFromID(id objc.ID) WKWebExtensionCommand {
	return WKWebExtensionCommand{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionCommand] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionCommand.ActivationKey]: The primary key used to trigger the command, distinct from any modifier flags.
//   - [IWKWebExtensionCommand.SetActivationKey]
//   - [IWKWebExtensionCommand.Identifier]: A unique identifier for the command.
//   - [IWKWebExtensionCommand.MenuItem]: A menu item representation of the web extension command for use in menus.
//   - [IWKWebExtensionCommand.ModifierFlags]: The modifier flags used with the activation key to trigger the command.
//   - [IWKWebExtensionCommand.SetModifierFlags]
//   - [IWKWebExtensionCommand.Title]: A descriptive title for the command to help discoverability.
//   - [IWKWebExtensionCommand.WebExtensionContext]: The web extension context associated with the command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command
type IWKWebExtensionCommand interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The primary key used to trigger the command, distinct from any modifier flags.
	ActivationKey() string
	SetActivationKey(value string)
	// A unique identifier for the command.
	Identifier() string
	// A menu item representation of the web extension command for use in menus.
	MenuItem() objectivec.IObject
	// The modifier flags used with the activation key to trigger the command.
	ModifierFlags() objectivec.IObject
	SetModifierFlags(value objectivec.IObject)
	// A descriptive title for the command to help discoverability.
	Title() string
	// The web extension context associated with the command.
	WebExtensionContext() IWKWebExtensionContext
}

// Init initializes the instance.
func (w WKWebExtensionCommand) Init() WKWebExtensionCommand {
	rv := objc.Send[WKWebExtensionCommand](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionCommand) Autorelease() WKWebExtensionCommand {
	rv := objc.Send[WKWebExtensionCommand](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionCommand creates a new WKWebExtensionCommand instance.
func NewWKWebExtensionCommand() WKWebExtensionCommand {
	class := getWKWebExtensionCommandClass()
	rv := objc.Send[WKWebExtensionCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The primary key used to trigger the command, distinct from any modifier
// flags.
//
// # Discussion
//
// This property can be customized within the app to avoid conflicts with
// existing shortcuts or to enable user personalization.
//
// It should accurately represent the activation key as used by the app, which
// the extension can use to display the complete shortcut in its interface.
//
// If no shortcut is desired for the command, the property should be set to
// `nil`. This value should be saved and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/activationKey
func (w WKWebExtensionCommand) ActivationKey() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("activationKey"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebExtensionCommand) SetActivationKey(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setActivationKey:"), objc.String(value))
}

// A unique identifier for the command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/id
func (w WKWebExtensionCommand) Identifier() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// A menu item representation of the web extension command for use in menus.
//
// # Discussion
//
// Provides a representation of the web extension command as a menu item to
// display in the app.
//
// Selecting the menu item will perform the command, offering a convenient and
// visual way for users to execute this web extension command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/menuItem
func (w WKWebExtensionCommand) MenuItem() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("menuItem"))
	return objectivec.Object{ID: rv}
}

// The modifier flags used with the activation key to trigger the command.
//
// # Discussion
//
// This property can be customized within the app to avoid conflicts with
// existing shortcuts or to enable user personalization. It should accurately
// represent the modifier keys as used by the app, which the extension can use
// to display the complete shortcut in its interface.
//
// If no modifiers are desired for the command, the property should be set to
// `0`. This value should be saved and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/modifierFlags
func (w WKWebExtensionCommand) ModifierFlags() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("modifierFlags"))
	return objectivec.Object{ID: rv}
}
func (w WKWebExtensionCommand) SetModifierFlags(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setModifierFlags:"), value)
}

// A descriptive title for the command to help discoverability.
//
// # Discussion
//
// This title can be displayed in user interface elements such as keyboard
// shortcuts lists or menu items to help users understand its purpose.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/title
func (w WKWebExtensionCommand) Title() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// The web extension context associated with the command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/webExtensionContext
func (w WKWebExtensionCommand) WebExtensionContext() IWKWebExtensionContext {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtensionContext"))
	return WKWebExtensionContextFromID(objc.ID(rv))
}
