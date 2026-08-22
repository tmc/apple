// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
type IMLE5ProgramLibraryOnDeviceAOTCompilationImpl interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	Container() IMLProgramE5Container
	CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error)
	ModelDisplayName() string
	SerializedMILText() string
	SetSerializedMILText(value string)
	InitWithIRProgramContainerConfigurationDeallocator(iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject, deallocator VoidHandler) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
	InitWithMILTextAtURLContainerConfiguration(url foundation.NSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
	InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration(url foundation.NSURL, program unsafe.Pointer, deallocator VoidHandler, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl
}

// Init initializes the instance.
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Init() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Autorelease() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ProgramLibraryOnDeviceAOTCompilationImpl creates a new MLE5ProgramLibraryOnDeviceAOTCompilationImpl instance.
func NewMLE5ProgramLibraryOnDeviceAOTCompilationImpl() MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	class := getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass()
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewE5ProgramLibraryOnDeviceAOTCompilationImplWithMILTextAtURLContainerConfiguration(url foundation.NSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	instance := getMLE5ProgramLibraryOnDeviceAOTCompilationImplClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMILTextAtURL:container:configuration:"), url, container, configuration)
	return MLE5ProgramLibraryOnDeviceAOTCompilationImplFromID(rv)
}

func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("createProgramLibraryHandleWithRespecialization:error:"), respecialization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(E5rtProgramLibraryRef), foundation.NSErrorFrom(errorPtr)
	}
	return E5rtProgramLibraryRef(rv), nil

}

var _mle5programlibraryondeviceaotcompilationimpl_initwithirprogram_container_configuration_deallocator_p3_key byte

func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithIRProgramContainerConfigurationDeallocator(iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject, deallocator VoidHandler) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	_block3, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](m.ID, objc.Sel("initWithIRProgram:container:configuration:deallocator:"), iRProgram, container, configuration, _block3)
	return rv
}
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithMILTextAtURLContainerConfiguration(url foundation.NSURL, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](m.ID, objc.Sel("initWithMILTextAtURL:container:configuration:"), url, container, configuration)
	return rv
}

var _mle5programlibraryondeviceaotcompilationimpl_initwithmiltextaturl_irprogram_deallocator_container_configuration_p2_key byte

func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration(url foundation.NSURL, program unsafe.Pointer, deallocator VoidHandler, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibraryOnDeviceAOTCompilationImpl {
	_block2, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[MLE5ProgramLibraryOnDeviceAOTCompilationImpl](m.ID, objc.Sel("initWithMILTextAtURL:irProgram:deallocator:container:configuration:"), url, program, _block2, container, configuration)
	return rv
}

func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) Container() IMLProgramE5Container {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("container"))
	return MLProgramE5ContainerFromID(objc.ID(rv))
}
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) ModelDisplayName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDisplayName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) SerializedMILText() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) SetSerializedMILText(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSerializedMILText:"), objc.String(value))
}

// InitWithIRProgramContainerConfigurationDeallocatorSync is a synchronous wrapper around [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLE5ProgramLibraryOnDeviceAOTCompilationImpl) InitWithIRProgramContainerConfigurationDeallocatorSync(ctx context.Context, iRProgram unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) error {
	done := make(chan struct{}, 1)
	m.InitWithIRProgramContainerConfigurationDeallocator(iRProgram, container, configuration, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
