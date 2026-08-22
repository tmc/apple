// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODAttributeMap] class.
var (
	_ODAttributeMapClass     ODAttributeMapClass
	_ODAttributeMapClassOnce sync.Once
)

func getODAttributeMapClass() ODAttributeMapClass {
	_ODAttributeMapClassOnce.Do(func() {
		_ODAttributeMapClass = ODAttributeMapClass{class: objc.GetClass("ODAttributeMap")}
	})
	return _ODAttributeMapClass
}

// GetODAttributeMapClass returns the class object for ODAttributeMap.
func GetODAttributeMapClass() ODAttributeMapClass {
	return getODAttributeMapClass()
}

type ODAttributeMapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODAttributeMapClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODAttributeMapClass) Alloc() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [ODAttributeMap.CustomAttributes]
//   - [ODAttributeMap.SetCustomAttributes]
//   - [ODAttributeMap.CustomQueryFunction]
//   - [ODAttributeMap.SetCustomQueryFunction]
//   - [ODAttributeMap.CustomTranslationFunction]
//   - [ODAttributeMap.SetCustomTranslationFunction]
//   - [ODAttributeMap.Value]
//   - [ODAttributeMap.SetValue]
//
// # Instance Methods
//
//   - [ODAttributeMap.SetStaticValue]
//   - [ODAttributeMap.SetVariableSubstitution]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap
type ODAttributeMap struct {
	objectivec.Object
}

// ODAttributeMapFromID constructs a [ODAttributeMap] from an objc.ID.
func ODAttributeMapFromID(id objc.ID) ODAttributeMap {
	return ODAttributeMap{objectivec.Object{ID: id}}
}

// NOTE: ODAttributeMap adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODAttributeMap] class.
//
// # Instance Properties
//
//   - [IODAttributeMap.CustomAttributes]
//   - [IODAttributeMap.SetCustomAttributes]
//   - [IODAttributeMap.CustomQueryFunction]
//   - [IODAttributeMap.SetCustomQueryFunction]
//   - [IODAttributeMap.CustomTranslationFunction]
//   - [IODAttributeMap.SetCustomTranslationFunction]
//   - [IODAttributeMap.Value]
//   - [IODAttributeMap.SetValue]
//
// # Instance Methods
//
//   - [IODAttributeMap.SetStaticValue]
//   - [IODAttributeMap.SetVariableSubstitution]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap
type IODAttributeMap interface {
	objectivec.IObject

	// Topic: Instance Properties

	CustomAttributes() foundation.INSArray
	SetCustomAttributes(value foundation.INSArray)
	CustomQueryFunction() string
	SetCustomQueryFunction(value string)
	CustomTranslationFunction() string
	SetCustomTranslationFunction(value string)
	Value() string
	SetValue(value string)

	// Topic: Instance Methods

	SetStaticValue(staticValue string)
	SetVariableSubstitution(variableSubstitution string)
}

// Init initializes the instance.
func (o ODAttributeMap) Init() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODAttributeMap) Autorelease() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODAttributeMap creates a new ODAttributeMap instance.
func NewODAttributeMap() ODAttributeMap {
	class := getODAttributeMapClass()
	rv := objc.Send[ODAttributeMap](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(staticValue:)
func NewODAttributeMapWithStaticValue(staticValue string) ODAttributeMap {
	rv := objc.Send[objc.ID](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithStaticValue:"), objc.String(staticValue))
	return ODAttributeMapFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(value:)
func NewODAttributeMapWithValue(value string) ODAttributeMap {
	rv := objc.Send[objc.ID](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithValue:"), objc.String(value))
	return ODAttributeMapFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setStaticValue(_:)
func (o ODAttributeMap) SetStaticValue(staticValue string) {
	objc.Send[objc.ID](o.ID, objc.Sel("setStaticValue:"), objc.String(staticValue))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setVariableSubstitution(_:)
func (o ODAttributeMap) SetVariableSubstitution(variableSubstitution string) {
	objc.Send[objc.ID](o.ID, objc.Sel("setVariableSubstitution:"), objc.String(variableSubstitution))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-swift.property
func (o ODAttributeMap) CustomAttributes() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("customAttributes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODAttributeMap) SetCustomAttributes(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setCustomAttributes:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-swift.property
func (o ODAttributeMap) CustomQueryFunction() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("customQueryFunction"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODAttributeMap) SetCustomQueryFunction(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setCustomQueryFunction:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-swift.property
func (o ODAttributeMap) CustomTranslationFunction() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("customTranslationFunction"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODAttributeMap) SetCustomTranslationFunction(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setCustomTranslationFunction:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-swift.property
func (o ODAttributeMap) Value() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("value"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODAttributeMap) SetValue(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setValue:"), objc.String(value))
}
