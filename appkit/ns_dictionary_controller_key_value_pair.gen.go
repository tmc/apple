// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDictionaryControllerKeyValuePair] class.
var (
	_NSDictionaryControllerKeyValuePairClass     NSDictionaryControllerKeyValuePairClass
	_NSDictionaryControllerKeyValuePairClassOnce sync.Once
)

func getNSDictionaryControllerKeyValuePairClass() NSDictionaryControllerKeyValuePairClass {
	_NSDictionaryControllerKeyValuePairClassOnce.Do(func() {
		_NSDictionaryControllerKeyValuePairClass = NSDictionaryControllerKeyValuePairClass{class: objc.GetClass("NSDictionaryControllerKeyValuePair")}
	})
	return _NSDictionaryControllerKeyValuePairClass
}

// GetNSDictionaryControllerKeyValuePairClass returns the class object for NSDictionaryControllerKeyValuePair.
func GetNSDictionaryControllerKeyValuePairClass() NSDictionaryControllerKeyValuePairClass {
	return getNSDictionaryControllerKeyValuePairClass()
}

type NSDictionaryControllerKeyValuePairClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDictionaryControllerKeyValuePairClass) Alloc() NSDictionaryControllerKeyValuePair {
	rv := objc.Send[NSDictionaryControllerKeyValuePair](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A set of methods implemented by arranged objects to give access to
// information about those objects.
//
// # Overview
// 
// [NSDictionaryControllerKeyValuePair] is an informal protocol that is
// implemented by objects returned by the [NSDictionaryController] method
// arrangedObjects. See [NSDictionaryController] for more information.
//
// # Instance Properties
//
//   - [NSDictionaryControllerKeyValuePair.ExplicitlyIncluded]
//   - [NSDictionaryControllerKeyValuePair.Key]
//   - [NSDictionaryControllerKeyValuePair.SetKey]
//   - [NSDictionaryControllerKeyValuePair.LocalizedKey]
//   - [NSDictionaryControllerKeyValuePair.SetLocalizedKey]
//   - [NSDictionaryControllerKeyValuePair.Value]
//   - [NSDictionaryControllerKeyValuePair.SetValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair
type NSDictionaryControllerKeyValuePair struct {
	objectivec.Object
}

// NSDictionaryControllerKeyValuePairFromID constructs a [NSDictionaryControllerKeyValuePair] from an objc.ID.
//
// A set of methods implemented by arranged objects to give access to
// information about those objects.
func NSDictionaryControllerKeyValuePairFromID(id objc.ID) NSDictionaryControllerKeyValuePair {
	return NSDictionaryControllerKeyValuePair{objectivec.Object{ID: id}}
}
// NOTE: NSDictionaryControllerKeyValuePair adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDictionaryControllerKeyValuePair] class.
//
// # Instance Properties
//
//   - [INSDictionaryControllerKeyValuePair.ExplicitlyIncluded]
//   - [INSDictionaryControllerKeyValuePair.Key]
//   - [INSDictionaryControllerKeyValuePair.SetKey]
//   - [INSDictionaryControllerKeyValuePair.LocalizedKey]
//   - [INSDictionaryControllerKeyValuePair.SetLocalizedKey]
//   - [INSDictionaryControllerKeyValuePair.Value]
//   - [INSDictionaryControllerKeyValuePair.SetValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair
type INSDictionaryControllerKeyValuePair interface {
	objectivec.IObject

	// Topic: Instance Properties

	ExplicitlyIncluded() bool
	Key() string
	SetKey(value string)
	LocalizedKey() string
	SetLocalizedKey(value string)
	Value() objectivec.IObject
	SetValue(value objectivec.IObject)
}

// Init initializes the instance.
func (d NSDictionaryControllerKeyValuePair) Init() NSDictionaryControllerKeyValuePair {
	rv := objc.Send[NSDictionaryControllerKeyValuePair](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDictionaryControllerKeyValuePair) Autorelease() NSDictionaryControllerKeyValuePair {
	rv := objc.Send[NSDictionaryControllerKeyValuePair](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDictionaryControllerKeyValuePair creates a new NSDictionaryControllerKeyValuePair instance.
func NewNSDictionaryControllerKeyValuePair() NSDictionaryControllerKeyValuePair {
	class := getNSDictionaryControllerKeyValuePairClass()
	rv := objc.Send[NSDictionaryControllerKeyValuePair](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/isExplicitlyIncluded
func (d NSDictionaryControllerKeyValuePair) ExplicitlyIncluded() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isExplicitlyIncluded"))
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/key
func (d NSDictionaryControllerKeyValuePair) Key() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDictionaryControllerKeyValuePair) SetKey(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/localizedKey
func (d NSDictionaryControllerKeyValuePair) LocalizedKey() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("localizedKey"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDictionaryControllerKeyValuePair) SetLocalizedKey(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocalizedKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/value
func (d NSDictionaryControllerKeyValuePair) Value() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("value"))
	return objectivec.Object{ID: rv}
}
func (d NSDictionaryControllerKeyValuePair) SetValue(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setValue:"), value)
}

