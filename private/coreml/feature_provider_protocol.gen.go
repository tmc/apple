// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLFeatureProvider protocol.
type MLFeatureProvider interface {
	objectivec.IObject

	// FeatureNames protocol.
	FeatureNames() objectivec.IObject

	// FeatureValueForName protocol.
	FeatureValueForName(name objectivec.IObject) objectivec.IObject
}

// MLFeatureProviderObject wraps an existing Objective-C object that conforms to the MLFeatureProvider protocol.
type MLFeatureProviderObject struct {
	objectivec.Object
}

func (o MLFeatureProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLFeatureProviderObjectFromID constructs a [MLFeatureProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLFeatureProviderObjectFromID(id objc.ID) MLFeatureProviderObject {
	return MLFeatureProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLFeatureProviderObject) FeatureNames() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("featureNames"))
	return objectivec.Object{ID: rv}
}
func (o MLFeatureProviderObject) FeatureValueForName(name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("featureValueForName:"), name)
	return objectivec.Object{ID: rv}
}
