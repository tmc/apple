// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetResourceFactoryInMemoryImpl] class.
var (
	_MLModelAssetResourceFactoryInMemoryImplClass     MLModelAssetResourceFactoryInMemoryImplClass
	_MLModelAssetResourceFactoryInMemoryImplClassOnce sync.Once
)

func getMLModelAssetResourceFactoryInMemoryImplClass() MLModelAssetResourceFactoryInMemoryImplClass {
	_MLModelAssetResourceFactoryInMemoryImplClassOnce.Do(func() {
		_MLModelAssetResourceFactoryInMemoryImplClass = MLModelAssetResourceFactoryInMemoryImplClass{class: objc.GetClass("MLModelAssetResourceFactoryInMemoryImpl")}
	})
	return _MLModelAssetResourceFactoryInMemoryImplClass
}

// GetMLModelAssetResourceFactoryInMemoryImplClass returns the class object for MLModelAssetResourceFactoryInMemoryImpl.
func GetMLModelAssetResourceFactoryInMemoryImplClass() MLModelAssetResourceFactoryInMemoryImplClass {
	return getMLModelAssetResourceFactoryInMemoryImplClass()
}

type MLModelAssetResourceFactoryInMemoryImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetResourceFactoryInMemoryImplClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetResourceFactoryInMemoryImplClass) Alloc() MLModelAssetResourceFactoryInMemoryImpl {
	rv := objc.Send[MLModelAssetResourceFactoryInMemoryImpl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelAssetResourceFactoryInMemoryImpl.CompiledModelURL]
//   - [MLModelAssetResourceFactoryInMemoryImpl.ModelAssetDescriptionWithError]
//   - [MLModelAssetResourceFactoryInMemoryImpl.ModelStructureWithError]
//   - [MLModelAssetResourceFactoryInMemoryImpl.ModelWithConfigurationError]
//   - [MLModelAssetResourceFactoryInMemoryImpl.InitWithArchiveData]
//   - [MLModelAssetResourceFactoryInMemoryImpl.DebugDescription]
//   - [MLModelAssetResourceFactoryInMemoryImpl.Description]
//   - [MLModelAssetResourceFactoryInMemoryImpl.Hash]
//   - [MLModelAssetResourceFactoryInMemoryImpl.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl
type MLModelAssetResourceFactoryInMemoryImpl struct {
	objectivec.Object
}

// MLModelAssetResourceFactoryInMemoryImplFromID constructs a [MLModelAssetResourceFactoryInMemoryImpl] from an objc.ID.
func MLModelAssetResourceFactoryInMemoryImplFromID(id objc.ID) MLModelAssetResourceFactoryInMemoryImpl {
	return MLModelAssetResourceFactoryInMemoryImpl{objectivec.Object{ID: id}}
}

// Ensure MLModelAssetResourceFactoryInMemoryImpl implements IMLModelAssetResourceFactoryInMemoryImpl.
var _ IMLModelAssetResourceFactoryInMemoryImpl = MLModelAssetResourceFactoryInMemoryImpl{}

// An interface definition for the [MLModelAssetResourceFactoryInMemoryImpl] class.
//
// # Methods
//
//   - [IMLModelAssetResourceFactoryInMemoryImpl.CompiledModelURL]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.ModelAssetDescriptionWithError]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.ModelStructureWithError]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.ModelWithConfigurationError]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.InitWithArchiveData]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.DebugDescription]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.Description]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.Hash]
//   - [IMLModelAssetResourceFactoryInMemoryImpl.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl
type IMLModelAssetResourceFactoryInMemoryImpl interface {
	objectivec.IObject

	// Topic: Methods

	CompiledModelURL() foundation.INSURL
	ModelAssetDescriptionWithError() (objectivec.IObject, error)
	ModelStructureWithError() (objectivec.IObject, error)
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	InitWithArchiveData(data objectivec.IObject) MLModelAssetResourceFactoryInMemoryImpl
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLModelAssetResourceFactoryInMemoryImpl) Init() MLModelAssetResourceFactoryInMemoryImpl {
	rv := objc.Send[MLModelAssetResourceFactoryInMemoryImpl](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetResourceFactoryInMemoryImpl) Autorelease() MLModelAssetResourceFactoryInMemoryImpl {
	rv := objc.Send[MLModelAssetResourceFactoryInMemoryImpl](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetResourceFactoryInMemoryImpl creates a new MLModelAssetResourceFactoryInMemoryImpl instance.
func NewMLModelAssetResourceFactoryInMemoryImpl() MLModelAssetResourceFactoryInMemoryImpl {
	class := getMLModelAssetResourceFactoryInMemoryImplClass()
	rv := objc.Send[MLModelAssetResourceFactoryInMemoryImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/initWithArchiveData:
func NewModelAssetResourceFactoryInMemoryImplWithArchiveData(data objectivec.IObject) MLModelAssetResourceFactoryInMemoryImpl {
	instance := getMLModelAssetResourceFactoryInMemoryImplClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArchiveData:"), data)
	return MLModelAssetResourceFactoryInMemoryImplFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/modelAssetDescriptionWithError:
func (m MLModelAssetResourceFactoryInMemoryImpl) ModelAssetDescriptionWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelAssetDescriptionWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/modelStructureWithError:
func (m MLModelAssetResourceFactoryInMemoryImpl) ModelStructureWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/modelWithConfiguration:error:
func (m MLModelAssetResourceFactoryInMemoryImpl) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/initWithArchiveData:
func (m MLModelAssetResourceFactoryInMemoryImpl) InitWithArchiveData(data objectivec.IObject) MLModelAssetResourceFactoryInMemoryImpl {
	rv := objc.Send[MLModelAssetResourceFactoryInMemoryImpl](m.ID, objc.Sel("initWithArchiveData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/compiledModelURL
func (m MLModelAssetResourceFactoryInMemoryImpl) CompiledModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/debugDescription
func (m MLModelAssetResourceFactoryInMemoryImpl) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/description
func (m MLModelAssetResourceFactoryInMemoryImpl) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/hash
func (m MLModelAssetResourceFactoryInMemoryImpl) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryInMemoryImpl/superclass
func (m MLModelAssetResourceFactoryInMemoryImpl) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
