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

// Alloc allocates memory for a new instance of the class.
func (mc MLKeyClass) Alloc() MLKey {
	rv := objc.Send[MLKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for machine learning key types.
//
// # Overview
// 
// You don’t create use this class directly. Instead, use a class that
// inherits from this one, such as [MLParameterKey] or [MLMetricKey].
//
// # Retrieving a key’s information
//
//   - [MLKey.Name]: The name of the machine learning key.
//   - [MLKey.Scope]: The applicable scope of the machine learning key.
//
// See: https://developer.apple.com/documentation/CoreML/MLKey
type MLKey struct {
	objectivec.Object
}

// MLKeyFromID constructs a [MLKey] from an objc.ID.
//
// An abstract base class for machine learning key types.
func MLKeyFromID(id objc.ID) MLKey {
	return MLKey{objectivec.Object{ID: id}}
}
// NOTE: MLKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLKey] class.
//
// # Retrieving a key’s information
//
//   - [IMLKey.Name]: The name of the machine learning key.
//   - [IMLKey.Scope]: The applicable scope of the machine learning key.
//
// See: https://developer.apple.com/documentation/CoreML/MLKey
type IMLKey interface {
	objectivec.IObject

	// Topic: Retrieving a key’s information

	// The name of the machine learning key.
	Name() string
	// The applicable scope of the machine learning key.
	Scope() string

	EncodeWithCoder(coder foundation.INSCoder)
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

func (k MLKey) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](k.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The name of the machine learning key.
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/name
func (k MLKey) Name() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The applicable scope of the machine learning key.
//
// See: https://developer.apple.com/documentation/CoreML/MLKey/scope
func (k MLKey) Scope() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("scope"))
	return foundation.NSStringFromID(rv).String()
}

