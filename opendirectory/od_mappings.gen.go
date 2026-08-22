// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODMappings] class.
var (
	_ODMappingsClass     ODMappingsClass
	_ODMappingsClassOnce sync.Once
)

func getODMappingsClass() ODMappingsClass {
	_ODMappingsClassOnce.Do(func() {
		_ODMappingsClass = ODMappingsClass{class: objc.GetClass("ODMappings")}
	})
	return _ODMappingsClass
}

// GetODMappingsClass returns the class object for ODMappings.
func GetODMappingsClass() ODMappingsClass {
	return getODMappingsClass()
}

type ODMappingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODMappingsClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODMappingsClass) Alloc() ODMappings {
	rv := objc.Send[ODMappings](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [ODMappings.Comment]
//   - [ODMappings.SetComment]
//   - [ODMappings.Function]
//   - [ODMappings.SetFunction]
//   - [ODMappings.FunctionAttributes]
//   - [ODMappings.SetFunctionAttributes]
//   - [ODMappings.Identifier]
//   - [ODMappings.SetIdentifier]
//   - [ODMappings.RecordTypes]
//   - [ODMappings.TemplateName]
//   - [ODMappings.SetTemplateName]
//
// # Instance Methods
//
//   - [ODMappings.RecordMapForStandardRecordType]
//   - [ODMappings.SetRecordMapForStandardRecordType]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings
type ODMappings struct {
	objectivec.Object
}

// ODMappingsFromID constructs a [ODMappings] from an objc.ID.
func ODMappingsFromID(id objc.ID) ODMappings {
	return ODMappings{objectivec.Object{ID: id}}
}

// NOTE: ODMappings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODMappings] class.
//
// # Instance Properties
//
//   - [IODMappings.Comment]
//   - [IODMappings.SetComment]
//   - [IODMappings.Function]
//   - [IODMappings.SetFunction]
//   - [IODMappings.FunctionAttributes]
//   - [IODMappings.SetFunctionAttributes]
//   - [IODMappings.Identifier]
//   - [IODMappings.SetIdentifier]
//   - [IODMappings.RecordTypes]
//   - [IODMappings.TemplateName]
//   - [IODMappings.SetTemplateName]
//
// # Instance Methods
//
//   - [IODMappings.RecordMapForStandardRecordType]
//   - [IODMappings.SetRecordMapForStandardRecordType]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings
type IODMappings interface {
	objectivec.IObject

	// Topic: Instance Properties

	Comment() string
	SetComment(value string)
	Function() string
	SetFunction(value string)
	FunctionAttributes() foundation.INSArray
	SetFunctionAttributes(value foundation.INSArray)
	Identifier() string
	SetIdentifier(value string)
	RecordTypes() foundation.INSArray
	TemplateName() string
	SetTemplateName(value string)

	// Topic: Instance Methods

	RecordMapForStandardRecordType(stdType string) IODRecordMap
	SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType string)
}

// Init initializes the instance.
func (o ODMappings) Init() ODMappings {
	rv := objc.Send[ODMappings](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODMappings) Autorelease() ODMappings {
	rv := objc.Send[ODMappings](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODMappings creates a new ODMappings instance.
func NewODMappings() ODMappings {
	class := getODMappingsClass()
	rv := objc.Send[ODMappings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordMap(forStandardRecordType:)
func (o ODMappings) RecordMapForStandardRecordType(stdType string) IODRecordMap {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordMapForStandardRecordType:"), objc.String(stdType))
	return ODRecordMapFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/setRecordMap(_:forStandardRecordType:)
func (o ODMappings) SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType string) {
	objc.Send[objc.ID](o.ID, objc.Sel("setRecordMap:forStandardRecordType:"), map_, objc.String(stdType))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o ODMappings) Comment() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comment"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODMappings) SetComment(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setComment:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o ODMappings) Function() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("function"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODMappings) SetFunction(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setFunction:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o ODMappings) FunctionAttributes() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionAttributes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODMappings) SetFunctionAttributes(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setFunctionAttributes:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o ODMappings) Identifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODMappings) SetIdentifier(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-swift.property
func (o ODMappings) RecordTypes() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordTypes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o ODMappings) TemplateName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("templateName"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODMappings) SetTemplateName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setTemplateName:"), objc.String(value))
}
