// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModelAssetResourceFactoryImpl protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryImpl
type MLModelAssetResourceFactoryImpl interface {
	objectivec.IObject
}

// MLModelAssetResourceFactoryImplObject wraps an existing Objective-C object that conforms to the MLModelAssetResourceFactoryImpl protocol.
type MLModelAssetResourceFactoryImplObject struct {
	objectivec.Object
}
func (o MLModelAssetResourceFactoryImplObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLModelAssetResourceFactoryImplObjectFromID constructs a [MLModelAssetResourceFactoryImplObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLModelAssetResourceFactoryImplObjectFromID(id objc.ID) MLModelAssetResourceFactoryImplObject {
	return MLModelAssetResourceFactoryImplObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryImpl/compiledModelURL
func (o MLModelAssetResourceFactoryImplObject) CompiledModelURL() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("compiledModelURL"))
	return objectivec.Object{ID: rv}
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryImpl/modelAssetDescriptionWithError:
func (o MLModelAssetResourceFactoryImplObject) ModelAssetDescriptionWithError() (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("modelAssetDescriptionWithError:"))
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryImpl/modelStructureWithError:
func (o MLModelAssetResourceFactoryImplObject) ModelStructureWithError() (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("modelStructureWithError:"))
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryImpl/modelWithConfiguration:error:
func (o MLModelAssetResourceFactoryImplObject) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("modelWithConfiguration:error:"), configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}

