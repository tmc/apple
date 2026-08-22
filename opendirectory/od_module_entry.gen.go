// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODModuleEntry] class.
var (
	_ODModuleEntryClass     ODModuleEntryClass
	_ODModuleEntryClassOnce sync.Once
)

func getODModuleEntryClass() ODModuleEntryClass {
	_ODModuleEntryClassOnce.Do(func() {
		_ODModuleEntryClass = ODModuleEntryClass{class: objc.GetClass("ODModuleEntry")}
	})
	return _ODModuleEntryClass
}

// GetODModuleEntryClass returns the class object for ODModuleEntry.
func GetODModuleEntryClass() ODModuleEntryClass {
	return getODModuleEntryClass()
}

type ODModuleEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODModuleEntryClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODModuleEntryClass) Alloc() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [ODModuleEntry.Mappings]
//   - [ODModuleEntry.SetMappings]
//   - [ODModuleEntry.Name]
//   - [ODModuleEntry.SetName]
//   - [ODModuleEntry.SupportedOptions]
//   - [ODModuleEntry.UuidString]
//   - [ODModuleEntry.SetUuidString]
//   - [ODModuleEntry.XpcServiceName]
//   - [ODModuleEntry.SetXpcServiceName]
//
// # Instance Methods
//
//   - [ODModuleEntry.Option]
//   - [ODModuleEntry.SetOptionValue]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry
type ODModuleEntry struct {
	objectivec.Object
}

// ODModuleEntryFromID constructs a [ODModuleEntry] from an objc.ID.
func ODModuleEntryFromID(id objc.ID) ODModuleEntry {
	return ODModuleEntry{objectivec.Object{ID: id}}
}

// NOTE: ODModuleEntry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODModuleEntry] class.
//
// # Instance Properties
//
//   - [IODModuleEntry.Mappings]
//   - [IODModuleEntry.SetMappings]
//   - [IODModuleEntry.Name]
//   - [IODModuleEntry.SetName]
//   - [IODModuleEntry.SupportedOptions]
//   - [IODModuleEntry.UuidString]
//   - [IODModuleEntry.SetUuidString]
//   - [IODModuleEntry.XpcServiceName]
//   - [IODModuleEntry.SetXpcServiceName]
//
// # Instance Methods
//
//   - [IODModuleEntry.Option]
//   - [IODModuleEntry.SetOptionValue]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry
type IODModuleEntry interface {
	objectivec.IObject

	// Topic: Instance Properties

	Mappings() IODMappings
	SetMappings(value IODMappings)
	Name() string
	SetName(value string)
	SupportedOptions() foundation.INSArray
	UuidString() string
	SetUuidString(value string)
	XpcServiceName() string
	SetXpcServiceName(value string)

	// Topic: Instance Methods

	Option(optionName string) objectivec.IObject
	SetOptionValue(optionName string, value objectivec.IObject)
}

// Init initializes the instance.
func (o ODModuleEntry) Init() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODModuleEntry) Autorelease() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODModuleEntry creates a new ODModuleEntry instance.
func NewODModuleEntry() ODModuleEntry {
	class := getODModuleEntryClass()
	rv := objc.Send[ODModuleEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func NewODModuleEntryWithNameXpcServiceName(name string, xpcServiceName string) ODModuleEntry {
	rv := objc.Send[objc.ID](objc.ID(getODModuleEntryClass().class), objc.Sel("moduleEntryWithName:xpcServiceName:"), objc.String(name), objc.String(xpcServiceName))
	return ODModuleEntryFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/option(_:)
func (o ODModuleEntry) Option(optionName string) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("option:"), objc.String(optionName))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/setOption(_:value:)
func (o ODModuleEntry) SetOptionValue(optionName string, value objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("setOption:value:"), objc.String(optionName), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o ODModuleEntry) Mappings() IODMappings {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mappings"))
	return ODMappingsFromID(objc.ID(rv))
}
func (o ODModuleEntry) SetMappings(value IODMappings) {
	objc.Send[struct{}](o.ID, objc.Sel("setMappings:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o ODModuleEntry) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODModuleEntry) SetName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-swift.property
func (o ODModuleEntry) SupportedOptions() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("supportedOptions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o ODModuleEntry) UuidString() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuidString"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODModuleEntry) SetUuidString(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setUuidString:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o ODModuleEntry) XpcServiceName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("xpcServiceName"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODModuleEntry) SetXpcServiceName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setXpcServiceName:"), objc.String(value))
}
