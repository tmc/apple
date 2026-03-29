// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLKey] class.
var (
	_MLKeyClass     MLKeyClass
	_MLKeyClassOnce sync.Once
)

func getMLKeyClass() MLKeyClass {
	_MLKeyClassOnce.Do(func() {
		_MLKeyClass = MLKeyClass{class: objc.GetClass("MLKey")}
	})
	return _MLKeyClass
}

// GetMLKeyClass returns the class object for MLKey.
func GetMLKeyClass() MLKeyClass {
	return getMLKeyClass()
}

type MLKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLKeyClass) Alloc() MLKey {
	rv := objc.Send[MLKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLKey.DeletingPrefixingScope]
//   - [MLKey.HasGlobalScope]
//   - [MLKey.HasSameNameAsKey]
//   - [MLKey.ScopedTo]
//   - [MLKey.InitWithCoder]
//   - [MLKey.InitWithKeyName]
//   - [MLKey.InitWithKeyNameScope]
// See: https://developer.apple.com/documentation/CoreML/MLKey
type MLKey struct {
	objectivec.Object
}

// MLKeyFromID constructs a [MLKey] from an objc.ID.
func MLKeyFromID(id objc.ID) MLKey {
	return MLKey{objectivec.Object{ID: id}}
}
// Ensure MLKey implements IMLKey.
var _ IMLKey = MLKey{}

// An interface definition for the [MLKey] class.
//
// # Methods
//
//   - [IMLKey.DeletingPrefixingScope]
//   - [IMLKey.HasGlobalScope]
//   - [IMLKey.HasSameNameAsKey]
//   - [IMLKey.ScopedTo]
//   - [IMLKey.InitWithCoder]
//   - [IMLKey.InitWithKeyName]
//   - [IMLKey.InitWithKeyNameScope]
//
// See: https://developer.apple.com/documentation/CoreML/MLKey
type IMLKey interface {
	objectivec.IObject

	// Topic: Methods

	DeletingPrefixingScope(scope objectivec.IObject) objectivec.IObject
	HasGlobalScope() bool
	HasSameNameAsKey(key objectivec.IObject) bool
	ScopedTo(to objectivec.IObject) objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) MLKey
	InitWithKeyName(name objectivec.IObject) MLKey
	InitWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey
}

// Init initializes the instance.
func (k MLKey) Init() MLKey {
	rv := objc.Send[MLKey](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MLKey) Autorelease() MLKey {
	rv := objc.Send[MLKey](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKey creates a new MLKey instance.
func NewMLKey() MLKey {
	class := getMLKeyClass()
	rv := objc.Send[MLKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithCoder:
func NewKeyWithCoder(coder objectivec.IObject) MLKey {
	instance := getMLKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:
func NewKeyWithKeyName(name objectivec.IObject) MLKey {
	instance := getMLKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:"), name)
	return MLKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:scope:
func NewKeyWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey {
	instance := getMLKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return MLKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/deletingPrefixingScope:
func (k MLKey) DeletingPrefixingScope(scope objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("deletingPrefixingScope:"), scope)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLKey/hasGlobalScope
func (k MLKey) HasGlobalScope() bool {
	rv := objc.Send[bool](k.ID, objc.Sel("hasGlobalScope"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/hasSameNameAsKey:
func (k MLKey) HasSameNameAsKey(key objectivec.IObject) bool {
	rv := objc.Send[bool](k.ID, objc.Sel("hasSameNameAsKey:"), key)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/scopedTo:
func (k MLKey) ScopedTo(to objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("scopedTo:"), to)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithCoder:
func (k MLKey) InitWithCoder(coder foundation.INSCoder) MLKey {
	rv := objc.Send[MLKey](k.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:
func (k MLKey) InitWithKeyName(name objectivec.IObject) MLKey {
	rv := objc.Send[MLKey](k.ID, objc.Sel("initWithKeyName:"), name)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:scope:
func (k MLKey) InitWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey {
	rv := objc.Send[MLKey](k.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKey/supportsSecureCoding
func (_MLKeyClass MLKeyClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLKeyClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

