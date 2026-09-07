// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelTensorSharedEventRegistry] class.
var (
	_MLModelTensorSharedEventRegistryClass     MLModelTensorSharedEventRegistryClass
	_MLModelTensorSharedEventRegistryClassOnce sync.Once
)

func getMLModelTensorSharedEventRegistryClass() MLModelTensorSharedEventRegistryClass {
	_MLModelTensorSharedEventRegistryClassOnce.Do(func() {
		_MLModelTensorSharedEventRegistryClass = MLModelTensorSharedEventRegistryClass{class: objc.GetClass("_TtC6CoreMLP33_F670D341DBCB097E94978F9EEEE1312D32MLModelTensorSharedEventRegistry")}
	})
	return _MLModelTensorSharedEventRegistryClass
}

// GetMLModelTensorSharedEventRegistryClass returns the class object for _TtC6CoreMLP33_F670D341DBCB097E94978F9EEEE1312D32MLModelTensorSharedEventRegistry.
func GetMLModelTensorSharedEventRegistryClass() MLModelTensorSharedEventRegistryClass {
	return getMLModelTensorSharedEventRegistryClass()
}

type MLModelTensorSharedEventRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelTensorSharedEventRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelTensorSharedEventRegistryClass) Alloc() MLModelTensorSharedEventRegistry {
	rv := objc.SendIfResponds[MLModelTensorSharedEventRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLModelTensorSharedEventRegistry struct {
	objectivec.Object
}

// MLModelTensorSharedEventRegistryFromID constructs a [MLModelTensorSharedEventRegistry] from an objc.ID.
func MLModelTensorSharedEventRegistryFromID(id objc.ID) MLModelTensorSharedEventRegistry {
	return MLModelTensorSharedEventRegistry{objectivec.Object{ID: id}}
}

// Ensure MLModelTensorSharedEventRegistry implements IMLModelTensorSharedEventRegistry.
var _ IMLModelTensorSharedEventRegistry = MLModelTensorSharedEventRegistry{}

// An interface definition for the [MLModelTensorSharedEventRegistry] class.
type IMLModelTensorSharedEventRegistry interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelTensorSharedEventRegistry) Init() MLModelTensorSharedEventRegistry {
	rv := objc.SendIfResponds[MLModelTensorSharedEventRegistry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelTensorSharedEventRegistry) Autorelease() MLModelTensorSharedEventRegistry {
	rv := objc.SendIfResponds[MLModelTensorSharedEventRegistry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelTensorSharedEventRegistry creates a new MLModelTensorSharedEventRegistry instance.
func NewMLModelTensorSharedEventRegistry() MLModelTensorSharedEventRegistry {
	class := getMLModelTensorSharedEventRegistryClass()
	rv := objc.SendIfResponds[MLModelTensorSharedEventRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}
