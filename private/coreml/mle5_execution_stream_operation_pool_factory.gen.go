// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ExecutionStreamOperationPoolFactory] class.
var (
	_MLE5ExecutionStreamOperationPoolFactoryClass     MLE5ExecutionStreamOperationPoolFactoryClass
	_MLE5ExecutionStreamOperationPoolFactoryClassOnce sync.Once
)

func getMLE5ExecutionStreamOperationPoolFactoryClass() MLE5ExecutionStreamOperationPoolFactoryClass {
	_MLE5ExecutionStreamOperationPoolFactoryClassOnce.Do(func() {
		_MLE5ExecutionStreamOperationPoolFactoryClass = MLE5ExecutionStreamOperationPoolFactoryClass{class: objc.GetClass("MLE5ExecutionStreamOperationPoolFactory")}
	})
	return _MLE5ExecutionStreamOperationPoolFactoryClass
}

// GetMLE5ExecutionStreamOperationPoolFactoryClass returns the class object for MLE5ExecutionStreamOperationPoolFactory.
func GetMLE5ExecutionStreamOperationPoolFactoryClass() MLE5ExecutionStreamOperationPoolFactoryClass {
	return getMLE5ExecutionStreamOperationPoolFactoryClass()
}

type MLE5ExecutionStreamOperationPoolFactoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ExecutionStreamOperationPoolFactoryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ExecutionStreamOperationPoolFactoryClass) Alloc() MLE5ExecutionStreamOperationPoolFactory {
	rv := objc.Send[MLE5ExecutionStreamOperationPoolFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPoolFactory
type MLE5ExecutionStreamOperationPoolFactory struct {
	objectivec.Object
}

// MLE5ExecutionStreamOperationPoolFactoryFromID constructs a [MLE5ExecutionStreamOperationPoolFactory] from an objc.ID.
func MLE5ExecutionStreamOperationPoolFactoryFromID(id objc.ID) MLE5ExecutionStreamOperationPoolFactory {
	return MLE5ExecutionStreamOperationPoolFactory{objectivec.Object{ID: id}}
}

// Ensure MLE5ExecutionStreamOperationPoolFactory implements IMLE5ExecutionStreamOperationPoolFactory.
var _ IMLE5ExecutionStreamOperationPoolFactory = MLE5ExecutionStreamOperationPoolFactory{}

// An interface definition for the [MLE5ExecutionStreamOperationPoolFactory] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPoolFactory
type IMLE5ExecutionStreamOperationPoolFactory interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e MLE5ExecutionStreamOperationPoolFactory) Init() MLE5ExecutionStreamOperationPoolFactory {
	rv := objc.Send[MLE5ExecutionStreamOperationPoolFactory](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5ExecutionStreamOperationPoolFactory) Autorelease() MLE5ExecutionStreamOperationPoolFactory {
	rv := objc.Send[MLE5ExecutionStreamOperationPoolFactory](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStreamOperationPoolFactory creates a new MLE5ExecutionStreamOperationPoolFactory instance.
func NewMLE5ExecutionStreamOperationPoolFactory() MLE5ExecutionStreamOperationPoolFactory {
	class := getMLE5ExecutionStreamOperationPoolFactoryClass()
	rv := objc.Send[MLE5ExecutionStreamOperationPoolFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPoolFactory/createPoolFromLibrary:functionName:modelDescription:modelConfiguration:modelSignpostId:compilerVersionInfo:
func (_MLE5ExecutionStreamOperationPoolFactoryClass MLE5ExecutionStreamOperationPoolFactoryClass) CreatePoolFromLibraryFunctionNameModelDescriptionModelConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLE5ExecutionStreamOperationPoolFactoryClass.class), objc.Sel("createPoolFromLibrary:functionName:modelDescription:modelConfiguration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return objectivec.Object{ID: rv}
}
