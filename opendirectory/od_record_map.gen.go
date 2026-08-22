// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODRecordMap] class.
var (
	_ODRecordMapClass     ODRecordMapClass
	_ODRecordMapClassOnce sync.Once
)

func getODRecordMapClass() ODRecordMapClass {
	_ODRecordMapClassOnce.Do(func() {
		_ODRecordMapClass = ODRecordMapClass{class: objc.GetClass("ODRecordMap")}
	})
	return _ODRecordMapClass
}

// GetODRecordMapClass returns the class object for ODRecordMap.
func GetODRecordMapClass() ODRecordMapClass {
	return getODRecordMapClass()
}

type ODRecordMapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODRecordMapClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODRecordMapClass) Alloc() ODRecordMap {
	rv := objc.Send[ODRecordMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [ODRecordMap.Attributes]
//   - [ODRecordMap.Native]
//   - [ODRecordMap.SetNative]
//   - [ODRecordMap.OdPredicate]
//   - [ODRecordMap.SetOdPredicate]
//   - [ODRecordMap.StandardAttributeTypes]
//
// # Instance Methods
//
//   - [ODRecordMap.AttributeMapForStandardAttribute]
//   - [ODRecordMap.SetAttributeMapForStandardAttribute]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap
type ODRecordMap struct {
	objectivec.Object
}

// ODRecordMapFromID constructs a [ODRecordMap] from an objc.ID.
func ODRecordMapFromID(id objc.ID) ODRecordMap {
	return ODRecordMap{objectivec.Object{ID: id}}
}

// NOTE: ODRecordMap adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODRecordMap] class.
//
// # Instance Properties
//
//   - [IODRecordMap.Attributes]
//   - [IODRecordMap.Native]
//   - [IODRecordMap.SetNative]
//   - [IODRecordMap.OdPredicate]
//   - [IODRecordMap.SetOdPredicate]
//   - [IODRecordMap.StandardAttributeTypes]
//
// # Instance Methods
//
//   - [IODRecordMap.AttributeMapForStandardAttribute]
//   - [IODRecordMap.SetAttributeMapForStandardAttribute]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap
type IODRecordMap interface {
	objectivec.IObject

	// Topic: Instance Properties

	Attributes() foundation.INSDictionary
	Native() string
	SetNative(value string)
	OdPredicate() foundation.INSDictionary
	SetOdPredicate(value foundation.INSDictionary)
	StandardAttributeTypes() foundation.INSArray

	// Topic: Instance Methods

	AttributeMapForStandardAttribute(standardAttribute string) IODAttributeMap
	SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute string)
}

// Init initializes the instance.
func (o ODRecordMap) Init() ODRecordMap {
	rv := objc.Send[ODRecordMap](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODRecordMap) Autorelease() ODRecordMap {
	rv := objc.Send[ODRecordMap](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODRecordMap creates a new ODRecordMap instance.
func NewODRecordMap() ODRecordMap {
	class := getODRecordMapClass()
	rv := objc.Send[ODRecordMap](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributeMap(forStandardAttribute:)
func (o ODRecordMap) AttributeMapForStandardAttribute(standardAttribute string) IODAttributeMap {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributeMapForStandardAttribute:"), objc.String(standardAttribute))
	return ODAttributeMapFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/setAttribute(_:forStandardAttribute:)
func (o ODRecordMap) SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute string) {
	objc.Send[objc.ID](o.ID, objc.Sel("setAttributeMap:forStandardAttribute:"), attributeMap, objc.String(standardAttribute))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-swift.property
func (o ODRecordMap) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o ODRecordMap) Native() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("native"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODRecordMap) SetNative(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setNative:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o ODRecordMap) OdPredicate() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("odPredicate"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (o ODRecordMap) SetOdPredicate(value foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("setOdPredicate:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/standardAttributeTypes
func (o ODRecordMap) StandardAttributeTypes() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("standardAttributeTypes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
