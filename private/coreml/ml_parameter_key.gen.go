// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLParameterKey] class.
var (
	_MLParameterKeyClass     MLParameterKeyClass
	_MLParameterKeyClassOnce sync.Once
)

func getMLParameterKeyClass() MLParameterKeyClass {
	_MLParameterKeyClassOnce.Do(func() {
		_MLParameterKeyClass = MLParameterKeyClass{class: objc.GetClass("MLParameterKey")}
	})
	return _MLParameterKeyClass
}

// GetMLParameterKeyClass returns the class object for MLParameterKey.
func GetMLParameterKeyClass() MLParameterKeyClass {
	return getMLParameterKeyClass()
}

type MLParameterKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLParameterKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLParameterKeyClass) Alloc() MLParameterKey {
	rv := objc.Send[MLParameterKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLParameterKey
type MLParameterKey struct {
	MLKey
}

// MLParameterKeyFromID constructs a [MLParameterKey] from an objc.ID.
func MLParameterKeyFromID(id objc.ID) MLParameterKey {
	return MLParameterKey{MLKey: MLKeyFromID(id)}
}
// Ensure MLParameterKey implements IMLParameterKey.
var _ IMLParameterKey = MLParameterKey{}

// An interface definition for the [MLParameterKey] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey
type IMLParameterKey interface {
	IMLKey
}

// Init initializes the instance.
func (p MLParameterKey) Init() MLParameterKey {
	rv := objc.Send[MLParameterKey](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLParameterKey) Autorelease() MLParameterKey {
	rv := objc.Send[MLParameterKey](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterKey creates a new MLParameterKey instance.
func NewMLParameterKey() MLParameterKey {
	class := getMLParameterKeyClass()
	rv := objc.Send[MLParameterKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithCoder:
func NewParameterKeyWithCoder(coder objectivec.IObject) MLParameterKey {
	instance := getMLParameterKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLParameterKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/initWithKeyName:
func NewParameterKeyWithKeyName(name objectivec.IObject) MLParameterKey {
	instance := getMLParameterKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:"), name)
	return MLParameterKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:scope:
func NewParameterKeyWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLParameterKey {
	instance := getMLParameterKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return MLParameterKeyFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/maxDepth
func (_MLParameterKeyClass MLParameterKeyClass) MaxDepth() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("maxDepth"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/minChildWeight
func (_MLParameterKeyClass MLParameterKeyClass) MinChildWeight() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("minChildWeight"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/numClasses
func (_MLParameterKeyClass MLParameterKeyClass) NumClasses() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("numClasses"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/numTrees
func (_MLParameterKeyClass MLParameterKeyClass) NumTrees() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("numTrees"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/objective
func (_MLParameterKeyClass MLParameterKeyClass) Objective() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("objective"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/precisionRecallCurves
func (_MLParameterKeyClass MLParameterKeyClass) PrecisionRecallCurves() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("precisionRecallCurves"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/updateType
func (_MLParameterKeyClass MLParameterKeyClass) UpdateType() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("updateType"))
	return objectivec.Object{ID: rv}
}

