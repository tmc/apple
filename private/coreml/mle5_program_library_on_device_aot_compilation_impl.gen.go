// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ProgramLibraryOnDeviceAOTCompilationImpl] class.
var (
	_MLE5ProgramLibraryOnDeviceAOTCompilationImplClass     MLE5ProgramLibraryOnDeviceAOTCompilationImplClass
	_MLE5ProgramLibraryOnDeviceAOTCompilationImplClassOnce sync.Once
)

func getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass() MLE5ProgramLibraryOnDeviceAOTCompilationImplClass {
	_MLE5ProgramLibraryOnDeviceAOTCompilationImplClassOnce.Do(func() {
		_MLE5ProgramLibraryOnDeviceAOTCompilationImplClass = MLE5ProgramLibraryOnDeviceAOTCompilationImplClass{class: objc.GetClass("MLE5ProgramLibraryOnDeviceAOTCompilationImpl")}
	})
	return _MLE5ProgramLibraryOnDeviceAOTCompilationImplClass
}

// GetMLE5ProgramLibraryOnDeviceAOTCompilationImplClass returns the class object for MLE5ProgramLibraryOnDeviceAOTCompilationImpl.
func GetMLE5ProgramLibraryOnDeviceAOTCompilationImplClass() MLE5ProgramLibraryOnDeviceAOTCompilationImplClass {
	return getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass()
}

type MLE5ProgramLibraryOnDeviceAOTCompilationImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ProgramLibraryOnDeviceAOTCompilationImplClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ProgramLibraryOnDeviceAOTCompilationImplClass) Alloc() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.Send[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.Configuration]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.Container]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.CreateProgramLibraryHandleWithRespecializationError]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.ModelDisplayName]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.SerializedMILText]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.SetSerializedMILText]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLContainerConfiguration]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration]
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl
type MLE5ProgramLibraryOnDeviceAOTCompilationImpl struct {
	objectivec.Object
}

// MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID constructs a [MLE5ProgramLibraryOnDeviceAOTCompilationImpl] from an objc.ID.
func MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID(id objc.ID) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	return MLE5ProgramLibraryOnDeviceAOTCompilationImpl{objectivec.Object{ID: id}}
}
// Ensure MLE5ProgramLibraryOnDeviceAOTCompilationImpl implements IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.
var _ IMLE5ProgramLibraryOnDeviceAOTCompilationImpl = MLE5ProgramLibraryOnDeviceAOTCompilationImpl{}

// An interface definition for the [MLE5ProgramLibraryOnDeviceAOTCompilationImpl] class.
//
// # Methods
//
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.Configuration]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.Container]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.CreateProgramLibraryHandleWithRespecializationError]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.ModelDisplayName]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.SerializedMILText]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.SetSerializedMILText]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLContainerConfiguration]
//   - [IMLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl
type IMLE5ProgramLibraryOnDeviceAOTCompilationImpl interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	Container() IMLProgramE5Container
	CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (objectivec.IObject, error)
	ModelDisplayName() string
	SerializedMILText() string
	SetSerializedMILText(value string)
	InitWithIRProgramContainerConfigurationDeallocator(iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject, deallocator VoidHandler) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
	InitWithMILTextAtURLContainerConfiguration(url foundation.INSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
	InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration(url foundation.INSURL, program unsafe.Pointer, deallocator VoidHandler, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
}

// Init initializes the instance.
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Init() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.Send[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Autorelease() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.Send[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ProgramLibraryOnDeviceAOTCompilationImpl creates a new MLE5ProgramLibraryOnDeviceAOTCompilationImpl instance.
func NewMLE5ProgramLibraryOnDeviceAOTCompilationImpl() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	class := getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass()
	rv := objc.Send[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/initWithMILTextAtURL:container:configuration:
func NewE5ProgramLibraryOnDeviceAOTCompilationImplWithMILTextAtURLContainerConfiguration(url foundation.INSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	instance := getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMILTextAtURL:container:configuration:"), url, container, configuration)
	return MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/createProgramLibraryHandleWithRespecialization:error:
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("createProgramLibraryHandleWithRespecialization:error:"), respecialization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/initWithIRProgram:container:configuration:deallocator:
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithIRProgramContainerConfigurationDeallocator(iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject, deallocator VoidHandler) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
_block3, _ := NewVoidBlock(deallocator)
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithIRProgram:container:configuration:deallocator:"), iRProgram, container, configuration, _block3)
	return MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/initWithMILTextAtURL:container:configuration:
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithMILTextAtURLContainerConfiguration(url foundation.INSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.Send[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](e.ID, objc.Sel("initWithMILTextAtURL:container:configuration:"), url, container, configuration)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/initWithMILTextAtURL:irProgram:deallocator:container:configuration:
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration(url foundation.INSURL, program unsafe.Pointer, deallocator VoidHandler, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
_block2, _ := NewVoidBlock(deallocator)
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithMILTextAtURL:irProgram:deallocator:container:configuration:"), url, program, _block2, container, configuration)
	return MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/configuration
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/container
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Container() IMLProgramE5Container {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("container"))
	return MLProgramE5ContainerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/modelDisplayName
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) ModelDisplayName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDisplayName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryOnDeviceAOTCompilationImpl/serializedMILText
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) SerializedMILText() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) SetSerializedMILText(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setSerializedMILText:"), objc.String(value))
}

// InitWithIRProgramContainerConfigurationDeallocatorSync is a synchronous wrapper around [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (e MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithIRProgramContainerConfigurationDeallocatorSync(ctx context.Context, iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.InitWithIRProgramContainerConfigurationDeallocator(iRProgram, container, configuration, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

