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
func (m MLGKDecisionTree) Init() MLGKDecisionTree {
	rv := objc.Send[MLGKDecisionTree](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGKDecisionTree) Autorelease() MLGKDecisionTree {
	rv := objc.Send[MLGKDecisionTree](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGKDecisionTree creates a new MLGKDecisionTree instance.
func NewMLGKDecisionTree() MLGKDecisionTree {
	class := getMLGKDecisionTreeClass()
	rv := objc.Send[MLGKDecisionTree](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLGKDecisionTree) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (m MLGKDecisionTree) _initWithFlattenedTree(tree objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_initWithFlattenedTree:"), tree)
	return objectivec.Object{ID: rv}
}

// InitWithFlattenedTree is an exported wrapper for the private method _initWithFlattenedTree.
func (m MLGKDecisionTree) InitWithFlattenedTree(tree objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_initWithFlattenedTree:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithFlattenedTree:"}
		return nil, err
	}
	return m._initWithFlattenedTree(tree), nil
}

// CanInitWithFlattenedTree reports whether the receiver responds to the private selector _initWithFlattenedTree:.
func (m MLGKDecisionTree) CanInitWithFlattenedTree() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_initWithFlattenedTree:"))
}
func (m MLGKDecisionTree) _loadModelAssetWithModelAtPathWithError(path objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_loadModelAssetWithModelAtPath:withError:"), path, error_)
}

// LoadModelAssetWithModelAtPathWithError is an exported wrapper for the private method _loadModelAssetWithModelAtPathWithError.
func (m MLGKDecisionTree) LoadModelAssetWithModelAtPathWithError(path objectivec.IObject, error_ objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_loadModelAssetWithModelAtPath:withError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_loadModelAssetWithModelAtPath:withError:"}
		return err
	}
	m._loadModelAssetWithModelAtPathWithError(path, error_)
	return nil
}

// CanLoadModelAssetWithModelAtPathWithError reports whether the receiver responds to the private selector _loadModelAssetWithModelAtPath:withError:.
func (m MLGKDecisionTree) CanLoadModelAssetWithModelAtPathWithError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_loadModelAssetWithModelAtPath:withError:"))
}
func (m MLGKDecisionTree) _makeInferenceFromAnswersWithError(answers objectivec.IObject, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_makeInferenceFromAnswers:withError:"), answers, error_)
	return objectivec.Object{ID: rv}
}

// MakeInferenceFromAnswersWithError is an exported wrapper for the private method _makeInferenceFromAnswersWithError.
func (m MLGKDecisionTree) MakeInferenceFromAnswersWithError(answers objectivec.IObject, error_ objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_makeInferenceFromAnswers:withError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_makeInferenceFromAnswers:withError:"}
		return nil, err
	}
	return m._makeInferenceFromAnswersWithError(answers, error_), nil
}

// CanMakeInferenceFromAnswersWithError reports whether the receiver responds to the private selector _makeInferenceFromAnswers:withError:.
func (m MLGKDecisionTree) CanMakeInferenceFromAnswersWithError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_makeInferenceFromAnswers:withError:"))
}
func (m MLGKDecisionTree) _saveModelAssetWithModelToPathWithError(path objectivec.IObject, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_saveModelAssetWithModelToPath:withError:"), path, error_)
	return rv
}

// SaveModelAssetWithModelToPathWithError is an exported wrapper for the private method _saveModelAssetWithModelToPathWithError.
func (m MLGKDecisionTree) SaveModelAssetWithModelToPathWithError(path objectivec.IObject, error_ objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_saveModelAssetWithModelToPath:withError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_saveModelAssetWithModelToPath:withError:"}
		return false, err
	}
	return m._saveModelAssetWithModelToPathWithError(path, error_), nil
}

// CanSaveModelAssetWithModelToPathWithError reports whether the receiver responds to the private selector _saveModelAssetWithModelToPath:withError:.
func (m MLGKDecisionTree) CanSaveModelAssetWithModelToPathWithError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_saveModelAssetWithModelToPath:withError:"))
}

func (m MLGKDecisionTree) _attributes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_attributes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanAttributes reports whether the receiver responds to the private selector _attributes.
func (m MLGKDecisionTree) CanAttributes() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_attributes"))
}

// Attributes is an exported wrapper for the private property _attributes.
func (m MLGKDecisionTree) Attributes() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_attributes")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_attributes"}
	}
	return m._attributes(), nil
}
func (m MLGKDecisionTree) Set_attributes(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("set_attributes:"), value)
}
func (m MLGKDecisionTree) _objectStore() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_objectStore"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// CanObjectStore reports whether the receiver responds to the private selector _objectStore.
func (m MLGKDecisionTree) CanObjectStore() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_objectStore"))
}

// ObjectStore is an exported wrapper for the private property _objectStore.
func (m MLGKDecisionTree) ObjectStore() (foundation.INSOrderedSet, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_objectStore")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_objectStore"}
	}
	return m._objectStore(), nil
}
func (m MLGKDecisionTree) Set_objectStore(value foundation.INSOrderedSet) {
	objc.Send[struct{}](m.ID, objc.Sel("set_objectStore:"), value)
}
