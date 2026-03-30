// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGKDecisionTree] class.
var (
	_MLGKDecisionTreeClass     MLGKDecisionTreeClass
	_MLGKDecisionTreeClassOnce sync.Once
)

func getMLGKDecisionTreeClass() MLGKDecisionTreeClass {
	_MLGKDecisionTreeClassOnce.Do(func() {
		_MLGKDecisionTreeClass = MLGKDecisionTreeClass{class: objc.GetClass("MLGKDecisionTree")}
	})
	return _MLGKDecisionTreeClass
}

// GetMLGKDecisionTreeClass returns the class object for MLGKDecisionTree.
func GetMLGKDecisionTreeClass() MLGKDecisionTreeClass {
	return getMLGKDecisionTreeClass()
}

type MLGKDecisionTreeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGKDecisionTreeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGKDecisionTreeClass) Alloc() MLGKDecisionTree {
	rv := objc.Send[MLGKDecisionTree](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGKDecisionTree._attributes]
//   - [MLGKDecisionTree.Set_attributes]
//   - [MLGKDecisionTree._init]
//   - [MLGKDecisionTree._initWithFlattenedTree]
//   - [MLGKDecisionTree._loadModelAssetWithModelAtPathWithError]
//   - [MLGKDecisionTree._makeInferenceFromAnswersWithError]
//   - [MLGKDecisionTree._objectStore]
//   - [MLGKDecisionTree.Set_objectStore]
//   - [MLGKDecisionTree._saveModelAssetWithModelToPathWithError]
//
// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree
type MLGKDecisionTree struct {
	objectivec.Object
}

// MLGKDecisionTreeFromID constructs a [MLGKDecisionTree] from an objc.ID.
func MLGKDecisionTreeFromID(id objc.ID) MLGKDecisionTree {
	return MLGKDecisionTree{objectivec.Object{ID: id}}
}

// Ensure MLGKDecisionTree implements IMLGKDecisionTree.
var _ IMLGKDecisionTree = MLGKDecisionTree{}

// An interface definition for the [MLGKDecisionTree] class.
//
// # Methods
//
//   - [IMLGKDecisionTree._attributes]
//   - [IMLGKDecisionTree.Set_attributes]
//   - [IMLGKDecisionTree._init]
//   - [IMLGKDecisionTree._initWithFlattenedTree]
//   - [IMLGKDecisionTree._loadModelAssetWithModelAtPathWithError]
//   - [IMLGKDecisionTree._makeInferenceFromAnswersWithError]
//   - [IMLGKDecisionTree._objectStore]
//   - [IMLGKDecisionTree.Set_objectStore]
//   - [IMLGKDecisionTree._saveModelAssetWithModelToPathWithError]
//
// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree
type IMLGKDecisionTree interface {
	objectivec.IObject

	// Topic: Methods

	_attributes() foundation.INSArray
	Set_attributes(value foundation.INSArray)
	_init() objectivec.IObject
	_initWithFlattenedTree(tree objectivec.IObject) objectivec.IObject
	_loadModelAssetWithModelAtPathWithError(path objectivec.IObject, error_ objectivec.IObject)
	_makeInferenceFromAnswersWithError(answers objectivec.IObject, error_ objectivec.IObject) objectivec.IObject
	_objectStore() foundation.INSOrderedSet
	Set_objectStore(value foundation.INSOrderedSet)
	_saveModelAssetWithModelToPathWithError(path objectivec.IObject, error_ objectivec.IObject) bool
}

// Init initializes the instance.
func (g MLGKDecisionTree) Init() MLGKDecisionTree {
	rv := objc.Send[MLGKDecisionTree](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGKDecisionTree) Autorelease() MLGKDecisionTree {
	rv := objc.Send[MLGKDecisionTree](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGKDecisionTree creates a new MLGKDecisionTree instance.
func NewMLGKDecisionTree() MLGKDecisionTree {
	class := getMLGKDecisionTreeClass()
	rv := objc.Send[MLGKDecisionTree](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_init
func (g MLGKDecisionTree) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_initWithFlattenedTree:
func (g MLGKDecisionTree) _initWithFlattenedTree(tree objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_initWithFlattenedTree:"), tree)
	return objectivec.Object{ID: rv}
}

// InitWithFlattenedTree is an exported wrapper for the private method _initWithFlattenedTree.
func (g MLGKDecisionTree) InitWithFlattenedTree(tree objectivec.IObject) objectivec.IObject {
	return g._initWithFlattenedTree(tree)
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_loadModelAssetWithModelAtPath:withError:
func (g MLGKDecisionTree) _loadModelAssetWithModelAtPathWithError(path objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_loadModelAssetWithModelAtPath:withError:"), path, error_)
}

// LoadModelAssetWithModelAtPathWithError is an exported wrapper for the private method _loadModelAssetWithModelAtPathWithError.
func (g MLGKDecisionTree) LoadModelAssetWithModelAtPathWithError(path objectivec.IObject, error_ objectivec.IObject) {
	g._loadModelAssetWithModelAtPathWithError(path, error_)
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_makeInferenceFromAnswers:withError:
func (g MLGKDecisionTree) _makeInferenceFromAnswersWithError(answers objectivec.IObject, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_makeInferenceFromAnswers:withError:"), answers, error_)
	return objectivec.Object{ID: rv}
}

// MakeInferenceFromAnswersWithError is an exported wrapper for the private method _makeInferenceFromAnswersWithError.
func (g MLGKDecisionTree) MakeInferenceFromAnswersWithError(answers objectivec.IObject, error_ objectivec.IObject) objectivec.IObject {
	return g._makeInferenceFromAnswersWithError(answers, error_)
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_saveModelAssetWithModelToPath:withError:
func (g MLGKDecisionTree) _saveModelAssetWithModelToPathWithError(path objectivec.IObject, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_saveModelAssetWithModelToPath:withError:"), path, error_)
	return rv
}

// SaveModelAssetWithModelToPathWithError is an exported wrapper for the private method _saveModelAssetWithModelToPathWithError.
func (g MLGKDecisionTree) SaveModelAssetWithModelToPathWithError(path objectivec.IObject, error_ objectivec.IObject) bool {
	return g._saveModelAssetWithModelToPathWithError(path, error_)
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_attributes
func (g MLGKDecisionTree) _attributes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_attributes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MLGKDecisionTree) Set_attributes(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("set_attributes:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLGKDecisionTree/_objectStore
func (g MLGKDecisionTree) _objectStore() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_objectStore"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (g MLGKDecisionTree) Set_objectStore(value foundation.INSOrderedSet) {
	objc.Send[struct{}](g.ID, objc.Sel("set_objectStore:"), value)
}
