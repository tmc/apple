// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDictionaryController] class.
var (
	_NSDictionaryControllerClass     NSDictionaryControllerClass
	_NSDictionaryControllerClassOnce sync.Once
)

func getNSDictionaryControllerClass() NSDictionaryControllerClass {
	_NSDictionaryControllerClassOnce.Do(func() {
		_NSDictionaryControllerClass = NSDictionaryControllerClass{class: objc.GetClass("NSDictionaryController")}
	})
	return _NSDictionaryControllerClass
}

// GetNSDictionaryControllerClass returns the class object for NSDictionaryController.
func GetNSDictionaryControllerClass() NSDictionaryControllerClass {
	return getNSDictionaryControllerClass()
}

type NSDictionaryControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDictionaryControllerClass) Alloc() NSDictionaryController {
	rv := objc.Send[NSDictionaryController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bindings-compatible controller that manages the display and editing of a
// dictionary of key-value pairs.
//
// # Overview
// 
// [NSDictionaryController] transforms the contents of a dictionary into an
// array of key-value pairs that can be bound to user interface items such as
// the columns of an [NSTableView].
// 
// The content of an [NSDictionaryController] instance is specified using the
// inherited method [NSDictionaryController.Content] or by binding an [NSDictionary] instance to the
// [contentDictionary] binding. New key/value pairs inserted into the
// dictionary are created using the [NSDictionaryController.NewObject] method. The initial key name
// is set to the string returned by [NSDictionaryController.InitialKey] . The initial key name is
// copied to the newly inserted object, while the object returned by
// [NSDictionaryController.InitialValue] is simply retained. As new items are inserted the controller
// enumerates the initial key name, resulting in key names such as “key”,
// “key1”, “key2”, and so on. This behavior can be customized by
// overriding [NSDictionaryController.NewObject].
// 
// An [NSDictionaryController] instance can be configured to exclude specified
// keys in a dictionary from being returned by [NSDictionaryController.ArrangedObjects] using the
// [NSDictionaryController.ExcludedKeys] property. Similarly, you can specify an array of key names
// that are always included in the arranged objects, even if they are not
// present in the content dictionary, using the [NSDictionaryController.IncludedKeys] property.
// 
// [NSDictionaryController] supports providing localized key names for the
// keys in the dictionary, allowing a user-friendly representation of the key
// name to be displayed. The localized key names are specified by a dictionary
// (using [NSDictionaryController.LocalizedKeyDictionary]) or by providing a strings table (using
// [NSDictionaryController.LocalizedKeyDictionary]).
// 
// The [NSDictionaryController.ArrangedObjects] method returns an array of objects that implement the
// [NSDictionaryControllerKeyValuePair] informal protocol. User interface
// controls are bound to the arranged objects array using key paths such as:
// `arrangedObjects.Key()` (displays the key name), `arrangedObjects.Value()`
// (displays the value for the key), or `arrangedObjects.LocalizedKey()`
// (displays the localized key name). See [NSDictionaryControllerKeyValuePair]
// for more information.
// 
// [NSDictionaryController] overrides [NSDictionaryController.ArrangedObjects] to return an array of
// objects that implement the [NSDictionaryControllerKeyValuePair] informal
// protocol. See [NSDictionaryControllerKeyValuePair] and [Cocoa Bindings
// Programming Topics] for more information.
// 
// The constants listed below are used to specify a binding to
// [bind(_:to:withKeyPath:options:)], [infoForBinding(_:)], [unbind(_:)], and
// [valueClassForBinding(_:)]. See the [Cocoa Bindings Reference] for more
// information.
// 
// - [contentDictionary] - [includedKeys] - [excludedKeys] -
// [localizedKeyDictionary] - [initialKey] - [initialValue]
//
// [Cocoa Bindings Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaBindings/CocoaBindings.html#//apple_ref/doc/uid/10000167i
// [Cocoa Bindings Reference]: https://developer.apple.com/library/archive/documentation/Cocoa/Reference/CocoaBindingsRef/CocoaBindingsRef.html#//apple_ref/doc/uid/10000189i
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [bind(_:to:withKeyPath:options:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/bind(_:to:withKeyPath:options:)
// [contentDictionary]: https://developer.apple.com/documentation/AppKit/NSBindingName/contentDictionary
// [excludedKeys]: https://developer.apple.com/documentation/AppKit/NSBindingName/excludedKeys
// [includedKeys]: https://developer.apple.com/documentation/AppKit/NSBindingName/includedKeys
// [infoForBinding(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/infoForBinding(_:)
// [initialKey]: https://developer.apple.com/documentation/AppKit/NSBindingName/initialKey
// [initialValue]: https://developer.apple.com/documentation/AppKit/NSBindingName/initialValue
// [localizedKeyDictionary]: https://developer.apple.com/documentation/AppKit/NSBindingName/localizedKeyDictionary
// [unbind(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/unbind(_:)
// [valueClassForBinding(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/valueClassForBinding(_:)
//
// # Localizing Key Names
//
//   - [NSDictionaryController.LocalizedKeyDictionary]: The localized key names that are displayed by the receiver in place of the key names.
//   - [NSDictionaryController.SetLocalizedKeyDictionary]
//   - [NSDictionaryController.LocalizedKeyTable]: the strings file used to localize key names.
//   - [NSDictionaryController.SetLocalizedKeyTable]
//
// # Keys to Display
//
//   - [NSDictionaryController.IncludedKeys]: The key names that are represented by a key-value pair, even if they are not present in the receiver’s content dictionary.
//   - [NSDictionaryController.SetIncludedKeys]
//   - [NSDictionaryController.ExcludedKeys]: The key names that are never displayed in the user interface items bound to the receiver.
//   - [NSDictionaryController.SetExcludedKeys]
//
// # Setting Initial Key and Values
//
//   - [NSDictionaryController.InitialKey]: The string used as the initial key name for a newly inserted item.
//   - [NSDictionaryController.SetInitialKey]
//   - [NSDictionaryController.InitialValue]: The string used as the initial value for a newly inserted item.
//   - [NSDictionaryController.SetInitialValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController
type NSDictionaryController struct {
	NSArrayController
}

// NSDictionaryControllerFromID constructs a [NSDictionaryController] from an objc.ID.
//
// A bindings-compatible controller that manages the display and editing of a
// dictionary of key-value pairs.
func NSDictionaryControllerFromID(id objc.ID) NSDictionaryController {
	return NSDictionaryController{NSArrayController: NSArrayControllerFromID(id)}
}
// NOTE: NSDictionaryController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDictionaryController] class.
//
// # Localizing Key Names
//
//   - [INSDictionaryController.LocalizedKeyDictionary]: The localized key names that are displayed by the receiver in place of the key names.
//   - [INSDictionaryController.SetLocalizedKeyDictionary]
//   - [INSDictionaryController.LocalizedKeyTable]: the strings file used to localize key names.
//   - [INSDictionaryController.SetLocalizedKeyTable]
//
// # Keys to Display
//
//   - [INSDictionaryController.IncludedKeys]: The key names that are represented by a key-value pair, even if they are not present in the receiver’s content dictionary.
//   - [INSDictionaryController.SetIncludedKeys]
//   - [INSDictionaryController.ExcludedKeys]: The key names that are never displayed in the user interface items bound to the receiver.
//   - [INSDictionaryController.SetExcludedKeys]
//
// # Setting Initial Key and Values
//
//   - [INSDictionaryController.InitialKey]: The string used as the initial key name for a newly inserted item.
//   - [INSDictionaryController.SetInitialKey]
//   - [INSDictionaryController.InitialValue]: The string used as the initial value for a newly inserted item.
//   - [INSDictionaryController.SetInitialValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController
type INSDictionaryController interface {
	INSArrayController

	// Topic: Localizing Key Names

	// The localized key names that are displayed by the receiver in place of the key names.
	LocalizedKeyDictionary() foundation.INSDictionary
	SetLocalizedKeyDictionary(value foundation.INSDictionary)
	// the strings file used to localize key names.
	LocalizedKeyTable() string
	SetLocalizedKeyTable(value string)

	// Topic: Keys to Display

	// The key names that are represented by a key-value pair, even if they are not present in the receiver’s content dictionary.
	IncludedKeys() []string
	SetIncludedKeys(value []string)
	// The key names that are never displayed in the user interface items bound to the receiver.
	ExcludedKeys() []string
	SetExcludedKeys(value []string)

	// Topic: Setting Initial Key and Values

	// The string used as the initial key name for a newly inserted item.
	InitialKey() string
	SetInitialKey(value string)
	// The string used as the initial value for a newly inserted item.
	InitialValue() objectivec.IObject
	SetInitialValue(value objectivec.IObject)

	// A constant that identifies a content dictionary binding.
	ContentDictionary() NSBindingName
}

// Init initializes the instance.
func (d NSDictionaryController) Init() NSDictionaryController {
	rv := objc.Send[NSDictionaryController](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDictionaryController) Autorelease() NSDictionaryController {
	rv := objc.Send[NSDictionaryController](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDictionaryController creates a new NSDictionaryController instance.
func NewNSDictionaryController() NSDictionaryController {
	class := getNSDictionaryControllerClass()
	rv := objc.Send[NSDictionaryController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(coder:)
func NewDictionaryControllerWithCoder(coder foundation.INSCoder) NSDictionaryController {
	instance := getNSDictionaryControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDictionaryControllerFromID(rv)
}

// Initializes and returns an [NSObjectController] object with the given
// content.
//
// content: The content for the receiver.
//
// # Return Value
// 
// The initialized object controller, with its content object set to
// `content`.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(content:)
func NewDictionaryControllerWithContent(content objectivec.IObject) NSDictionaryController {
	instance := getNSDictionaryControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), content)
	return NSDictionaryControllerFromID(rv)
}

// The localized key names that are displayed by the receiver in place of the
// key names.
//
// # Discussion
// 
// The dictionary contains the key names as the keys, and the localized key
// names as the corresponding values.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/localizedKeyDictionary
func (d NSDictionaryController) LocalizedKeyDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("localizedKeyDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d NSDictionaryController) SetLocalizedKeyDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocalizedKeyDictionary:"), value)
}
// the strings file used to localize key names.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/localizedKeyTable
func (d NSDictionaryController) LocalizedKeyTable() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("localizedKeyTable"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDictionaryController) SetLocalizedKeyTable(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocalizedKeyTable:"), objc.String(value))
}
// The key names that are represented by a key-value pair, even if they are
// not present in the receiver’s content dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/includedKeys
func (d NSDictionaryController) IncludedKeys() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("includedKeys"))
	return objc.ConvertSliceToStrings(rv)
}
func (d NSDictionaryController) SetIncludedKeys(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setIncludedKeys:"), objectivec.StringSliceToNSArray(value))
}
// The key names that are never displayed in the user interface items bound to
// the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/excludedKeys
func (d NSDictionaryController) ExcludedKeys() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("excludedKeys"))
	return objc.ConvertSliceToStrings(rv)
}
func (d NSDictionaryController) SetExcludedKeys(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setExcludedKeys:"), objectivec.StringSliceToNSArray(value))
}
// The string used as the initial key name for a newly inserted item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/initialKey
func (d NSDictionaryController) InitialKey() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initialKey"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDictionaryController) SetInitialKey(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setInitialKey:"), objc.String(value))
}
// The string used as the initial value for a newly inserted item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryController/initialValue
func (d NSDictionaryController) InitialValue() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initialValue"))
	return objectivec.Object{ID: rv}
}
func (d NSDictionaryController) SetInitialValue(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setInitialValue:"), value)
}
// A constant that identifies a content dictionary binding.
//
// See: https://developer.apple.com/documentation/appkit/nsbindingname/contentdictionary
func (d NSDictionaryController) ContentDictionary() NSBindingName {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NSContentDictionaryBinding"))
	return NSBindingName(foundation.NSStringFromID(rv).String())
}

