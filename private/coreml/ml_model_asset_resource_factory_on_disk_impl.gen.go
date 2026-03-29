// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetResourceFactoryOnDiskImpl] class.
var (
	_MLModelAssetResourceFactoryOnDiskImplClass     MLModelAssetResourceFactoryOnDiskImplClass
	_MLModelAssetResourceFactoryOnDiskImplClassOnce sync.Once
)

func getMLModelAssetResourceFactoryOnDiskImplClass() MLModelAssetResourceFactoryOnDiskImplClass {
	_MLModelAssetResourceFactoryOnDiskImplClassOnce.Do(func() {
		_MLModelAssetResourceFactoryOnDiskImplClass = MLModelAssetResourceFactoryOnDiskImplClass{class: objc.GetClass("MLModelAssetResourceFactoryOnDiskImpl")}
	})
	return _MLModelAssetResourceFactoryOnDiskImplClass
}

// GetMLModelAssetResourceFactoryOnDiskImplClass returns the class object for MLModelAssetResourceFactoryOnDiskImpl.
func GetMLModelAssetResourceFactoryOnDiskImplClass() MLModelAssetResourceFactoryOnDiskImplClass {
	return getMLModelAssetResourceFactoryOnDiskImplClass()
}

type MLModelAssetResourceFactoryOnDiskImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetResourceFactoryOnDiskImplClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetResourceFactoryOnDiskImplClass) Alloc() MLModelAssetResourceFactoryOnDiskImpl {
	rv := objc.Send[MLModelAssetResourceFactoryOnDiskImpl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAssetResourceFactoryOnDiskImpl.CompiledModelURL]
//   - [MLModelAssetResourceFactoryOnDiskImpl.ModelAssetDescriptionWithError]
//   - [MLModelAssetResourceFactoryOnDiskImpl.ModelStructureWithError]
//   - [MLModelAssetResourceFactoryOnDiskImpl.ModelURL]
//   - [MLModelAssetResourceFactoryOnDiskImpl.ModelWithConfigurationError]
//   - [MLModelAssetResourceFactoryOnDiskImpl.InitWithModelURLError]
//   - [MLModelAssetResourceFactoryOnDiskImpl.DebugDescription]
//   - [MLModelAssetResourceFactoryOnDiskImpl.Description]
//   - [MLModelAssetResourceFactoryOnDiskImpl.Hash]
//   - [MLModelAssetResourceFactoryOnDiskImpl.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl
type MLModelAssetResourceFactoryOnDiskImpl struct {
	objectivec.Object
}

// MLModelAssetResourceFactoryOnDiskImplFromID constructs a [MLModelAssetResourceFactoryOnDiskImpl] from an objc.ID.
func MLModelAssetResourceFactoryOnDiskImplFromID(id objc.ID) MLModelAssetResourceFactoryOnDiskImpl {
	return MLModelAssetResourceFactoryOnDiskImpl{objectivec.Object{ID: id}}
}
// Ensure MLModelAssetResourceFactoryOnDiskImpl implements IMLModelAssetResourceFactoryOnDiskImpl.
var _ IMLModelAssetResourceFactoryOnDiskImpl = MLModelAssetResourceFactoryOnDiskImpl{}

// An interface definition for the [MLModelAssetResourceFactoryOnDiskImpl] class.
//
// # Methods
//
//   - [IMLModelAssetResourceFactoryOnDiskImpl.CompiledModelURL]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.ModelAssetDescriptionWithError]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.ModelStructureWithError]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.ModelURL]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.ModelWithConfigurationError]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.InitWithModelURLError]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.DebugDescription]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.Description]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.Hash]
//   - [IMLModelAssetResourceFactoryOnDiskImpl.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl
type IMLModelAssetResourceFactoryOnDiskImpl interface {
	objectivec.IObject

	// Topic: Methods

	CompiledModelURL() foundation.INSURL
	ModelAssetDescriptionWithError() (objectivec.IObject, error)
	ModelStructureWithError() (objectivec.IObject, error)
	ModelURL() foundation.INSURL
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	InitWithModelURLError(url foundation.INSURL) (MLModelAssetResourceFactoryOnDiskImpl, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLModelAssetResourceFactoryOnDiskImpl) Init() MLModelAssetResourceFactoryOnDiskImpl {
	rv := objc.Send[MLModelAssetResourceFactoryOnDiskImpl](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetResourceFactoryOnDiskImpl) Autorelease() MLModelAssetResourceFactoryOnDiskImpl {
	rv := objc.Send[MLModelAssetResourceFactoryOnDiskImpl](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetResourceFactoryOnDiskImpl creates a new MLModelAssetResourceFactoryOnDiskImpl instance.
func NewMLModelAssetResourceFactoryOnDiskImpl() MLModelAssetResourceFactoryOnDiskImpl {
	class := getMLModelAssetResourceFactoryOnDiskImplClass()
	rv := objc.Send[MLModelAssetResourceFactoryOnDiskImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/initWithModelURL:error:
func NewModelAssetResourceFactoryOnDiskImplWithModelURLError(url foundation.INSURL) (MLModelAssetResourceFactoryOnDiskImpl, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetResourceFactoryOnDiskImplClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetResourceFactoryOnDiskImpl{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetResourceFactoryOnDiskImplFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/modelAssetDescriptionWithError:
func (m MLModelAssetResourceFactoryOnDiskImpl) ModelAssetDescriptionWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelAssetDescriptionWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/modelStructureWithError:
func (m MLModelAssetResourceFactoryOnDiskImpl) ModelStructureWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/modelWithConfiguration:error:
func (m MLModelAssetResourceFactoryOnDiskImpl) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/initWithModelURL:error:
func (m MLModelAssetResourceFactoryOnDiskImpl) InitWithModelURLError(url foundation.INSURL) (MLModelAssetResourceFactoryOnDiskImpl, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetResourceFactoryOnDiskImpl{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetResourceFactoryOnDiskImplFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/compiledModelURL
func (m MLModelAssetResourceFactoryOnDiskImpl) CompiledModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/debugDescription
func (m MLModelAssetResourceFactoryOnDiskImpl) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/description
func (m MLModelAssetResourceFactoryOnDiskImpl) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/hash
func (m MLModelAssetResourceFactoryOnDiskImpl) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/modelURL
func (m MLModelAssetResourceFactoryOnDiskImpl) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactoryOnDiskImpl/superclass
func (m MLModelAssetResourceFactoryOnDiskImpl) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

