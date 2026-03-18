// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ListFormatter] class.
var (
	_ListFormatterClass     ListFormatterClass
	_ListFormatterClassOnce sync.Once
)

func getListFormatterClass() ListFormatterClass {
	_ListFormatterClassOnce.Do(func() {
		_ListFormatterClass = ListFormatterClass{class: objc.GetClass("NSListFormatter")}
	})
	return _ListFormatterClass
}

// GetListFormatterClass returns the class object for NSListFormatter.
func GetListFormatterClass() ListFormatterClass {
	return getListFormatterClass()
}

type ListFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (lc ListFormatterClass) Alloc() ListFormatter {
	rv := objc.Send[ListFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides locale-correct formatting of a list of items using
// the appropriate separator and conjunction.
//
// # Overview
// 
// The list formatter isn’t aware of the context where the formatted string
// will be used and doesn’t provide capitalization customization of the list
// items. The formatted result may not be grammatically correct if placed in a
// sentence, and it should only be used in a standalone manner.
//
// # Converting Arrays to Formatted Lists
//
//   - [ListFormatter.StringFromItems]: Creates a formatted string for an array of items.
//
// # Configuring Formatter Options
//
//   - [ListFormatter.ItemFormatter]: An object that formats each item in the list.
//   - [ListFormatter.SetItemFormatter]
//   - [ListFormatter.Locale]: The locale to use when formatting items in the list.
//   - [ListFormatter.SetLocale]
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter
type ListFormatter struct {
	NSFormatter
}

// ListFormatterFromID constructs a [ListFormatter] from an objc.ID.
//
// An object that provides locale-correct formatting of a list of items using
// the appropriate separator and conjunction.
func ListFormatterFromID(id objc.ID) ListFormatter {
	return NSListFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSListFormatterFromID is an alias for [ListFormatterFromID] for cross-framework compatibility.
func NSListFormatterFromID(id objc.ID) ListFormatter { return ListFormatterFromID(id) }
// NOTE: ListFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ListFormatter] class.
//
// # Converting Arrays to Formatted Lists
//
//   - [IListFormatter.StringFromItems]: Creates a formatted string for an array of items.
//
// # Configuring Formatter Options
//
//   - [IListFormatter.ItemFormatter]: An object that formats each item in the list.
//   - [IListFormatter.SetItemFormatter]
//   - [IListFormatter.Locale]: The locale to use when formatting items in the list.
//   - [IListFormatter.SetLocale]
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter
type IListFormatter interface {
	INSFormatter

	// Topic: Converting Arrays to Formatted Lists

	// Creates a formatted string for an array of items.
	StringFromItems(items INSArray) string

	// Topic: Configuring Formatter Options

	// An object that formats each item in the list.
	ItemFormatter() INSFormatter
	SetItemFormatter(value INSFormatter)
	// The locale to use when formatting items in the list.
	Locale() INSLocale
	SetLocale(value INSLocale)
}

// Init initializes the instance.
func (l ListFormatter) Init() ListFormatter {
	rv := objc.Send[ListFormatter](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l ListFormatter) Autorelease() ListFormatter {
	rv := objc.Send[ListFormatter](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewListFormatter creates a new ListFormatter instance.
func NewListFormatter() ListFormatter {
	class := getListFormatterClass()
	rv := objc.Send[ListFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewListFormatterWithCoder(coder INSCoder) ListFormatter {
	instance := getListFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ListFormatterFromID(rv)
}

// Creates a formatted string for an array of items.
//
// items: An array of objects to format as a list.
//
// # Return Value
// 
// A formatted string representing the list of objects in an array. Returns
// `nil` if the formatter can’t generate a description for all objects in
// the array, or if `obj` is `nil`.
//
// # Discussion
// 
// The list formatter uses [ItemFormatter] to format each item in the array.
// If [ItemFormatter] doesn’t apply to a particular item, the list formatter
// falls back to the item’s [DescriptionWithLocale] or
// [LocalizedDescription] if implemented. If those methods aren’t
// implemented, the formatter uses [description] instead.
//
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter/string(from:)
func (l ListFormatter) StringFromItems(items INSArray) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("stringFromItems:"), items)
	return NSStringFromID(rv).String()
}

// Constructs a formatted string from an array of strings that uses the list
// format specific to the current locale.
//
// strings: An array of strings to join together in a list.
//
// # Return Value
// 
// A formatted string that joins together a list of strings using a
// locale-specific list format.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter/localizedString(byJoining:)
func (_ListFormatterClass ListFormatterClass) LocalizedStringByJoiningStrings(strings []string) string {
	rv := objc.Send[objc.ID](objc.ID(_ListFormatterClass.class), objc.Sel("localizedStringByJoiningStrings:"), objectivec.StringSliceToNSArray(strings))
	return NSStringFromID(rv).String()
}

// An object that formats each item in the list.
//
// # Discussion
// 
// If this property isn’t set, the list formatter falls back to the item’s
// [DescriptionWithLocale] or [LocalizedDescription] methods if implemented.
// If those methods aren’t implemented, the formatter uses [description]
// instead.
//
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter/itemFormatter
func (l ListFormatter) ItemFormatter() INSFormatter {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("itemFormatter"))
	return NSFormatterFromID(objc.ID(rv))
}
func (l ListFormatter) SetItemFormatter(value INSFormatter) {
	objc.Send[struct{}](l.ID, objc.Sel("setItemFormatter:"), value)
}

// The locale to use when formatting items in the list.
//
// # Discussion
// 
// The default value is [AutoupdatingCurrentLocale]. If you set this property
// to `nil`, the formatter resets to using [AutoupdatingCurrentLocale].
//
// See: https://developer.apple.com/documentation/Foundation/ListFormatter/locale
func (l ListFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (l ListFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](l.ID, objc.Sel("setLocale:"), value)
}

