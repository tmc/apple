// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPropertyMapping] class.
var (
	_NSPropertyMappingClass     NSPropertyMappingClass
	_NSPropertyMappingClassOnce sync.Once
)

func getNSPropertyMappingClass() NSPropertyMappingClass {
	_NSPropertyMappingClassOnce.Do(func() {
		_NSPropertyMappingClass = NSPropertyMappingClass{class: objc.GetClass("NSPropertyMapping")}
	})
	return _NSPropertyMappingClass
}

// GetNSPropertyMappingClass returns the class object for NSPropertyMapping.
func GetNSPropertyMappingClass() NSPropertyMappingClass {
	return getNSPropertyMappingClass()
}

type NSPropertyMappingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPropertyMappingClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPropertyMappingClass) Alloc() NSPropertyMapping {
	rv := objc.Send[NSPropertyMapping](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mapping instance that specifies in a model how to map from a property in
// a source entity to a property in a destination entity.
//
// # Managing Mapping Attributes
//
//   - [NSPropertyMapping.Name]: The name of the property in the destination entity for the property mapping.
//   - [NSPropertyMapping.SetName]
//   - [NSPropertyMapping.ValueExpression]: The value expression for the property mapping.
//   - [NSPropertyMapping.SetValueExpression]
//   - [NSPropertyMapping.UserInfo]: The user info for the property mapping.
//   - [NSPropertyMapping.SetUserInfo]
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyMapping
type NSPropertyMapping struct {
	objectivec.Object
}

// NSPropertyMappingFromID constructs a [NSPropertyMapping] from an objc.ID.
//
// A mapping instance that specifies in a model how to map from a property in
// a source entity to a property in a destination entity.
func NSPropertyMappingFromID(id objc.ID) NSPropertyMapping {
	return NSPropertyMapping{objectivec.Object{ID: id}}
}

// NOTE: NSPropertyMapping adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPropertyMapping] class.
//
// # Managing Mapping Attributes
//
//   - [INSPropertyMapping.Name]: The name of the property in the destination entity for the property mapping.
//   - [INSPropertyMapping.SetName]
//   - [INSPropertyMapping.ValueExpression]: The value expression for the property mapping.
//   - [INSPropertyMapping.SetValueExpression]
//   - [INSPropertyMapping.UserInfo]: The user info for the property mapping.
//   - [INSPropertyMapping.SetUserInfo]
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyMapping
type INSPropertyMapping interface {
	objectivec.IObject

	// Topic: Managing Mapping Attributes

	// The name of the property in the destination entity for the property mapping.
	Name() string
	SetName(value string)
	// The value expression for the property mapping.
	ValueExpression() foundation.NSExpression
	SetValueExpression(value foundation.NSExpression)
	// The user info for the property mapping.
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)
}

// Init initializes the instance.
func (p NSPropertyMapping) Init() NSPropertyMapping {
	rv := objc.Send[NSPropertyMapping](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPropertyMapping) Autorelease() NSPropertyMapping {
	rv := objc.Send[NSPropertyMapping](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPropertyMapping creates a new NSPropertyMapping instance.
func NewNSPropertyMapping() NSPropertyMapping {
	class := getNSPropertyMappingClass()
	rv := objc.Send[NSPropertyMapping](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the property in the destination entity for the property
// mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/name
func (p NSPropertyMapping) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPropertyMapping) SetName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setName:"), objc.String(value))
}

// The value expression for the property mapping.
//
// # Discussion
//
// The expression is used to create the value for the destination property.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/valueExpression
func (p NSPropertyMapping) ValueExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("valueExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}
func (p NSPropertyMapping) SetValueExpression(value foundation.NSExpression) {
	objc.Send[struct{}](p.ID, objc.Sel("setValueExpression:"), value)
}

// The user info for the property mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/userInfo
func (p NSPropertyMapping) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p NSPropertyMapping) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setUserInfo:"), value)
}
