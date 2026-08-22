// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLKey.DeletingPrefixingScope]
//   - [MLKey.HasGlobalScope]
//   - [MLKey.HasSameNameAsKey]
//   - [MLKey.ScopedTo]
//   - [MLKey.InitWithKeyName]
//   - [MLKey.InitWithKeyNameScope]
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
//   - [IMLKey.InitWithKeyName]
//   - [IMLKey.InitWithKeyNameScope]
type IMLKey interface {
	objectivec.IObject

	// Topic: Methods

	DeletingPrefixingScope(scope objectivec.IObject) objectivec.IObject
	HasGlobalScope() bool
	HasSameNameAsKey(key objectivec.IObject) bool
	ScopedTo(to objectivec.IObject) objectivec.IObject
	InitWithKeyName(name objectivec.IObject) MLKey
	InitWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey
}

// Init initializes the instance.
func (m MLKey) Init() MLKey {
	rv := objc.SendIfResponds[MLKey](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLKey) Autorelease() MLKey {
	rv := objc.SendIfResponds[MLKey](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKey creates a new MLKey instance.
func NewMLKey() MLKey {
	class := getMLKeyClass()
	rv := objc.SendIfResponds[MLKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewKeyWithKeyName(name objectivec.IObject) MLKey {
	instance := getMLKeyClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithKeyName:"), name)
	return MLKeyFromID(rv)
}

func NewKeyWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey {
	instance := getMLKeyClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return MLKeyFromID(rv)
}

func (m MLKey) DeletingPrefixingScope(scope objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("deletingPrefixingScope:"), scope)
	return objectivec.Object{ID: rv}
}
func (m MLKey) HasGlobalScope() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGlobalScope"))
	return rv
}
func (m MLKey) HasSameNameAsKey(key objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasSameNameAsKey:"), key)
	return rv
}
func (m MLKey) ScopedTo(to objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("scopedTo:"), to)
	return objectivec.Object{ID: rv}
}
func (m MLKey) InitWithKeyName(name objectivec.IObject) MLKey {
	rv := objc.SendIfResponds[MLKey](m.ID, objc.Sel("initWithKeyName:"), name)
	return rv
}
func (m MLKey) InitWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLKey {
	rv := objc.SendIfResponds[MLKey](m.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return rv
}

func (_MLKeyClass MLKeyClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLKeyClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
